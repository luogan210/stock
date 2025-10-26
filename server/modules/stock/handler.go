package stock

import (
	"strconv"

	"server/handler"

	"github.com/gin-gonic/gin"
)

// StockHandler 股票处理器
type StockHandler struct {
	stockService StockService
}

// NewStockHandler 创建股票处理器
func NewStockHandler() *StockHandler {
	return &StockHandler{
		stockService: NewStockService(),
	}
}

// RegisterStockRoutes 注册股票路由
func RegisterStockRoutes(r *gin.RouterGroup) {
	handler := NewStockHandler()

	g := r.Group("/stocks")
	{
		g.POST("/create", handler.createStock)
		g.GET("/getList", handler.listStocks)
		g.GET("/getDetail/:id", handler.getStock)
		g.PUT("/update/:id", handler.updateStock)
		g.DELETE("/delete/:id", handler.deleteStock)
	}
}

// CreateStock 创建股票
func (h *StockHandler) createStock(c *gin.Context) {
	var req StockCreateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
		return
	}

	stock, err := h.stockService.CreateStock(&req)
	if err != nil {
		handler.Error(c, handler.CodeError, err.Error())
		return
	}

	handler.Success(c, gin.H{"id": stock.ID})
}

// GetStock 获取股票详情
func (h *StockHandler) getStock(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handler.Error(c, handler.CodeInvalid, "股票ID不能为空")
		return
	}

	stock, err := h.stockService.GetStock(id)
	if err != nil {
		handler.Error(c, handler.CodeNotFound, err.Error())
		return
	}

	handler.Success(c, stock)
}

// ListStocks 获取股票列表
func (h *StockHandler) listStocks(c *gin.Context) {
	req := &StockListRequest{
		Keyword:  c.Query("keyword"),
		Region:   c.Query("region"),
		Category: c.Query("category"),
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

	response, err := h.stockService.ListStocks(req)
	if err != nil {
		handler.Error(c, handler.CodeError, err.Error())
		return
	}

	handler.Success(c, response)
}

// UpdateStock 更新股票
func (h *StockHandler) updateStock(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handler.Error(c, handler.CodeInvalid, "股票ID不能为空")
		return
	}

	var req StockUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.Error(c, handler.CodeInvalid, "参数错误: "+err.Error())
		return
	}

	stock, err := h.stockService.UpdateStock(id, &req)
	if err != nil {
		handler.Error(c, handler.CodeError, err.Error())
		return
	}

	handler.Success(c, stock)
}

// DeleteStock 删除股票
func (h *StockHandler) deleteStock(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		handler.Error(c, handler.CodeInvalid, "股票ID不能为空")
		return
	}

	err := h.stockService.DeleteStock(id)
	if err != nil {
		handler.Error(c, handler.CodeError, err.Error())
		return
	}

	handler.Success(c, gin.H{"id": id})
}
