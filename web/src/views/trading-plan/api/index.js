// 交易计划相关 API 服务
import { request } from '@/services/apiService'

// 交易计划相关 API
export const tradingPlanApi = {
  // 获取交易计划列表
  getPlans: (params = {}) => request.get('/trading-plans', params),
  
  // 获取交易计划详情
  getPlanDetail: (id) => request.get(`/trading-plans/${id}`),
  
  // 创建交易计划
  createPlan: (data) => request.post('/trading-plans', data),
  
  // 更新交易计划
  updatePlan: (id, data) => request.put(`/trading-plans/${id}`, data),
  
  // 删除交易计划
  deletePlan: (id) => request.delete(`/trading-plans/${id}`),
  
  // 更新计划状态
  updatePlanStatus: (id, status) => request.patch(`/trading-plans/${id}/status`, { status }),
  
  // 获取活跃计划
  getActivePlans: () => request.get('/trading-plans/active'),
  
  // 获取已完成计划
  getCompletedPlans: () => request.get('/trading-plans/completed'),
  
  // 批量更新计划状态
  batchUpdatePlanStatus: (ids, status) => request.patch('/trading-plans/batch-status', { ids, status }),
  
  // 复制交易计划
  copyPlan: (id) => request.post(`/trading-plans/${id}/copy`),
  
  // 导出交易计划
  exportPlans: (params = {}) => request.get('/trading-plans/export', params)
}

export default tradingPlanApi
