package handler

import (
	"crypto/md5"
	"fmt"
	"go-demo/config"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// UploadHandler 文件上传处理
type UploadHandler struct {
	*BaseHandler
	uploadDir string
	uploads   map[string]*UploadSession
	mu        sync.RWMutex
}

// UploadSession 上传会话
type UploadSession struct {
	FileID         string
	FileName       string
	FileSize       int64
	ChunkSize      int
	TotalChunks    int
	UploadedChunks map[int]bool
	CreatedAt      time.Time
	mu             sync.Mutex
}

// NewUploadHandler 创建上传处理器实例
func NewUploadHandler() *UploadHandler {
	cfg := config.Load()
	return &UploadHandler{
		BaseHandler: NewBaseHandler(),
		uploadDir:   cfg.UploadDir,
		uploads:     make(map[string]*UploadSession),
	}
}

// InitUpload 初始化上传
func (h *UploadHandler) InitUpload(c *gin.Context) {
	var req struct {
		FileID    string `json:"fileId" binding:"required"`
		FileName  string `json:"fileName" binding:"required"`
		FileSize  int64  `json:"fileSize" binding:"required"`
		ChunkSize int    `json:"chunkSize" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	cfg := config.Load()

	// 验证文件大小限制
	if req.FileSize > cfg.MaxUploadSizeBytes {
		h.ParamError(c, "文件大小超过限制")
		return
	}

	// 强制分片大小不超过配置
	if req.ChunkSize <= 0 || req.ChunkSize > cfg.ChunkSizeBytes {
		req.ChunkSize = cfg.ChunkSizeBytes
	}

	// 创建上传会话
	session := &UploadSession{
		FileID:         req.FileID,
		FileName:       req.FileName,
		FileSize:       req.FileSize,
		ChunkSize:      req.ChunkSize,
		TotalChunks:    int((req.FileSize + int64(req.ChunkSize) - 1) / int64(req.ChunkSize)),
		UploadedChunks: make(map[int]bool),
		CreatedAt:      time.Now(),
	}

	h.mu.Lock()
	h.uploads[req.FileID] = session
	h.mu.Unlock()

	// 创建临时目录
	tempDir := filepath.Join(h.uploadDir, req.FileID)
	if err := os.MkdirAll(tempDir, 0755); err != nil {
		h.ServerError(c, "创建临时目录失败")
		return
	}

	h.Success(c, gin.H{
		"file_id":      req.FileID,
		"total_chunks": session.TotalChunks,
		"chunk_size":   req.ChunkSize,
	})
}

// UploadChunk 上传分片
func (h *UploadHandler) UploadChunk(c *gin.Context) {
	fileID := c.PostForm("fileId")
	chunkIndexStr := c.PostForm("chunkIndex")

	if fileID == "" || chunkIndexStr == "" {
		h.ParamError(c, "缺少必要参数")
		return
	}

	chunkIndex, err := strconv.Atoi(chunkIndexStr)
	if err != nil {
		h.ParamError(c, "分片索引格式错误")
		return
	}

	// 获取上传会话
	h.mu.RLock()
	session, exists := h.uploads[fileID]
	h.mu.RUnlock()

	if !exists {
		h.ParamError(c, "上传会话不存在")
		return
	}

	// 获取上传的文件
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		h.ServerError(c, "获取上传文件失败")
		return
	}
	defer file.Close()

	// 保存分片文件
	chunkPath := filepath.Join(h.uploadDir, fileID, fmt.Sprintf("chunk_%d", chunkIndex))
	chunkFile, err := os.Create(chunkPath)
	if err != nil {
		h.ServerError(c, "创建分片文件失败")
		return
	}
	defer chunkFile.Close()

	// 计算MD5校验
	hash := md5.New()
	multiWriter := io.MultiWriter(chunkFile, hash)

	_, err = io.Copy(multiWriter, file)
	if err != nil {
		h.ServerError(c, "保存分片文件失败")
		return
	}

	// 更新会话状态
	session.mu.Lock()
	session.UploadedChunks[chunkIndex] = true
	session.mu.Unlock()

	h.Success(c, gin.H{
		"chunk_index":     chunkIndex,
		"uploaded_chunks": len(session.UploadedChunks),
		"total_chunks":    session.TotalChunks,
		"file_name":       header.Filename,
	})
}

// CompleteUpload 完成上传
func (h *UploadHandler) CompleteUpload(c *gin.Context) {
	// 先尝试手动解析JSON
	var jsonData map[string]interface{}
	if err := c.ShouldBindJSON(&jsonData); err != nil {
		h.ParamError(c, "JSON解析错误: "+err.Error())
		return
	}

	// 检查fileId字段是否存在
	fileIdValue, exists := jsonData["fileId"]
	if !exists {
		h.ParamError(c, "缺少fileId参数")
		return
	}

	fileId, ok := fileIdValue.(string)
	if !ok {
		h.ParamError(c, "fileId参数类型错误")
		return
	}

	if fileId == "" {
		h.ParamError(c, "fileId参数不能为空")
		return
	}

	// 获取上传会话
	h.mu.RLock()
	session, exists := h.uploads[fileId]
	h.mu.RUnlock()

	if !exists {
		h.ParamError(c, "上传会话不存在")
		return
	}

	// 检查所有分片是否上传完成
	session.mu.Lock()
	if len(session.UploadedChunks) != session.TotalChunks {
		session.mu.Unlock()
		h.ParamError(c, "文件分片未完全上传")
		return
	}
	session.mu.Unlock()

	// 合并文件
	finalPath := filepath.Join(h.uploadDir, session.FileName)
	finalFile, err := os.Create(finalPath)
	if err != nil {
		h.ServerError(c, "创建最终文件失败")
		return
	}
	defer finalFile.Close()

	// 按顺序合并分片
	for i := 0; i < session.TotalChunks; i++ {
		chunkPath := filepath.Join(h.uploadDir, fileId, fmt.Sprintf("chunk_%d", i))
		chunkFile, err := os.Open(chunkPath)
		if err != nil {
			h.ServerError(c, "读取分片文件失败")
			return
		}

		_, err = io.Copy(finalFile, chunkFile)
		chunkFile.Close()
		if err != nil {
			h.ServerError(c, "合并分片文件失败")
			return
		}
	}

	// 清理临时文件
	tempDir := filepath.Join(h.uploadDir, fileId)
	os.RemoveAll(tempDir)

	// 清理会话
	h.mu.Lock()
	delete(h.uploads, fileId)
	h.mu.Unlock()

	h.Success(c, gin.H{
		"file_path": finalPath,
		"file_size": session.FileSize,
		"file_name": session.FileName,
	})
}

// GetUploadProgress 获取上传进度
func (h *UploadHandler) GetUploadProgress(c *gin.Context) {
	fileID := c.Param("fileId")

	h.mu.RLock()
	session, exists := h.uploads[fileID]
	h.mu.RUnlock()

	if !exists {
		h.ParamError(c, "上传会话不存在")
		return
	}

	session.mu.Lock()
	progress := float64(len(session.UploadedChunks)) / float64(session.TotalChunks) * 100
	session.mu.Unlock()

	h.Success(c, gin.H{
		"progress":        progress,
		"uploaded_chunks": len(session.UploadedChunks),
		"total_chunks":    session.TotalChunks,
		"file_name":       session.FileName,
		"file_size":       session.FileSize,
	})
}

// RegisterUploadRoutes 注册上传相关路由
func RegisterUploadRoutes(rg *gin.RouterGroup) {
	uploadHandler := NewUploadHandler()

	uploads := rg.Group("/upload")
	{
		uploads.POST("/init", uploadHandler.InitUpload)                   // 初始化上传
		uploads.POST("/chunk", uploadHandler.UploadChunk)                 // 上传分片
		uploads.POST("/complete", uploadHandler.CompleteUpload)           // 完成上传
		uploads.GET("/progress/:fileId", uploadHandler.GetUploadProgress) // 获取进度
	}
}
