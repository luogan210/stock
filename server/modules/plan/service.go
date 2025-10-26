package plan

import (
	"fmt"
	"time"

	"server/utils"
)

// PlanService 计划服务接口
type PlanService interface {
	CreatePlan(req *PlanCreateRequest) (*Plan, error)
	GetPlan(id string) (*Plan, error)
	ListPlans(req *PlanListRequest) (*PlanListResponse, error)
	UpdatePlan(id string, req *PlanUpdateRequest) (*Plan, error)
	DeletePlan(id string) error
	UpdatePlanStatus(id string, status string) (*Plan, error)
}

// planService 计划服务实现
type planService struct {
	planRepo *PlanRepository
}

// NewPlanService 创建计划服务
func NewPlanService() PlanService {
	return &planService{
		planRepo: NewPlanRepository(),
	}
}

// CreatePlan 创建计划
func (s *planService) CreatePlan(req *PlanCreateRequest) (*Plan, error) {
	id := utils.GenerateID()

	plan := &Plan{
		ID:              id,
		Name:            req.Name,
		Type:            req.Type,
		StockCode:       req.StockCode,
		StockName:       req.StockName,
		Strategy:        req.Strategy,
		TradingStrategy: req.TradingStrategy,
		TargetPrice:     req.TargetPrice,
		Quantity:        req.Quantity,
		StopLoss:        req.StopLoss,
		TakeProfit:      req.TakeProfit,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		RiskLevel:       req.RiskLevel,
		Description:     req.Description,
		Remark:          req.Remark,
		Status:          "active",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	repo := NewPlanRepository()
	err := repo.Create(plan)
	if err != nil {
		return nil, fmt.Errorf("创建计划失败: %w", err)
	}

	return plan, nil
}

// GetPlan 获取计划详情
func (s *planService) GetPlan(id string) (*Plan, error) {
	repo := NewPlanRepository()
	plan, err := repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("获取计划失败: %w", err)
	}

	return plan, nil
}

// ListPlans 获取计划列表
func (s *planService) ListPlans(req *PlanListRequest) (*PlanListResponse, error) {
	repo := NewPlanRepository()
	plans, total, err := repo.List(req)
	if err != nil {
		return nil, fmt.Errorf("获取计划列表失败: %w", err)
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

	return &PlanListResponse{
		Items:    plans,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// UpdatePlan 更新计划
func (s *planService) UpdatePlan(id string, req *PlanUpdateRequest) (*Plan, error) {
	repo := NewPlanRepository()

	// 检查计划是否存在
	_, err := repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("计划不存在: %w", err)
	}

	// 更新计划
	err = repo.Update(id, req)
	if err != nil {
		return nil, fmt.Errorf("更新计划失败: %w", err)
	}

	// 返回更新后的计划
	updatedPlan, err := repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("获取更新后的计划失败: %w", err)
	}

	return updatedPlan, nil
}

// DeletePlan 删除计划
func (s *planService) DeletePlan(id string) error {
	repo := NewPlanRepository()

	// 检查计划是否存在
	_, err := repo.GetByID(id)
	if err != nil {
		return fmt.Errorf("计划不存在: %w", err)
	}

	// 删除计划
	err = repo.Delete(id)
	if err != nil {
		return fmt.Errorf("删除计划失败: %w", err)
	}

	return nil
}

// UpdatePlanStatus 更新计划状态
func (s *planService) UpdatePlanStatus(id string, status string) (*Plan, error) {
	req := &PlanUpdateRequest{Status: &status}
	return s.UpdatePlan(id, req)
}
