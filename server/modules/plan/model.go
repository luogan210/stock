package plan

import "time"

// Plan 交易计划模型
type Plan struct {
	ID              string    `json:"id" db:"id"`
	Name            string    `json:"name" db:"name"`
	Type            string    `json:"type" db:"type"`
	StockCode       string    `json:"stockCode" db:"stock_code"`
	StockName       string    `json:"stockName" db:"stock_name"`
	Strategy        string    `json:"strategy" db:"strategy"`
	TradingStrategy string    `json:"tradingStrategy" db:"trading_strategy"`
	TargetPrice     float64   `json:"targetPrice" db:"target_price"`
	Quantity        int       `json:"quantity" db:"quantity"`
	StopLoss        float64   `json:"stopLoss" db:"stop_loss"`
	TakeProfit      float64   `json:"takeProfit" db:"take_profit"`
	StartTime       string    `json:"startTime" db:"start_time"`
	EndTime         string    `json:"endTime" db:"end_time"`
	RiskLevel       string    `json:"riskLevel" db:"risk_level"`
	Description     string    `json:"description" db:"description"`
	Remark          string    `json:"remark" db:"remark"`
	Status          string    `json:"status" db:"status"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt       time.Time `json:"updatedAt" db:"updated_at"`
}

// PlanCreateRequest 创建计划请求
type PlanCreateRequest struct {
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
}

// PlanUpdateRequest 更新计划请求
type PlanUpdateRequest struct {
	Name            *string  `json:"name,omitempty"`
	Type            *string  `json:"type,omitempty"`
	StockCode       *string  `json:"stockCode,omitempty"`
	StockName       *string  `json:"stockName,omitempty"`
	Strategy        *string  `json:"strategy,omitempty"`
	TradingStrategy *string  `json:"tradingStrategy,omitempty"`
	TargetPrice     *float64 `json:"targetPrice,omitempty"`
	Quantity        *int     `json:"quantity,omitempty"`
	StopLoss        *float64 `json:"stopLoss,omitempty"`
	TakeProfit      *float64 `json:"takeProfit,omitempty"`
	StartTime       *string  `json:"startTime,omitempty"`
	EndTime         *string  `json:"endTime,omitempty"`
	RiskLevel       *string  `json:"riskLevel,omitempty"`
	Description     *string  `json:"description,omitempty"`
	Remark          *string  `json:"remark,omitempty"`
	Status          *string  `json:"status,omitempty"`
}

// PlanListRequest 计划列表请求
type PlanListRequest struct {
	Keyword   string `form:"keyword"`
	Type      string `form:"type"`
	Status    string `form:"status"`
	RiskLevel string `form:"riskLevel"`
	StockCode string `form:"stockCode"`
	Page      int    `form:"page"`
	PageSize  int    `form:"pageSize"`
}

// PlanListResponse 计划列表响应
type PlanListResponse struct {
	Items    []Plan `json:"list"`
	Total    int    `json:"total"`
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
}
