package review

import (
	"strconv"

	"server/handler"

	"github.com/gin-gonic/gin"
)

// RegisterReviewRoutes 注册复盘路由
func RegisterReviewRoutes(r *gin.RouterGroup) {
	reviewService := NewReviewService()

	g := r.Group("/reviews")
	{
		g.POST("/create", func(c *gin.Context) {
			var req ReviewCreateRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
				return
			}

			review, err := reviewService.CreateReview(&req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, gin.H{"id": review.ID})
		})

		g.GET("/getList", func(c *gin.Context) {
			req := &ReviewListRequest{
				Keyword:   c.Query("keyword"),
				Period:    c.Query("period"),
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

			response, err := reviewService.ListReviews(req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, response)
		})

		g.GET("/getDetail/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "复盘ID不能为空")
				return
			}

			review, err := reviewService.GetReview(id)
			if err != nil {
				handler.Error(c, handler.CodeNotFound, err.Error())
				return
			}

			handler.Success(c, review)
		})

		g.PUT("/update/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "复盘ID不能为空")
				return
			}

			var req ReviewUpdateRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
				return
			}

			review, err := reviewService.UpdateReview(id, &req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, review)
		})

		g.DELETE("/delete/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "复盘ID不能为空")
				return
			}

			err := reviewService.DeleteReview(id)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, gin.H{"id": id})
		})
	}
}
