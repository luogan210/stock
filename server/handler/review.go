package handler

import (
	"database/sql"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ReviewHandler 交易复盘处理
type ReviewHandler struct {
	*BaseHandler
}

// NewReviewHandler 创建交易复盘处理器实例
func NewReviewHandler() *ReviewHandler {
	return &ReviewHandler{
		BaseHandler: NewBaseHandler(),
	}
}

// Review 交易复盘模型
type Review struct {
	ID           string  `json:"id"`
	Period       string  `json:"period" binding:"required"`
	ReviewDate   string  `json:"reviewDate" binding:"required"`
	Title        string  `json:"title" binding:"required"`
	BuyCount     int     `json:"buyCount"`
	SellCount    int     `json:"sellCount"`
	TotalProfit  float64 `json:"totalProfit"`
	Summary      string  `json:"summary" binding:"required"`
	Improvements string  `json:"improvements"`
	CreatedAt    string  `json:"createdAt"`
	UpdatedAt    string  `json:"updatedAt"`
}

// RegisterReviewRoutes 注册交易复盘相关路由
func RegisterReviewRoutes(rg *gin.RouterGroup) {
	h := NewReviewHandler()
	g := rg.Group("/reviews")
	g.GET("/getList", h.listReviews)        // GET /api/reviews/getList - 获取复盘列表
	g.POST("/create", h.createReview)       // POST /api/reviews/create - 创建复盘
	g.GET("/getDetail/:id", h.getReview)    // GET /api/reviews/getDetail/:id - 获取复盘详情
	g.PUT("/update/:id", h.updateReview)    // PUT /api/reviews/update/:id - 更新复盘
	g.DELETE("/delete/:id", h.deleteReview) // DELETE /api/reviews/delete/:id - 删除复盘
}

func (h *ReviewHandler) getDB(c *gin.Context) *sql.DB {
	v, ok := c.Get("db")
	if !ok {
		return nil
	}
	db, _ := v.(*sql.DB)
	return db
}

func (h *ReviewHandler) listReviews(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	keyword := strings.TrimSpace(c.Query("keyword"))
	period := c.Query("period")

	q := `SELECT id, period, review_date, title, buy_count, sell_count, total_profit, summary, improvements, created_at, updated_at FROM reviews WHERE 1=1`
	args := []any{}
	if keyword != "" {
		q += " AND (title LIKE ? OR summary LIKE ?)"
		like := "%" + keyword + "%"
		args = append(args, like, like)
	}
	if period != "" {
		q += " AND period = ?"
		args = append(args, period)
	}
	q += " ORDER BY created_at DESC"

	rows, err := db.Query(q, args...)
	if err != nil {
		h.ServerError(c, "查询失败: "+err.Error())
		return
	}
	defer rows.Close()

	var items []Review
	for rows.Next() {
		var r Review
		err := rows.Scan(&r.ID, &r.Period, &r.ReviewDate, &r.Title, &r.BuyCount, &r.SellCount, &r.TotalProfit, &r.Summary, &r.Improvements, &r.CreatedAt, &r.UpdatedAt)
		if err != nil {
			h.ServerError(c, "数据解析失败: "+err.Error())
			return
		}
		items = append(items, r)
	}
	h.Success(c, gin.H{"list": items})
}

func (h *ReviewHandler) createReview(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	var r Review
	if err := c.ShouldBindJSON(&r); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	if r.ID == "" {
		r.ID = uuid.New().String()
	}

	_, err := db.Exec(`INSERT INTO reviews (id, period, review_date, title, buy_count, sell_count, total_profit, summary, improvements) VALUES (?,?,?,?,?,?,?,?,?)`,
		r.ID, r.Period, r.ReviewDate, r.Title, r.BuyCount, r.SellCount, r.TotalProfit, r.Summary, r.Improvements,
	)
	if err != nil {
		h.ServerError(c, "创建失败: "+err.Error())
		return
	}

	h.Success(c, r)
}

func (h *ReviewHandler) getReview(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	var r Review
	err := db.QueryRow(`SELECT id, period, review_date, title, buy_count, sell_count, total_profit, summary, improvements, created_at, updated_at FROM reviews WHERE id = ?`, id).
		Scan(&r.ID, &r.Period, &r.ReviewDate, &r.Title, &r.BuyCount, &r.SellCount, &r.TotalProfit, &r.Summary, &r.Improvements, &r.CreatedAt, &r.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			h.NotFoundError(c, "交易复盘不存在")
		} else {
			h.ServerError(c, "查询失败: "+err.Error())
		}
		return
	}

	h.Success(c, r)
}

func (h *ReviewHandler) updateReview(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	var r Review
	if err := c.ShouldBindJSON(&r); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	_, err := db.Exec(`UPDATE reviews SET period=?, review_date=?, title=?, buy_count=?, sell_count=?, total_profit=?, summary=?, improvements=? WHERE id=?`,
		r.Period, r.ReviewDate, r.Title, r.BuyCount, r.SellCount, r.TotalProfit, r.Summary, r.Improvements, id)
	if err != nil {
		h.ServerError(c, "更新失败: "+err.Error())
		return
	}

	h.Success(c, gin.H{"id": id})
}

func (h *ReviewHandler) deleteReview(c *gin.Context) {
	db := h.getDB(c)
	if db == nil {
		h.ServerError(c, "数据库未就绪")
		return
	}

	id := c.Param("id")
	_, err := db.Exec(`DELETE FROM reviews WHERE id=?`, id)
	if err != nil {
		h.ServerError(c, "删除失败: "+err.Error())
		return
	}

	h.Success(c, gin.H{"id": id})
}
