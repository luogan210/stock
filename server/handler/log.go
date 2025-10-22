package handler

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// LogHandler 交易日志处理
type LogHandler struct {
	*BaseHandler
}

// NewLogHandler 创建交易日志处理器实例
func NewLogHandler() *LogHandler {
	return &LogHandler{
		BaseHandler: NewBaseHandler(),
	}
}

// Log 交易日志模型
type Log struct {
	ID          string  `json:"id"`
	PlanName    string  `json:"planName"`
	StockCode   string  `json:"stockCode" binding:"required"`
	StockName   string  `json:"stockName"`
	Type        string  `json:"type" binding:"required"`
	TradingTime string  `json:"tradingTime" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Strategy    string  `json:"strategy"`
	Remark      string  `json:"remark"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
}

// RegisterLogRoutes 注册交易日志相关路由
func RegisterLogRoutes(rg *gin.RouterGroup) {
	h := NewLogHandler()
	g := rg.Group("/logs")
	g.GET("/getList", h.listLogs)        // GET /api/logs/getList - 获取日志列表
	g.POST("/create", h.createLog)       // POST /api/logs/create - 创建日志
	g.GET("/getDetail/:id", h.getLog)    // GET /api/logs/getDetail/:id - 获取日志详情
	g.PUT("/update/:id", h.updateLog)    // PUT /api/logs/update/:id - 更新日志
	g.DELETE("/delete/:id", h.deleteLog) // DELETE /api/logs/delete/:id - 删除日志
}

func (h *LogHandler) getDB(c *gin.Context) *sql.DB {
	v, ok := c.Get("db")
	if !ok {
		return nil
	}
	db, _ := v.(*sql.DB)
	return db
}

func (h *LogHandler) listLogs(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	keyword := strings.TrimSpace(c.Query("keyword"))
	stockCode := c.Query("stockCode")
	typeFilter := c.Query("type")

	q := `SELECT id, plan_name, stock_code, stock_name, type, trading_time, price, quantity, strategy, remark, created_at, updated_at FROM logs WHERE 1=1`
	args := []any{}
	if keyword != "" {
		q += " AND (stock_name LIKE ? OR plan_name LIKE ?)"
		like := "%" + keyword + "%"
		args = append(args, like, like)
	}
	if stockCode != "" {
		q += " AND stock_code = ?"
		args = append(args, stockCode)
	}
	if typeFilter != "" {
		q += " AND type = ?"
		args = append(args, typeFilter)
	}
	q += " ORDER BY created_at DESC"

	rows, err := db.Query(q, args...)
	if err != nil {
		h.ServerError(c, "查询失败: "+err.Error())
		return
	}
	defer rows.Close()

	var items []Log
	for rows.Next() {
		var l Log
		err := rows.Scan(&l.ID, &l.PlanName, &l.StockCode, &l.StockName, &l.Type, &l.TradingTime, &l.Price, &l.Quantity, &l.Strategy, &l.Remark, &l.CreatedAt, &l.UpdatedAt)
		if err != nil {
			h.ServerError(c, "数据解析失败: "+err.Error())
			return
		}
		items = append(items, l)
	}
	h.Success(c, gin.H{"list": items})
}

func (h *LogHandler) createLog(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	var l Log
	if err := c.ShouldBindJSON(&l); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	if l.ID == "" {
		l.ID = uuid.New().String()
	}

	_, err := db.Exec(`INSERT INTO logs (id, plan_name, stock_code, stock_name, type, trading_time, price, quantity, strategy, remark) VALUES (?,?,?,?,?,?,?,?,?,?)`,
		l.ID, l.PlanName, l.StockCode, l.StockName, l.Type, l.TradingTime, l.Price, l.Quantity, l.Strategy, l.Remark,
	)
	if err != nil {
		h.ServerError(c, "创建失败: "+err.Error())
		return
	}

	h.Success(c, l)
}

func (h *LogHandler) getLog(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	var l Log
	err := db.QueryRow(`SELECT id, plan_name, stock_code, stock_name, type, trading_time, price, quantity, strategy, remark, created_at, updated_at FROM logs WHERE id = ?`, id).
		Scan(&l.ID, &l.PlanName, &l.StockCode, &l.StockName, &l.Type, &l.TradingTime, &l.Price, &l.Quantity, &l.Strategy, &l.Remark, &l.CreatedAt, &l.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			h.NotFoundError(c, "交易日志不存在")
		} else {
			h.ServerError(c, "查询失败: "+err.Error())
		}
		return
	}

	h.Success(c, l)
}

func (h *LogHandler) updateLog(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	var l Log
	if err := c.ShouldBindJSON(&l); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	_, err := db.Exec(`UPDATE logs SET plan_name=?, stock_code=?, stock_name=?, type=?, trading_time=?, price=?, quantity=?, strategy=?, remark=? WHERE id=?`,
		l.PlanName, l.StockCode, l.StockName, l.Type, l.TradingTime, l.Price, l.Quantity, l.Strategy, l.Remark, id)
	if err != nil {
		h.ServerError(c, "更新失败: "+err.Error())
		return
	}

	h.Success(c, gin.H{"id": id})
}

func (h *LogHandler) deleteLog(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	_, err := db.Exec(`DELETE FROM logs WHERE id=?`, id)
	if err != nil {
		h.ServerError(c, "删除失败: "+err.Error())
		return
	}

	h.Success(c, gin.H{"id": id})
}
