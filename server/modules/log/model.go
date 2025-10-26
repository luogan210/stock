package log

import "time"

// Log 交易日志模型
type Log struct {
	ID          string    `json:"id" db:"id"`
	PlanName    string    `json:"planName" db:"plan_name"`
	StockCode   string    `json:"stockCode" db:"stock_code"`
	StockName   string    `json:"stockName" db:"stock_name"`
	Type        string    `json:"type" db:"type"`
	TradingTime string    `json:"tradingTime" db:"trading_time"`
	Price       float64   `json:"price" db:"price"`
	Quantity    int       `json:"quantity" db:"quantity"`
	Strategy    string    `json:"strategy" db:"strategy"`
	Remark      string    `json:"remark" db:"remark"`
	CreatedAt   time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt   time.Time `json:"updatedAt" db:"updated_at"`
}

// LogCreateRequest 创建日志请求
type LogCreateRequest struct {
	PlanName    string  `json:"planName"`
	StockCode   string  `json:"stockCode" binding:"required"`
	StockName   string  `json:"stockName"`
	Type        string  `json:"type" binding:"required"`
	TradingTime string  `json:"tradingTime" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required"`
	Strategy    string  `json:"strategy"`
	Remark      string  `json:"remark"`
}

// LogUpdateRequest 更新日志请求
type LogUpdateRequest struct {
	PlanName    *string  `json:"planName,omitempty"`
	StockCode   *string  `json:"stockCode,omitempty"`
	StockName   *string  `json:"stockName,omitempty"`
	Type        *string  `json:"type,omitempty"`
	TradingTime *string  `json:"tradingTime,omitempty"`
	Price       *float64 `json:"price,omitempty"`
	Quantity    *int     `json:"quantity,omitempty"`
	Strategy    *string  `json:"strategy,omitempty"`
	Remark      *string  `json:"remark,omitempty"`
}

// LogListRequest 日志列表请求
type LogListRequest struct {
	Keyword   string `form:"keyword"`
	Type      string `form:"type"`
	StockCode string `form:"stockCode"`
	PlanName  string `form:"planName"`
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
	Page      int    `form:"page"`
	PageSize  int    `form:"pageSize"`
}

// LogListResponse 日志列表响应
type LogListResponse struct {
	Items    []Log `json:"list"`
	Total    int   `json:"total"`
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
}
