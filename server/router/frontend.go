package router

import (
	"server/handler"

	"github.com/gin-gonic/gin"
)

// SetupFrontendRoutes 设置前端路由
func SetupFrontendRoutes(r *gin.Engine) {
	// 设置HTML模板
	r.LoadHTMLGlob("frontend/templates/*.html")

	// 静态文件服务
	r.Static("/static", "./frontend/static")

	// 创建前端处理器
	frontendHandler := handler.NewFrontendHandler()

	// 首页
	r.GET("/", frontendHandler.ServeIndex)

	// 登录页面
	r.GET("/login", frontendHandler.ServeLogin)
	r.GET("/login.html", frontendHandler.ServeLogin)

	// 注册页面
	r.GET("/register", frontendHandler.ServeRegister)
	r.GET("/register.html", frontendHandler.ServeRegister)

	// 上传页面
	r.GET("/upload", frontendHandler.ServeUpload)
	r.GET("/upload.html", frontendHandler.ServeUpload)
}
