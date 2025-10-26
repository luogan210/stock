package plan

import (
	"database/sql"
	"fmt"
	"server/storage"
	"strings"
	"time"
)

// PlanRepository 计划数据访问层
type PlanRepository struct{}

// NewPlanRepository 创建计划仓库
func NewPlanRepository() *PlanRepository {
	return &PlanRepository{}
}

// Create 创建计划
func (r *PlanRepository) Create(plan *Plan) error {
	query := `INSERT INTO plans (
		id, name, type, stock_code, stock_name, strategy, trading_strategy,
		target_price, quantity, stop_loss, take_profit, start_time, end_time,
		risk_level, description, remark, status, created_at, updated_at
	) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := storage.GetDB().Exec(query,
		plan.ID, plan.Name, plan.Type, plan.StockCode, plan.StockName,
		plan.Strategy, plan.TradingStrategy, plan.TargetPrice, plan.Quantity,
		plan.StopLoss, plan.TakeProfit, plan.StartTime, plan.EndTime,
		plan.RiskLevel, plan.Description, plan.Remark, plan.Status,
		plan.CreatedAt, plan.UpdatedAt,
	)

	return err
}

// GetByID 根据ID获取计划
func (r *PlanRepository) GetByID(id string) (*Plan, error) {
	query := `SELECT id, name, type, stock_code, stock_name, strategy, trading_strategy,
		target_price, quantity, stop_loss, take_profit, start_time, end_time,
		risk_level, description, remark, status, created_at, updated_at
		FROM plans WHERE id = ?`

	plan := &Plan{}
	err := storage.GetDB().QueryRow(query, id).Scan(
		&plan.ID, &plan.Name, &plan.Type, &plan.StockCode, &plan.StockName,
		&plan.Strategy, &plan.TradingStrategy, &plan.TargetPrice, &plan.Quantity,
		&plan.StopLoss, &plan.TakeProfit, &plan.StartTime, &plan.EndTime,
		&plan.RiskLevel, &plan.Description, &plan.Remark, &plan.Status,
		&plan.CreatedAt, &plan.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("计划不存在")
		}
		return nil, err
	}
	return plan, nil
}

// List 获取计划列表
func (r *PlanRepository) List(req *PlanListRequest) ([]Plan, int, error) {
	// 构建查询条件
	where := []string{"1=1"}
	args := []interface{}{}

	if req.Keyword != "" {
		where = append(where, "(name LIKE ? OR stock_name LIKE ? OR stock_code LIKE ?)")
		keyword := "%" + req.Keyword + "%"
		args = append(args, keyword, keyword, keyword)
	}
	if req.Type != "" {
		where = append(where, "type = ?")
		args = append(args, req.Type)
	}
	if req.Status != "" {
		where = append(where, "status = ?")
		args = append(args, req.Status)
	}
	if req.RiskLevel != "" {
		where = append(where, "risk_level = ?")
		args = append(args, req.RiskLevel)
	}
	if req.StockCode != "" {
		where = append(where, "stock_code = ?")
		args = append(args, req.StockCode)
	}

	whereClause := strings.Join(where, " AND ")
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM plans WHERE %s", whereClause)
	var total int
	err := storage.GetDB().QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	query := fmt.Sprintf(`SELECT id, name, type, stock_code, stock_name, strategy, trading_strategy,
		target_price, quantity, stop_loss, take_profit, start_time, end_time,
		risk_level, description, remark, status, created_at, updated_at
		FROM plans WHERE %s ORDER BY created_at DESC LIMIT ? OFFSET ?`, whereClause)

	args = append(args, pageSize, offset)
	rows, err := storage.GetDB().Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var plans []Plan
	for rows.Next() {
		plan := Plan{}
		err := rows.Scan(
			&plan.ID, &plan.Name, &plan.Type, &plan.StockCode, &plan.StockName,
			&plan.Strategy, &plan.TradingStrategy, &plan.TargetPrice, &plan.Quantity,
			&plan.StopLoss, &plan.TakeProfit, &plan.StartTime, &plan.EndTime,
			&plan.RiskLevel, &plan.Description, &plan.Remark, &plan.Status,
			&plan.CreatedAt, &plan.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		plans = append(plans, plan)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return plans, total, nil
}

// Update 更新计划
func (r *PlanRepository) Update(id string, req *PlanUpdateRequest) error {
	// 构建更新字段
	setParts := []string{}
	args := []interface{}{}

	if req.Name != nil {
		setParts = append(setParts, "name = ?")
		args = append(args, *req.Name)
	}
	if req.Type != nil {
		setParts = append(setParts, "type = ?")
		args = append(args, *req.Type)
	}
	if req.StockCode != nil {
		setParts = append(setParts, "stock_code = ?")
		args = append(args, *req.StockCode)
	}
	if req.StockName != nil {
		setParts = append(setParts, "stock_name = ?")
		args = append(args, *req.StockName)
	}
	if req.Strategy != nil {
		setParts = append(setParts, "strategy = ?")
		args = append(args, *req.Strategy)
	}
	if req.TradingStrategy != nil {
		setParts = append(setParts, "trading_strategy = ?")
		args = append(args, *req.TradingStrategy)
	}
	if req.TargetPrice != nil {
		setParts = append(setParts, "target_price = ?")
		args = append(args, *req.TargetPrice)
	}
	if req.Quantity != nil {
		setParts = append(setParts, "quantity = ?")
		args = append(args, *req.Quantity)
	}
	if req.StopLoss != nil {
		setParts = append(setParts, "stop_loss = ?")
		args = append(args, *req.StopLoss)
	}
	if req.TakeProfit != nil {
		setParts = append(setParts, "take_profit = ?")
		args = append(args, *req.TakeProfit)
	}
	if req.StartTime != nil {
		setParts = append(setParts, "start_time = ?")
		args = append(args, *req.StartTime)
	}
	if req.EndTime != nil {
		setParts = append(setParts, "end_time = ?")
		args = append(args, *req.EndTime)
	}
	if req.RiskLevel != nil {
		setParts = append(setParts, "risk_level = ?")
		args = append(args, *req.RiskLevel)
	}
	if req.Description != nil {
		setParts = append(setParts, "description = ?")
		args = append(args, *req.Description)
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
		return fmt.Errorf("没有要更新的字段")
	}

	setParts = append(setParts, "updated_at = ?")
	args = append(args, time.Now())
	args = append(args, id)

	query := fmt.Sprintf("UPDATE plans SET %s WHERE id = ?", strings.Join(setParts, ", "))
	result, err := storage.GetDB().Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("计划不存在")
	}
	return nil
}

// Delete 删除计划
func (r *PlanRepository) Delete(id string) error {
	query := "DELETE FROM plans WHERE id = ?"
	result, err := storage.GetDB().Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("计划不存在")
	}
	return nil
}
