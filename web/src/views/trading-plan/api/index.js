// 交易计划相关 API 服务
import { request } from '@/services/api.js'

// 获取交易计划列表
export const getPlans = (params = {}) => request.get('/plans/getList', params)

// 获取交易计划详情
export const getPlanDetail = (id) => request.get(`/plans/getDetail/${id}`)

// 创建交易计划
export const createPlan = (data) => request.post('/plans/create', data)

// 更新交易计划
export const updatePlan = (id, data) => request.put(`/plans/update/${id}`, data)

// 删除交易计划
export const deletePlan = (id) => request.delete(`/plans/delete/${id}`)

// 更新计划状态
export const updatePlanStatus = (id, status) => request.patch(`/plans/${id}/status`, { status })

// 获取活跃计划
export const getActivePlans = () => request.get('/plans/active')

// 获取已完成计划
export const getCompletedPlans = () => request.get('/plans/completed')

// 批量更新计划状态
export const batchUpdatePlanStatus = (ids, status) => request.patch('/plans/batch-status', { ids, status })

// 复制交易计划
export const copyPlan = (id) => request.post(`/plans/${id}/copy`)

// 导出交易计划
export const exportPlans = (params = {}) => request.get('/plans/export', params)

// 默认导出
export default {
  getPlans,
  getPlanDetail,
  createPlan,
  updatePlan,
  deletePlan,
  updatePlanStatus,
  getActivePlans,
  getCompletedPlans,
  batchUpdatePlanStatus,
  copyPlan,
  exportPlans
}
