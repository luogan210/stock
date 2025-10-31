package router

import (
	"github.com/gin-gonic/gin"
	"server/config"
	"server/handler"
	"server/middleware"
	"server/modules"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	cfg := config.Load()

	// Set gin mode based on env
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// 添加请求日志中间件
	r.Use(middleware.RequestLogger())

	// Optionally trust proxy
	if cfg.TrustProxy {
		_ = r.SetTrustedProxies(nil) // trust all; change to specific CIDRs if needed
	}

	// 设置前端路由（可选，保留静态页面与上传示例）
	SetupFrontendRoutes(r)

	// API路由组 - 应用认证中间件
	api := r.Group("/api")

	{
		// 注册所有模块路由
		modules.RegisterAllRoutes(api)
		// 注册其他路由（如上传等）
		handler.RegisterUploadRoutes(api)
	}

	return r
}
