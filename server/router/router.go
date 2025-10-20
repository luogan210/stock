package router

import (
	"go-demo/config"
	"go-demo/handler"
	"go-demo/middleware"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	cfg := config.Load()

	// Set gin mode based on env
	if cfg.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Optionally trust proxy
	if cfg.TrustProxy {
		_ = r.SetTrustedProxies(nil) // trust all; change to specific CIDRs if needed
	}

	// 设置前端路由
	SetupFrontendRoutes(r)

	// API路由组 - 应用认证中间件
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())
	{
		// 注册API路由
		handler.RegisterRoutes(api)
	}

	return r
}
