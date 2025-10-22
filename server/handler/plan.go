package handler

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// PlanHandler 交易计划处理
type PlanHandler struct {
	*BaseHandler
}

// Plan 交易计划结构
type Plan struct {
	ID              string  `json:"id"`
	Name            string  `json:"name" binding:"required"`
	Type            string  `json:"type" binding:"required"`
	StockCode       string  `json:"stockCode" binding:"required"`
	StockName       string  `json:"stockName" binding:"required"`
	Strategy        string  `json:"strategy"`
	TradingStrategy string  `json:"tradingStrategy"`
	TargetPrice     float64 `json:"targetPrice"`
	Quantity        int     `json:"quantity"`
	StopLoss        float64 `json:"stopLoss"`
	TakeProfit      float64 `json:"takeProfit"`
	StartTime       string  `json:"startTime"`
	EndTime         string  `json:"endTime"`
	RiskLevel       string  `json:"riskLevel"`
	Description     string  `json:"description"`
	Remark          string  `json:"remark"`
	Status          string  `json:"status"`
	CreatedAt       string  `json:"createdAt"`
	UpdatedAt       string  `json:"updatedAt"`
}

// NewPlanHandler 创建交易计划处理器实例
func NewPlanHandler() *PlanHandler {
	return &PlanHandler{
		BaseHandler: NewBaseHandler(),
	}
}

// RegisterPlanRoutes 注册交易计划路由
func RegisterPlanRoutes(r *gin.RouterGroup) {
	handler := NewPlanHandler()
	g := r.Group("/plans")
	g.GET("/getList", handler.ListPlans)        // GET /api/plans/getList - 获取计划列表
	g.POST("/create", handler.CreatePlan)       // POST /api/plans/create - 创建计划
	g.GET("/getDetail/:id", handler.GetPlan)    // GET /api/plans/getDetail/:id - 获取计划详情
	g.PUT("/update/:id", handler.UpdatePlan)    // PUT /api/plans/update/:id - 更新计划
	g.DELETE("/delete/:id", handler.DeletePlan) // DELETE /api/plans/delete/:id - 删除计划
}

// getDB 从上下文获取数据库连接
func (h *PlanHandler) getDB(c *gin.Context) *sql.DB {
	if db, exists := c.Get("db"); exists {
		return db.(*sql.DB)
	}
	return nil
}

// ListPlans 获取交易计划列表
func (h *PlanHandler) ListPlans(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	keyword := strings.TrimSpace(c.Query("keyword"))
	status := c.Query("status")
	stockCode := c.Query("stockCode")

	q := `SELECT id, name, type, stock_code, stock_name, strategy, trading_strategy, target_price, quantity, stop_loss, take_profit, start_time, end_time, risk_level, description, remark, status, created_at, updated_at FROM plans WHERE 1=1`
	args := []any{}
	if keyword != "" {
		q += " AND (name LIKE ? OR description LIKE ?)"
		like := "%" + keyword + "%"
		args = append(args, like, like)
	}
	if status != "" {
		q += " AND status = ?"
		args = append(args, status)
	}
	if stockCode != "" {
		q += " AND stock_code = ?"
		args = append(args, stockCode)
	}
	q += " ORDER BY created_at DESC"

	rows, err := db.Query(q, args...)
	if err != nil {
		h.ServerError(c, "查询失败: "+err.Error())
		return
	}
	defer rows.Close()

	var items []Plan
	for rows.Next() {
		var p Plan
		err := rows.Scan(&p.ID, &p.Name, &p.Type, &p.StockCode, &p.StockName, &p.Strategy, &p.TradingStrategy, &p.TargetPrice, &p.Quantity, &p.StopLoss, &p.TakeProfit, &p.StartTime, &p.EndTime, &p.RiskLevel, &p.Description, &p.Remark, &p.Status, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			h.ServerError(c, "数据解析失败: "+err.Error())
			return
		}
		items = append(items, p)
	}
	h.Success(c, gin.H{"list": items})
}

// CreatePlan 创建交易计划
func (h *PlanHandler) CreatePlan(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	var p Plan
	if err := c.ShouldBindJSON(&p); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	if p.ID == "" {
		p.ID = uuid.New().String()
	}
	if p.Status == "" {
		p.Status = "active"
	}

	_, err := db.Exec(`INSERT INTO plans (id, name, type, stock_code, stock_name, strategy, trading_strategy, target_price, quantity, stop_loss, take_profit, start_time, end_time, risk_level, description, remark, status) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
		p.ID, p.Name, p.Type, p.StockCode, p.StockName, p.Strategy, p.TradingStrategy, p.TargetPrice, p.Quantity, p.StopLoss, p.TakeProfit, p.StartTime, p.EndTime, p.RiskLevel, p.Description, p.Remark, p.Status,
	)
	if err != nil {
		h.ServerError(c, "创建失败: "+err.Error())
		return
	}

	h.Success(c, p)
}

// GetPlan 获取交易计划详情
func (h *PlanHandler) GetPlan(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	var p Plan
	err := db.QueryRow(`SELECT id, name, type, stock_code, stock_name, strategy, trading_strategy, target_price, quantity, stop_loss, take_profit, start_time, end_time, risk_level, description, remark, status, created_at, updated_at FROM plans WHERE id = ?`, id).
		Scan(&p.ID, &p.Name, &p.Type, &p.StockCode, &p.StockName, &p.Strategy, &p.TradingStrategy, &p.TargetPrice, &p.Quantity, &p.StopLoss, &p.TakeProfit, &p.StartTime, &p.EndTime, &p.RiskLevel, &p.Description, &p.Remark, &p.Status, &p.CreatedAt, &p.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			h.NotFoundError(c, "交易计划不存在")
		} else {
			h.ServerError(c, "查询失败: "+err.Error())
		}
		return
	}

	h.Success(c, p)
}

// UpdatePlan 更新交易计划
func (h *PlanHandler) UpdatePlan(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	var p Plan
	if err := c.ShouldBindJSON(&p); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	_, err := db.Exec(`UPDATE plans SET name=?, type=?, stock_code=?, stock_name=?, strategy=?, trading_strategy=?, target_price=?, quantity=?, stop_loss=?, take_profit=?, start_time=?, end_time=?, risk_level=?, description=?, remark=?, status=? WHERE id=?`,
		p.Name, p.Type, p.StockCode, p.StockName, p.Strategy, p.TradingStrategy, p.TargetPrice, p.Quantity, p.StopLoss, p.TakeProfit, p.StartTime, p.EndTime, p.RiskLevel, p.Description, p.Remark, p.Status, id)
	if err != nil {
		h.ServerError(c, "更新失败: "+err.Error())
		return
	}

	h.Success(c, gin.H{"id": id})
}

// DeletePlan 删除交易计划
func (h *PlanHandler) DeletePlan(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	_, err := db.Exec(`DELETE FROM plans WHERE id=?`, id)
	if err != nil {
		h.ServerError(c, "删除失败: "+err.Error())
		return
	}

	h.Success(c, gin.H{"id": id})
}
