package log

import (
	"database/sql"
	"fmt"
	"server/storage"
	"server/utils"
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
	utils.LogInfo("正在创建交易日志，股票代码: %s, 类型: %s", req.StockCode, req.Type)

	// 设置默认状态
	status := req.Status
	if status == "" {
		status = "pending"
	}

	log := &Log{
		ID:          id,
		Title:       req.Title,
		PlanName:    req.PlanName,
		StockCode:   req.StockCode,
		StockName:   req.StockName,
		Type:        req.Type,
		TradingTime: req.TradingTime,
		Price:       req.Price,
		Quantity:    req.Quantity,
		Strategy:    req.Strategy,
		Remark:      req.Remark,
		Status:      status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	query := `INSERT INTO logs (
		id, title, plan_name, stock_code, stock_name, type, trading_time,
		price, quantity, strategy, remark, status, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := storage.GetDB().Exec(query,
		log.ID, log.Title, log.PlanName, log.StockCode, log.StockName, log.Type,
		log.TradingTime, log.Price, log.Quantity, log.Strategy, log.Remark,
		log.Status, log.CreatedAt, log.UpdatedAt,
	)

	if err != nil {
		utils.LogError("创建交易日志失败，ID: %s, 错误: %v", id, err)
		return nil, fmt.Errorf("创建日志失败: %w", err)
	}

	utils.LogInfo("交易日志创建成功，ID: %s", id)
	return log, nil
}

// GetLog 获取日志详情
func (s *logService) GetLog(id string) (*Log, error) {
	utils.LogDebug("正在获取交易日志详情，ID: %s", id)
	query := `SELECT id, plan_name, stock_code, stock_name, type, trading_time,
		price, quantity, strategy, remark, created_at, updated_at, title, status
		FROM logs WHERE id = ?`

	log := &Log{}
	var title, planName, stockName, strategy, remark, status sql.NullString
	err := storage.GetDB().QueryRow(query, id).Scan(
		&log.ID, &planName, &log.StockCode, &stockName, &log.Type,
		&log.TradingTime, &log.Price, &log.Quantity, &strategy, &remark,
		&log.CreatedAt, &log.UpdatedAt, &title, &status,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			utils.LogWarning("交易日志不存在，ID: %s", id)
			return nil, fmt.Errorf("日志不存在")
		}
		utils.LogError("获取交易日志失败，ID: %s, 错误: %v", id, err)
		return nil, fmt.Errorf("获取日志失败: %w", err)
	}

	// 处理 NULL 值
	log.Title = title.String
	log.PlanName = planName.String
	log.StockName = stockName.String
	log.Strategy = strategy.String
	log.Remark = remark.String
	log.Status = status.String

	utils.LogDebug("成功获取交易日志详情，ID: %s", id)
	return log, nil
}

// ListLogs 获取日志列表
func (s *logService) ListLogs(req *LogListRequest) (*LogListResponse, error) {
	// 构建查询条件
	where := []string{"1=1"}
	args := []interface{}{}

	if req.Keyword != "" {
		where = append(where, "(title LIKE ? OR plan_name LIKE ? OR stock_name LIKE ? OR stock_code LIKE ?)")
		keyword := "%" + req.Keyword + "%"
		args = append(args, keyword, keyword, keyword, keyword)
	}

	if req.Type != "" {
		where = append(where, "type = ?")
		args = append(args, req.Type)
	}

	if req.Status != "" {
		where = append(where, "status = ?")
		args = append(args, req.Status)
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
		price, quantity, strategy, remark, created_at, updated_at, title, status
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
		var title, planName, stockName, strategy, remark, status sql.NullString
		err = rows.Scan(
			&log.ID, &planName, &log.StockCode, &stockName, &log.Type,
			&log.TradingTime, &log.Price, &log.Quantity, &strategy, &remark,
			&log.CreatedAt, &log.UpdatedAt, &title, &status,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描日志数据失败: %w", err)
		}

		// 处理 NULL 值
		log.Title = title.String
		log.PlanName = planName.String
		log.StockName = stockName.String
		log.Strategy = strategy.String
		log.Remark = remark.String
		log.Status = status.String

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
	utils.LogInfo("正在更新交易日志，ID: %s", id)
	// 检查日志是否存在
	_, err := s.GetLog(id)
	if err != nil {
		utils.LogWarning("更新交易日志失败，日志不存在，ID: %s", id)
		return nil, fmt.Errorf("日志不存在: %w", err)
	}

	// 构建更新字段
	setParts := []string{}
	args := []interface{}{}

	if req.Title != nil {
		setParts = append(setParts, "title = ?")
		args = append(args, *req.Title)
	}
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
	if req.Status != nil {
		setParts = append(setParts, "status = ?")
		args = append(args, *req.Status)
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
		utils.LogError("更新交易日志失败，ID: %s, 错误: %v", id, err)
		return nil, fmt.Errorf("更新日志失败: %w", err)
	}

	utils.LogInfo("交易日志更新成功，ID: %s", id)
	// 返回更新后的日志
	return s.GetLog(id)
}

// DeleteLog 删除日志
func (s *logService) DeleteLog(id string) error {
	utils.LogInfo("正在删除交易日志，ID: %s", id)
	// 检查日志是否存在
	_, err := s.GetLog(id)
	if err != nil {
		utils.LogWarning("删除交易日志失败，日志不存在，ID: %s", id)
		return fmt.Errorf("日志不存在: %w", err)
	}

	query := "DELETE FROM logs WHERE id = ?"
	result, err := storage.GetDB().Exec(query, id)
	if err != nil {
		utils.LogError("删除交易日志失败，ID: %s, 错误: %v", id, err)
		return fmt.Errorf("删除日志失败: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		utils.LogError("获取删除结果失败，ID: %s, 错误: %v", id, err)
		return fmt.Errorf("获取删除结果失败: %w", err)
	}

	if rowsAffected == 0 {
		utils.LogWarning("删除交易日志失败，未找到记录，ID: %s", id)
		return fmt.Errorf("日志不存在")
	}

	utils.LogInfo("交易日志删除成功，ID: %s", id)
	return nil
}
