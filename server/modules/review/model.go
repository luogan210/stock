package review

import "time"

// Review 交易复盘模型
type Review struct {
	ID           string    `json:"id" db:"id"`
	Period       string    `json:"period" db:"period"`
	ReviewDate   string    `json:"reviewDate" db:"review_date"`
	Title        string    `json:"title" db:"title"`
	BuyCount     int       `json:"buyCount" db:"buy_count"`
	SellCount    int       `json:"sellCount" db:"sell_count"`
	TotalProfit  float64   `json:"totalProfit" db:"total_profit"`
	Summary      string    `json:"summary" db:"summary"`
	Improvements string    `json:"improvements" db:"improvements"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt    time.Time `json:"updatedAt" db:"updated_at"`
}

// ReviewCreateRequest 创建复盘请求
type ReviewCreateRequest struct {
	Period       string  `json:"period" binding:"required"`
	ReviewDate   string  `json:"reviewDate" binding:"required"`
	Title        string  `json:"title" binding:"required"`
	BuyCount     int     `json:"buyCount"`
	SellCount    int     `json:"sellCount"`
	TotalProfit  float64 `json:"totalProfit"`
	Summary      string  `json:"summary" binding:"required"`
	Improvements string  `json:"improvements"`
}

// ReviewUpdateRequest 更新复盘请求
type ReviewUpdateRequest struct {
	Period       *string  `json:"period,omitempty"`
	ReviewDate   *string  `json:"reviewDate,omitempty"`
	Title        *string  `json:"title,omitempty"`
	BuyCount     *int     `json:"buyCount,omitempty"`
	SellCount    *int     `json:"sellCount,omitempty"`
	TotalProfit  *float64 `json:"totalProfit,omitempty"`
	Summary      *string  `json:"summary,omitempty"`
	Improvements *string  `json:"improvements,omitempty"`
}

// ReviewListRequest 复盘列表请求
type ReviewListRequest struct {
	Keyword   string `form:"keyword"`
	Period    string `form:"period"`
	StartDate string `form:"startDate"`
	EndDate   string `form:"endDate"`
	Page      int    `form:"page"`
	PageSize  int    `form:"pageSize"`
}

// ReviewListResponse 复盘列表响应
type ReviewListResponse struct {
	Items    []Review `json:"list"`
	Total    int      `json:"total"`
	Page     int      `json:"page"`
	PageSize int      `json:"pageSize"`
}
