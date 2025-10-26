package review

import (
	"database/sql"
	"fmt"
	"server/storage"
	"strings"
	"time"
)

// ReviewService 复盘服务接口
type ReviewService interface {
	CreateReview(req *ReviewCreateRequest) (*Review, error)
	GetReview(id string) (*Review, error)
	ListReviews(req *ReviewListRequest) (*ReviewListResponse, error)
	UpdateReview(id string, req *ReviewUpdateRequest) (*Review, error)
	DeleteReview(id string) error
}

// reviewService 复盘服务实现
type reviewService struct{}

// NewReviewService 创建复盘服务
func NewReviewService() ReviewService {
	return &reviewService{}
}

// CreateReview 创建复盘
func (s *reviewService) CreateReview(req *ReviewCreateRequest) (*Review, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	review := &Review{
		ID:           id,
		Period:       req.Period,
		ReviewDate:   req.ReviewDate,
		Title:        req.Title,
		BuyCount:     req.BuyCount,
		SellCount:    req.SellCount,
		TotalProfit:  req.TotalProfit,
		Summary:      req.Summary,
		Improvements: req.Improvements,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	query := `INSERT INTO reviews (
		id, period, review_date, title, buy_count, sell_count,
		total_profit, summary, improvements, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := storage.GetDB().Exec(query,
		review.ID, review.Period, review.ReviewDate, review.Title,
		review.BuyCount, review.SellCount, review.TotalProfit, review.Summary,
		review.Improvements, review.CreatedAt, review.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("创建复盘失败: %w", err)
	}

	return review, nil
}

// GetReview 获取复盘详情
func (s *reviewService) GetReview(id string) (*Review, error) {
	query := `SELECT id, period, review_date, title, buy_count, sell_count,
		total_profit, summary, improvements, created_at, updated_at
		FROM reviews WHERE id = ?`

	review := &Review{}
	err := storage.GetDB().QueryRow(query, id).Scan(
		&review.ID, &review.Period, &review.ReviewDate, &review.Title,
		&review.BuyCount, &review.SellCount, &review.TotalProfit, &review.Summary,
		&review.Improvements, &review.CreatedAt, &review.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("复盘不存在")
		}
		return nil, fmt.Errorf("获取复盘失败: %w", err)
	}

	return review, nil
}

// ListReviews 获取复盘列表
func (s *reviewService) ListReviews(req *ReviewListRequest) (*ReviewListResponse, error) {
	// 构建查询条件
	where := []string{"1=1"}
	args := []interface{}{}

	if req.Keyword != "" {
		where = append(where, "(title LIKE ? OR summary LIKE ?)")
		keyword := "%" + req.Keyword + "%"
		args = append(args, keyword, keyword)
	}

	if req.Period != "" {
		where = append(where, "period = ?")
		args = append(args, req.Period)
	}

	if req.StartDate != "" {
		where = append(where, "review_date >= ?")
		args = append(args, req.StartDate)
	}

	if req.EndDate != "" {
		where = append(where, "review_date <= ?")
		args = append(args, req.EndDate)
	}

	whereClause := strings.Join(where, " AND ")

	// 获取总数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM reviews WHERE %s", whereClause)
	var total int
	err := storage.GetDB().QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("获取复盘总数失败: %w", err)
	}

	// 设置分页参数
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	// 获取数据
	query := fmt.Sprintf(`SELECT id, period, review_date, title, buy_count, sell_count,
		total_profit, summary, improvements, created_at, updated_at
		FROM reviews WHERE %s ORDER BY review_date DESC LIMIT ? OFFSET ?`, whereClause)

	args = append(args, pageSize, offset)

	rows, err := storage.GetDB().Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("查询复盘列表失败: %w", err)
	}
	defer rows.Close()

	var reviews []Review
	for rows.Next() {
		review := &Review{}
		err = rows.Scan(
			&review.ID, &review.Period, &review.ReviewDate, &review.Title,
			&review.BuyCount, &review.SellCount, &review.TotalProfit, &review.Summary,
			&review.Improvements, &review.CreatedAt, &review.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描复盘数据失败: %w", err)
		}
		reviews = append(reviews, *review)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历复盘数据失败: %w", err)
	}

	return &ReviewListResponse{
		Items:    reviews,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// UpdateReview 更新复盘
func (s *reviewService) UpdateReview(id string, req *ReviewUpdateRequest) (*Review, error) {
	// 检查复盘是否存在
	_, err := s.GetReview(id)
	if err != nil {
		return nil, fmt.Errorf("复盘不存在: %w", err)
	}

	// 构建更新字段
	setParts := []string{}
	args := []interface{}{}

	if req.Period != nil {
		setParts = append(setParts, "period = ?")
		args = append(args, *req.Period)
	}
	if req.ReviewDate != nil {
		setParts = append(setParts, "review_date = ?")
		args = append(args, *req.ReviewDate)
	}
	if req.Title != nil {
		setParts = append(setParts, "title = ?")
		args = append(args, *req.Title)
	}
	if req.BuyCount != nil {
		setParts = append(setParts, "buy_count = ?")
		args = append(args, *req.BuyCount)
	}
	if req.SellCount != nil {
		setParts = append(setParts, "sell_count = ?")
		args = append(args, *req.SellCount)
	}
	if req.TotalProfit != nil {
		setParts = append(setParts, "total_profit = ?")
		args = append(args, *req.TotalProfit)
	}
	if req.Summary != nil {
		setParts = append(setParts, "summary = ?")
		args = append(args, *req.Summary)
	}
	if req.Improvements != nil {
		setParts = append(setParts, "improvements = ?")
		args = append(args, *req.Improvements)
	}

	if len(setParts) == 0 {
		return nil, fmt.Errorf("没有要更新的字段")
	}

	// 添加更新时间
	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())

	// 添加WHERE条件
	args = append(args, id)

	query := fmt.Sprintf("UPDATE reviews SET %s WHERE id = ?", strings.Join(setParts, ", "))
	_, err = storage.GetDB().Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("更新复盘失败: %w", err)
	}

	// 返回更新后的复盘
	return s.GetReview(id)
}

// DeleteReview 删除复盘
func (s *reviewService) DeleteReview(id string) error {
	// 检查复盘是否存在
	_, err := s.GetReview(id)
	if err != nil {
		return fmt.Errorf("复盘不存在: %w", err)
	}

	query := "DELETE FROM reviews WHERE id = ?"
	result, err := storage.GetDB().Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除复盘失败: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取删除结果失败: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("复盘不存在")
	}

	return nil
}
