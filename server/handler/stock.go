package handler

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Stock struct {
	ID       string `json:"id"`
	Code     string `json:"code" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Region   string `json:"region"`
	Currency string `json:"currency"`
	Category string `json:"category"`
	Enabled  bool   `json:"enabled"`
	Remark   string `json:"remark"`
}

// generateID 生成唯一ID
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}

// In a fuller setup, the DB would be injected. For a minimal version, get it from context.
func getDB(c *gin.Context) *sql.DB {
	v, ok := c.Get("db")
	if !ok {
		return nil
	}
	db, _ := v.(*sql.DB)
	return db
}

func RegisterStockRoutes(r *gin.RouterGroup) {
	g := r.Group("/stocks")
	g.GET("/getList", listStocks)
	g.GET("/getDetail/:id", getStock)
	g.POST("/create", createStock)
	g.PUT("/update/:id", updateStock)
	g.DELETE("/delete/:id", deleteStock)
}

func listStocks(c *gin.Context) {
	db := getDB(c)
	if db == nil {
		Error(c, CodeError, "数据库未就绪")
		return
	}

	keyword := strings.TrimSpace(c.Query("keyword"))
	region := c.Query("region")
	category := c.Query("category")

	q := `SELECT id, code, name, region, currency, category, enabled, remark FROM stocks WHERE 1=1`
	args := []any{}
	if keyword != "" {
		q += " AND (code LIKE ? OR name LIKE ?)"
		like := "%" + keyword + "%"
		args = append(args, like, like)
	}
	if region != "" {
		q += " AND region = ?"
		args = append(args, region)
	}
	if category != "" {
		q += " AND category = ?"
		args = append(args, category)
	}

	rows, err := db.Query(q, args...)
	if err != nil {
		Error(c, CodeError, "查询失败: "+err.Error())
		return
	}
	defer rows.Close()

	var items []Stock
	for rows.Next() {
		var s Stock
		var enabled int
		if err := rows.Scan(&s.ID, &s.Code, &s.Name, &s.Region, &s.Currency, &s.Category, &enabled, &s.Remark); err != nil {
			Error(c, CodeError, "数据解析失败: "+err.Error())
			return
		}
		s.Enabled = enabled == 1
		items = append(items, s)
	}
	Success(c, gin.H{"list": items})
}

func createStock(c *gin.Context) {
	db := getDB(c)
	if db == nil {
		Error(c, CodeError, "数据库未就绪")
		return
	}
	var s Stock
	if err := c.ShouldBindJSON(&s); err != nil {
		Error(c, CodeInvalid, "参数错误: "+err.Error())
		return
	}
	if s.ID == "" {
		s.ID = generateID()
	}
	enabled := 0
	if s.Enabled {
		enabled = 1
	}
	
	_, err := db.Exec(`INSERT INTO stocks (id, code, name, region, currency, category, enabled, remark) VALUES (?,?,?,?,?,?,?,?)`,
		s.ID, s.Code, s.Name, s.Region, s.Currency, s.Category, enabled, s.Remark,
	)
	if err != nil {
		Error(c, CodeError, "创建失败: "+err.Error())
		return
	}
	Success(c, gin.H{"id": s.ID})
}

func getStock(c *gin.Context) {
	db := getDB(c)
	if db == nil {
		Error(c, CodeError, "数据库未就绪")
		return
	}

	id := c.Param("id")
	var s Stock
	var enabled int
	err := db.QueryRow(`SELECT id, code, name, region, currency, category, enabled, remark FROM stocks WHERE id = ?`, id).
		Scan(&s.ID, &s.Code, &s.Name, &s.Region, &s.Currency, &s.Category, &enabled, &s.Remark)
	if err != nil {
		if err == sql.ErrNoRows {
			Error(c, CodeNotFound, "股票不存在")
		} else {
			Error(c, CodeError, "查询失败: "+err.Error())
		}
		return
	}
	s.Enabled = enabled == 1
	Success(c, s)
}

func updateStock(c *gin.Context) {
	db := getDB(c)
	if db == nil {
		Error(c, CodeError, "数据库未就绪")
		return
	}

	id := c.Param("id")
	var s Stock
	if err := c.ShouldBindJSON(&s); err != nil {
		Error(c, CodeInvalid, "参数错误: "+err.Error())
		return
	}
	s.ID = id
	enabled := 0
	if s.Enabled {
		enabled = 1
	}
	_, err := db.Exec(`UPDATE stocks SET code=?, name=?, region=?, currency=?, category=?, enabled=?, remark=? WHERE id=?`,
		s.Code, s.Name, s.Region, s.Currency, s.Category, enabled, s.Remark, s.ID)
	if err != nil {
		Error(c, CodeError, "更新失败: "+err.Error())
		return
	}
	Success(c, gin.H{"id": s.ID})
}

func deleteStock(c *gin.Context) {
	db := getDB(c)
	if db == nil {
		Error(c, CodeError, "数据库未就绪")
		return
	}

	id := c.Param("id")
	_, err := db.Exec(`DELETE FROM stocks WHERE id = ?`, id)
	if err != nil {
		Error(c, CodeError, "删除失败: "+err.Error())
		return
	}
	Success(c, gin.H{"id": id})
}
