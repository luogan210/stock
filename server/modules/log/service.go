package log

import (
	"database/sql"
	"fmt"
	"server/storage"
	"strings"
	"time"
)

// LogService 日志服务接口
type LogService interface {
	CreateLog(req *LogCreateRequest) (*Log, error)
	GetLog(id string) (*Log, error)
	ListLogs(req *LogListRequest) (*LogListResponse, error)
	UpdateLog(id string, req *LogUpdateRequest) (*Log, error)
	DeleteLog(id string) error
}

// logService 日志服务实现
type logService struct{}

// NewLogService 创建日志服务
func NewLogService() LogService {
	return &logService{}
}

// CreateLog 创建日志
func (s *logService) CreateLog(req *LogCreateRequest) (*Log, error) {
	id := fmt.Sprintf("%d", time.Now().UnixNano())

	log := &Log{
		ID:          id,
		PlanName:    req.PlanName,
		StockCode:   req.StockCode,
		StockName:   req.StockName,
		Type:        req.Type,
		TradingTime: req.TradingTime,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Strategy:    req.Strategy,
		Remark:      req.Remark,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `INSERT INTO logs (
		id, plan_name, stock_code, stock_name, type, trading_time,
		price, quantity, strategy, remark, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := storage.GetDB().Exec(query,
		log.ID, log.PlanName, log.StockCode, log.StockName, log.Type,
		log.TradingTime, log.Price, log.Quantity, log.Strategy, log.Remark,
		log.CreatedAt, log.UpdatedAt,
	)

	if err != nil {
		return nil, fmt.Errorf("创建日志失败: %w", err)
	}

	return log, nil
}

// GetLog 获取日志详情
func (s *logService) GetLog(id string) (*Log, error) {
	query := `SELECT id, plan_name, stock_code, stock_name, type, trading_time,
		price, quantity, strategy, remark, created_at, updated_at
		FROM logs WHERE id = ?`

	log := &Log{}
	err := storage.GetDB().QueryRow(query, id).Scan(
		&log.ID, &log.PlanName, &log.StockCode, &log.StockName, &log.Type,
		&log.TradingTime, &log.Price, &log.Quantity, &log.Strategy, &log.Remark,
		&log.CreatedAt, &log.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("日志不存在")
		}
		return nil, fmt.Errorf("获取日志失败: %w", err)
	}

	return log, nil
}

// ListLogs 获取日志列表
func (s *logService) ListLogs(req *LogListRequest) (*LogListResponse, error) {
	// 构建查询条件
	where := []string{"1=1"}
	args := []interface{}{}

	if req.Keyword != "" {
		where = append(where, "(plan_name LIKE ? OR stock_name LIKE ? OR stock_code LIKE ?)")
		keyword := "%" + req.Keyword + "%"
		args = append(args, keyword, keyword, keyword)
	}

	if req.Type != "" {
		where = append(where, "type = ?")
		args = append(args, req.Type)
	}

	if req.StockCode != "" {
		where = append(where, "stock_code = ?")
		args = append(args, req.StockCode)
	}

	if req.PlanName != "" {
		where = append(where, "plan_name LIKE ?")
		args = append(args, "%"+req.PlanName+"%")
	}

	if req.StartDate != "" {
		where = append(where, "trading_time >= ?")
		args = append(args, req.StartDate)
	}

	if req.EndDate != "" {
		where = append(where, "trading_time <= ?")
		args = append(args, req.EndDate)
	}

	whereClause := strings.Join(where, " AND ")

	// 获取总数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM logs WHERE %s", whereClause)
	var total int
	err := storage.GetDB().QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, fmt.Errorf("获取日志总数失败: %w", err)
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
	query := fmt.Sprintf(`SELECT id, plan_name, stock_code, stock_name, type, trading_time,
		price, quantity, strategy, remark, created_at, updated_at
		FROM logs WHERE %s ORDER BY trading_time DESC LIMIT ? OFFSET ?`, whereClause)

	args = append(args, pageSize, offset)

	rows, err := storage.GetDB().Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("查询日志列表失败: %w", err)
	}
	defer rows.Close()

	var logs []Log
	for rows.Next() {
		log := &Log{}
		err = rows.Scan(
			&log.ID, &log.PlanName, &log.StockCode, &log.StockName, &log.Type,
			&log.TradingTime, &log.Price, &log.Quantity, &log.Strategy, &log.Remark,
			&log.CreatedAt, &log.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描日志数据失败: %w", err)
		}
		logs = append(logs, *log)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历日志数据失败: %w", err)
	}

	return &LogListResponse{
		Items:    logs,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// UpdateLog 更新日志
func (s *logService) UpdateLog(id string, req *LogUpdateRequest) (*Log, error) {
	// 检查日志是否存在
	_, err := s.GetLog(id)
	if err != nil {
		return nil, fmt.Errorf("日志不存在: %w", err)
	}

	// 构建更新字段
	setParts := []string{}
	args := []interface{}{}

	if req.PlanName != nil {
		setParts = append(setParts, "plan_name = ?")
		args = append(args, *req.PlanName)
	}
	if req.StockCode != nil {
		setParts = append(setParts, "stock_code = ?")
		args = append(args, *req.StockCode)
	}
	if req.StockName != nil {
		setParts = append(setParts, "stock_name = ?")
		args = append(args, *req.StockName)
	}
	if req.Type != nil {
		setParts = append(setParts, "type = ?")
		args = append(args, *req.Type)
	}
	if req.TradingTime != nil {
		setParts = append(setParts, "trading_time = ?")
		args = append(args, *req.TradingTime)
	}
	if req.Price != nil {
		setParts = append(setParts, "price = ?")
		args = append(args, *req.Price)
	}
	if req.Quantity != nil {
		setParts = append(setParts, "quantity = ?")
		args = append(args, *req.Quantity)
	}
	if req.Strategy != nil {
		setParts = append(setParts, "strategy = ?")
		args = append(args, *req.Strategy)
	}
	if req.Remark != nil {
		setParts = append(setParts, "remark = ?")
		args = append(args, *req.Remark)
	}

	if len(setParts) == 0 {
		return nil, fmt.Errorf("没有要更新的字段")
	}

	// 添加更新时间
	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())

	// 添加WHERE条件
	args = append(args, id)

	query := fmt.Sprintf("UPDATE logs SET %s WHERE id = ?", strings.Join(setParts, ", "))
	_, err = storage.GetDB().Exec(query, args...)
	if err != nil {
		return nil, fmt.Errorf("更新日志失败: %w", err)
	}

	// 返回更新后的日志
	return s.GetLog(id)
}

// DeleteLog 删除日志
func (s *logService) DeleteLog(id string) error {
	// 检查日志是否存在
	_, err := s.GetLog(id)
	if err != nil {
		return fmt.Errorf("日志不存在: %w", err)
	}

	query := "DELETE FROM logs WHERE id = ?"
	result, err := storage.GetDB().Exec(query, id)
	if err != nil {
		return fmt.Errorf("删除日志失败: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("获取删除结果失败: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("日志不存在")
	}

	return nil
}
