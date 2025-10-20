package handler

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(rg *gin.RouterGroup) {
	// 注册Hello相关路由
	RegisterHelloRoutes(rg)

	// 注册用户路由
	RegisterUserRoutes(rg)

	// 注册示例路由
	RegisterExampleRoutes(rg)

	// 注册微信相关路由
	RegisterWechatRoutes(rg)

	// 注册文件上传路由
	RegisterUploadRoutes(rg)
}
