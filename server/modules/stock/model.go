package stock

import "time"

// Stock 股票模型
type Stock struct {
	ID        string    `json:"id" db:"id"`
	Code      string    `json:"code" db:"code"`
	Name      string    `json:"name" db:"name"`
	Region    string    `json:"region" db:"region"`
	Currency  string    `json:"currency" db:"currency"`
	Category  string    `json:"category" db:"category"`
	Enabled   bool      `json:"enabled" db:"enabled"`
	Remark    string    `json:"remark" db:"remark"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

// StockCreateRequest 创建股票请求
type StockCreateRequest struct {
	Code     string `json:"code" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Region   string `json:"region"`
	Currency string `json:"currency"`
	Category string `json:"category"`
	Enabled  bool   `json:"enabled"`
	Remark   string `json:"remark"`
}

// StockUpdateRequest 更新股票请求
type StockUpdateRequest struct {
	Code     *string `json:"code,omitempty"`
	Name     *string `json:"name,omitempty"`
	Region   *string `json:"region,omitempty"`
	Currency *string `json:"currency,omitempty"`
	Category *string `json:"category,omitempty"`
	Enabled  *bool   `json:"enabled,omitempty"`
	Remark   *string `json:"remark,omitempty"`
}

// StockListRequest 股票列表请求
type StockListRequest struct {
	Keyword  string `form:"keyword"`
	Region   string `form:"region"`
	Category string `form:"category"`
	Page     int    `form:"page"`
	PageSize int    `form:"pageSize"`
}

// StockListResponse 股票列表响应
type StockListResponse struct {
	Items    []Stock `json:"list"`
	Total    int     `json:"total"`
	Page     int     `json:"page"`
	PageSize int     `json:"pageSize"`
}
