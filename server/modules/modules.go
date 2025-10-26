package modules

import (
	"server/modules/log"
	"server/modules/plan"
	"server/modules/review"
	"server/modules/stock"

	"github.com/gin-gonic/gin"
)

// RegisterAllRoutes 注册所有模块的路由
func RegisterAllRoutes(r *gin.RouterGroup) {
	// 注册股票模块路由
	stock.RegisterStockRoutes(r)

	// 注册计划模块路由
	plan.RegisterPlanRoutes(r)

	// 注册日志模块路由
	log.RegisterLogRoutes(r)

	// 注册复盘模块路由
	review.RegisterReviewRoutes(r)
}
