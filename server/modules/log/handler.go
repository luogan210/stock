package log

import (
	"strconv"

	"server/handler"

	"github.com/gin-gonic/gin"
)

// RegisterLogRoutes 注册日志路由
func RegisterLogRoutes(r *gin.RouterGroup) {
	logService := NewLogService()

	g := r.Group("/logs")
	{
		g.POST("/create", func(c *gin.Context) {
			var req LogCreateRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
				return
			}

			log, err := logService.CreateLog(&req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, gin.H{"id": log.ID})
		})

		g.GET("/getList", func(c *gin.Context) {
			req := &LogListRequest{
				Keyword:   c.Query("keyword"),
				Type:      c.Query("type"),
				StockCode: c.Query("stockCode"),
				PlanName:  c.Query("planName"),
				StartDate: c.Query("startDate"),
				EndDate:   c.Query("endDate"),
			}

			// 解析分页参数
			if pageStr := c.Query("page"); pageStr != "" {
				if page, err := strconv.Atoi(pageStr); err == nil && page > 0 {
					req.Page = page
				}
			}
			if pageSizeStr := c.Query("pageSize"); pageSizeStr != "" {
				if pageSize, err := strconv.Atoi(pageSizeStr); err == nil && pageSize > 0 {
					req.PageSize = pageSize
				}
			}

			response, err := logService.ListLogs(req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, response)
		})

		g.GET("/getDetail/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "日志ID不能为空")
				return
			}

			log, err := logService.GetLog(id)
			if err != nil {
				handler.Error(c, handler.CodeNotFound, err.Error())
				return
			}

			handler.Success(c, log)
		})

		g.PUT("/update/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "日志ID不能为空")
				return
			}

			var req LogUpdateRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
				return
			}

			log, err := logService.UpdateLog(id, &req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, log)
		})

		g.DELETE("/delete/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "日志ID不能为空")
				return
			}

			err := logService.DeleteLog(id)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, gin.H{"id": id})
		})
	}
}
