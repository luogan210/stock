package stock

import (
	"fmt"
	"server/utils"
	"time"
)

// StockService 股票服务接口
type StockService interface {
	CreateStock(req *StockCreateRequest) (*Stock, error)
	GetStock(id string) (*Stock, error)
	ListStocks(req *StockListRequest) (*StockListResponse, error)
	UpdateStock(id string, req *StockUpdateRequest) (*Stock, error)
	DeleteStock(id string) error
}

// stockService 股票服务实现
type stockService struct{}

// NewStockService 创建股票服务
func NewStockService() StockService {
	return &stockService{}
}

// CreateStock 创建股票
func (s *stockService) CreateStock(req *StockCreateRequest) (*Stock, error) {
	id := utils.GenerateID()

	stock := &Stock{
		ID:        id,
		Code:      req.Code,
		Name:      req.Name,
		Region:    req.Region,
		Currency:  req.Currency,
		Category:  req.Category,
		Enabled:   req.Enabled,
		Remark:    req.Remark,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	repo := NewStockRepository()
	err := repo.Create(stock)
	if err != nil {
		return nil, fmt.Errorf("创建股票失败: %w", err)
	}

	return stock, nil
}

// GetStock 获取股票详情
func (s *stockService) GetStock(id string) (*Stock, error) {
	repo := NewStockRepository()
	stock, err := repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("获取股票失败: %w", err)
	}

	return stock, nil
}

// ListStocks 获取股票列表
func (s *stockService) ListStocks(req *StockListRequest) (*StockListResponse, error) {
	repo := NewStockRepository()
	stocks, total, err := repo.List(req)
	if err != nil {
		return nil, fmt.Errorf("获取股票列表失败: %w", err)
	}

	// 设置默认分页参数
	page := req.Page
	if page <= 0 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize <= 0 {
		pageSize = 10
	}

	return &StockListResponse{
		Items:    stocks,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// UpdateStock 更新股票
func (s *stockService) UpdateStock(id string, req *StockUpdateRequest) (*Stock, error) {
	repo := NewStockRepository()

	// 检查股票是否存在
	_, err := repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("股票不存在: %w", err)
	}

	// 更新股票
	err = repo.Update(id, req)
	if err != nil {
		return nil, fmt.Errorf("更新股票失败: %w", err)
	}

	// 返回更新后的股票
	updatedStock, err := repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("获取更新后的股票失败: %w", err)
	}

	return updatedStock, nil
}

// DeleteStock 删除股票
func (s *stockService) DeleteStock(id string) error {
	repo := NewStockRepository()

	// 检查股票是否存在
	_, err := repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("股票不存在: %w", err)
	}

	// 删除股票
	err = repo.Delete(id)
	if err != nil {
		return fmt.Errorf("删除股票失败: %w", err)
	}

	return nil
}
