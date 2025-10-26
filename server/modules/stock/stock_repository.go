package stock

import (
	"database/sql"
	"fmt"
	"server/storage"
	"strings"
)

// StockRepository 股票数据访问层
type StockRepository struct{}

// NewStockRepository 创建股票仓库
func NewStockRepository() *StockRepository {
	return &StockRepository{}
}

// Create 创建股票
func (r *StockRepository) Create(stock *Stock) error {
	query := `INSERT INTO stocks (id, code, name, region, currency, category, enabled, remark, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := storage.GetDB().Exec(query,
		stock.ID, stock.Code, stock.Name, stock.Region, stock.Currency,
		stock.Category, stock.Enabled, stock.Remark, stock.CreatedAt, stock.UpdatedAt,
	)

	return err
}

// GetByID 根据ID获取股票
func (r *StockRepository) GetByID(id string) (*Stock, error) {
	query := `SELECT id, code, name, region, currency, category, enabled, remark, created_at, updated_at
		FROM stocks WHERE id = ?`

	stock := &Stock{}
	err := storage.GetDB().QueryRow(query, id).Scan(
		&stock.ID, &stock.Code, &stock.Name, &stock.Region, &stock.Currency,
		&stock.Category, &stock.Enabled, &stock.Remark, &stock.CreatedAt, &stock.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("股票不存在")
		}
		return nil, err
	}

	return stock, nil
}

// List 获取股票列表
func (r *StockRepository) List(req *StockListRequest) ([]Stock, int, error) {
	// 构建查询条件
	where := []string{"1=1"}
	args := []interface{}{}

	if req.Keyword != "" {
		where = append(where, "(code LIKE ? OR name LIKE ?)")
		keyword := "%" + req.Keyword + "%"
		args = append(args, keyword, keyword)
	}

	if req.Region != "" {
		where = append(where, "region = ?")
		args = append(args, req.Region)
	}

	if req.Category != "" {
		where = append(where, "category = ?")
		args = append(args, req.Category)
	}

	whereClause := strings.Join(where, " AND ")

	// 获取总数
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM stocks WHERE %s", whereClause)
	var total int
	err := storage.GetDB().QueryRow(countQuery, args...).Scan(&total)
	if err != nil {
		return nil, 0, err
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
	query := fmt.Sprintf(`SELECT id, code, name, region, currency, category, enabled, remark, created_at, updated_at
		FROM stocks WHERE %s ORDER BY created_at DESC LIMIT ? OFFSET ?`, whereClause)

	args = append(args, pageSize, offset)

	rows, err := storage.GetDB().Query(query, args...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var stocks []Stock
	for rows.Next() {
		stock := Stock{}
		err := rows.Scan(
			&stock.ID, &stock.Code, &stock.Name, &stock.Region, &stock.Currency,
			&stock.Category, &stock.Enabled, &stock.Remark, &stock.CreatedAt, &stock.UpdatedAt,
		)
		if err != nil {
			return nil, 0, err
		}
		stocks = append(stocks, stock)
	}

	// 检查遍历过程中是否有错误
	if err = rows.Err(); err != nil {
		return nil, 0, err
	}

	return stocks, total, nil
}

// Update 更新股票
func (r *StockRepository) Update(id string, req *StockUpdateRequest) error {
	// 构建更新字段
	setParts := []string{}
	args := []interface{}{}

	if req.Code != nil {
		setParts = append(setParts, "code = ?")
		args = append(args, *req.Code)
	}
	if req.Name != nil {
		setParts = append(setParts, "name = ?")
		args = append(args, *req.Name)
	}
	if req.Region != nil {
		setParts = append(setParts, "region = ?")
		args = append(args, *req.Region)
	}
	if req.Currency != nil {
		setParts = append(setParts, "currency = ?")
		args = append(args, *req.Currency)
	}
	if req.Category != nil {
		setParts = append(setParts, "category = ?")
		args = append(args, *req.Category)
	}
	if req.Enabled != nil {
		setParts = append(setParts, "enabled = ?")
		args = append(args, *req.Enabled)
	}
	if req.Remark != nil {
		setParts = append(setParts, "remark = ?")
		args = append(args, *req.Remark)
	}

	if len(setParts) == 0 {
		return fmt.Errorf("没有要更新的字段")
	}

	// 添加更新时间
	setParts = append(setParts, "updated_at = ?")
	args = append(args, "NOW()")

	// 添加WHERE条件
	args = append(args, id)

	query := fmt.Sprintf("UPDATE stocks SET %s WHERE id = ?", strings.Join(setParts, ", "))
	result, err := storage.GetDB().Exec(query, args...)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("股票不存在")
	}

	return nil
}

// Delete 删除股票
func (r *StockRepository) Delete(id string) error {
	query := "DELETE FROM stocks WHERE id = ?"
	result, err := storage.GetDB().Exec(query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("股票不存在")
	}

	return nil
}
