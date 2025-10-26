// 交易计划相关 API 服务
import { request } from '@/services/apiService'

// 交易计划相关 API
export const tradingPlanApi = {
  // 获取交易计划列表
  getPlans: (params = {}) => request.get('/plans/getList', params),
  
  // 获取交易计划详情
  getPlanDetail: (id) => request.get(`/plans/getDetail/${id}`),
  
  // 创建交易计划
  createPlan: (data) => request.post('/plans/create', data),
  
  // 更新交易计划
  updatePlan: (id, data) => request.put(`/plans/update/${id}`, data),
  
  // 删除交易计划
  deletePlan: (id) => request.delete(`/plans/delete/${id}`),
  
  // 更新计划状态
  updatePlanStatus: (id, status) => request.patch(`/plans/${id}/status`, { status }),
  
  // 获取活跃计划
  getActivePlans: () => request.get('/plans/active'),
  
  // 获取已完成计划
  getCompletedPlans: () => request.get('/plans/completed'),
  
  // 批量更新计划状态
  batchUpdatePlanStatus: (ids, status) => request.patch('/plans/batch-status', { ids, status }),
  
  // 复制交易计划
  copyPlan: (id) => request.post(`/plans/${id}/copy`),
  
  // 导出交易计划
  exportPlans: (params = {}) => request.get('/plans/export', params)
}

export default tradingPlanApi
