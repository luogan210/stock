// 交易日志相关 API 服务
import { request } from '@/services/apiService'

// 交易日志相关 API
export const tradingLogApi = {
  // 获取交易日志列表
  getLogs: (params = {}) => request.get('/logs/getList', params),
  
  // 获取交易日志详情
  getLogDetail: (id) => request.get(`/logs/getDetail/${id}`),
  
  // 创建交易日志
  createLog: (data) => request.post('/logs/create', data),
  
  // 更新交易日志
  updateLog: (id, data) => request.put(`/logs/update/${id}`, data),
  
  // 删除交易日志
  deleteLog: (id) => request.delete(`/logs/delete/${id}`)
}

export default tradingLogApi
