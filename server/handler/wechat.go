package handler

import (
	"go-demo/model"
	"go-demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// WechatHandler 微信小程序处理器
type WechatHandler struct {
	*BaseHandler
	wechatService *service.WechatService
}

// NewWechatHandler 创建微信处理器实例
func NewWechatHandler() *WechatHandler {
	return &WechatHandler{
		BaseHandler:   NewBaseHandler(),
		wechatService: service.NewWechatService(),
	}
}

// GenerateUrlLink 生成微信小程序urlLink
func (h *WechatHandler) GenerateUrlLink(c *gin.Context) {
	// 1. 参数绑定和验证
	var req model.WechatUrlLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	// 2. 调用Service层
	response, err := h.wechatService.GenerateUrlLink(&req)
	if err != nil {
		h.ServerError(c, err.Error())
		return
	}

	// 3. 响应格式化
	h.Success(c, response)
}

// GetUrlLinkInfo 获取urlLink信息
func (h *WechatHandler) GetUrlLinkInfo(c *gin.Context) {
	// 1. 参数验证
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.ParamError(c, "记录ID格式错误")
		return
	}

	// 2. 调用Service层
	urlLinkInfo, err := h.wechatService.GetUrlLinkInfo(uint(id))
	if err != nil {
		h.ServerError(c, err.Error())
		return
	}

	// 3. 响应格式化
	h.Success(c, urlLinkInfo)
}

// RegisterWechatRoutes 注册微信相关路由
func RegisterWechatRoutes(rg *gin.RouterGroup) {
	wechatHandler := NewWechatHandler()

	// 微信相关路由
	wechat := rg.Group("/wechat")
	{
		wechat.POST("/url-link", wechatHandler.GenerateUrlLink)   // 生成urlLink
		wechat.GET("/url-link/:id", wechatHandler.GetUrlLinkInfo) // 获取urlLink信息
	}
}
