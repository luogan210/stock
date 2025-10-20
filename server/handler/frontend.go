package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// FrontendHandler 前端处理器
type FrontendHandler struct {
	BaseHandler
}

// NewFrontendHandler 创建前端处理器
func NewFrontendHandler() *FrontendHandler {
	return &FrontendHandler{}
}

// ServeIndex 服务首页
func (h *FrontendHandler) ServeIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Go Demo - 全栈应用",
	})
}

// ServeLogin 服务登录页面
func (h *FrontendHandler) ServeLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"title": "登录 - Go Demo",
	})
}

// ServeRegister 服务注册页面
func (h *FrontendHandler) ServeRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{
		"title": "注册 - Go Demo",
	})
}

// ServeUpload 服务上传页面
func (h *FrontendHandler) ServeUpload(c *gin.Context) {
	c.HTML(http.StatusOK, "upload.html", gin.H{
		"title": "大文件上传 - Go Demo",
	})
}
