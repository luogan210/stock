package plan

import (
	"strconv"

	"server/handler"

	"github.com/gin-gonic/gin"
)

// PlanHandler 计划处理器
type PlanHandler struct {
	planService PlanService
}

func NewPlanHandler() *PlanHandler {
	return &PlanHandler{
		planService: NewPlanService(),
	}

}

// RegisterPlanRoutes 注册计划路由
func RegisterPlanRoutes(r *gin.RouterGroup) {
	planService := NewPlanService()

	g := r.Group("/plans")
	{
		g.POST("/create", func(c *gin.Context) {
			var req PlanCreateRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
				return
			}
			plan, err := planService.CreatePlan(&req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}
			handler.Success(c, gin.H{"id": plan.ID})
		})
		g.GET("/getList", func(c *gin.Context) {
			req := &PlanListRequest{
				Keyword:   c.Query("keyword"),
				Type:      c.Query("type"),
				Status:    c.Query("status"),
				RiskLevel: c.Query("riskLevel"),
				StockCode: c.Query("stockCode"),
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

			response, err := planService.ListPlans(req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, response)
		})
		g.GET("/getDetail/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "计划ID不能为空")
				return
			}
			plan, err := planService.GetPlan(id)
			if err != nil {
				handler.Error(c, handler.CodeNotFound, err.Error())
				return
			}

			handler.Success(c, plan)
		})
		g.PUT("/update/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "计划ID不能为空")
				return
			}

			var req PlanUpdateRequest
			if err := c.ShouldBindJSON(&req); err != nil {
				handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
				return
			}

			plan, err := planService.UpdatePlan(id, &req)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, plan)
		})
		g.DELETE("/delete/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "计划ID不能为空")
				return
			}

			err := planService.DeletePlan(id)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, gin.H{"id": id})
		})
		g.PATCH("/status/:id", func(c *gin.Context) {
			id := c.Param("id")
			if id == "" {
				handler.Error(c, handler.CodeInvalid, "计划ID不能为空")
				return
			}

			var req struct {
				Status string `json:"status" binding:"required"`
			}
			if err := c.ShouldBindJSON(&req); err != nil {
				handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
				return
			}

			plan, err := planService.UpdatePlanStatus(id, req.Status)
			if err != nil {
				handler.Error(c, handler.CodeError, err.Error())
				return
			}

			handler.Success(c, plan)
		})
	}
}

// ListPlans 获取计划列表
func (h *PlanHandler) listPlans(c *gin.Context) {
	req := &PlanListRequest{
		Keyword:   c.Query("keyword"),
		Type:      c.Query("type"),
		Status:    c.Query("status"),
		RiskLevel: c.Query("riskLevel"),
		StockCode: c.Query("stockCode"),
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

	response, err := h.planService.ListPlans(req)
	if err != nil {
		handler.Error(c, handler.CodeError, err.Error())
		return
	}

	handler.Success(c, response)
}

// UpdatePlan 更新计划
func (h *PlanHandler) updatePlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handler.Error(c, handler.CodeInvalid, "计划ID不能为空")
		return
	}

	var req PlanUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
		return
	}

	plan, err := h.planService.UpdatePlan(id, &req)
	if err != nil {
		handler.Error(c, handler.CodeError, err.Error())
		return
	}

	handler.Success(c, plan)
}

// DeletePlan 删除计划
func (h *PlanHandler) deletePlan(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handler.Error(c, handler.CodeInvalid, "计划ID不能为空")
		return
	}

	err := h.planService.DeletePlan(id)
	if err != nil {
		handler.Error(c, handler.CodeError, err.Error())
		return
	}

	handler.Success(c, gin.H{"id": id})
}

// UpdatePlanStatus 更新计划状态
func (h *PlanHandler) updatePlanStatus(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handler.Error(c, handler.CodeInvalid, "计划ID不能为空")
		return
	}

	var req struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
		return
	}

	plan, err := h.planService.UpdatePlanStatus(id, req.Status)
	if err != nil {
		handler.Error(c, handler.CodeError, err.Error())
		return
	}

	handler.Success(c, plan)
}
