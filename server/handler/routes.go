package handler

import (
	"github.com/gin-gonic/gin"
)

// RegisterRoutes 注册所有路由
func RegisterRoutes(rg *gin.RouterGroup) {
	// 注册所有业务路由
	RegisterUploadRoutes(rg)
	RegisterStockRoutes(rg)
	RegisterPlanRoutes(rg)
	RegisterLogRoutes(rg)
	RegisterReviewRoutes(rg)
}
