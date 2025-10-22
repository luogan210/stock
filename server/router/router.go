package router

import (
	"server/config"
	"server/handler"
	"server/storage"

	"github.com/gin-gonic/gin"
)

// SetupRouter 设置路由
func SetupRouter(db *storage.DB) *gin.Engine {
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

	// 设置前端路由（可选，保留静态页面与上传示例）
	SetupFrontendRoutes(r)

	// API路由组 - 应用认证中间件
	api := r.Group("/api")
	// 开发阶段可选择关闭强制鉴权
	// api.Use(middleware.AuthMiddleware())
	// inject db into context
	api.Use(func(c *gin.Context) {
		if db != nil {
			c.Set("db", db.SQL)
		}
		c.Next()
	})
	{
		// 注册API路由
		handler.RegisterRoutes(api)
	}

	return r
}
