// 交易日志相关 API 服务
import { request } from '@/services/apiService'

// 交易日志相关 API
export const tradingLogApi = {
  // 获取交易日志列表
  getLogs: (params = {}) => request.get('/trading-logs', params),
  
  // 获取交易日志详情
  getLogDetail: (id) => request.get(`/trading-logs/${id}`),
  
  // 创建交易日志
  createLog: (data) => request.post('/trading-logs', data),
  
  // 更新交易日志
  updateLog: (id, data) => request.put(`/trading-logs/${id}`, data),
  
  // 删除交易日志
  deleteLog: (id) => request.delete(`/trading-logs/${id}`),
  
  // 更新日志状态
  updateLogStatus: (id, status) => request.patch(`/trading-logs/${id}/status`, { status }),
  
  // 获取成功日志
  getSuccessLogs: () => request.get('/trading-logs/success'),
  
  // 获取失败日志
  getFailedLogs: () => request.get('/trading-logs/failed'),
  
  // 获取总盈亏
  getTotalProfit: () => request.get('/trading-logs/total-profit'),
  
  // 导出交易日志
  exportLogs: (params = {}) => request.get('/trading-logs/export', params),
  
  // 批量删除日志
  batchDeleteLogs: (ids) => request.delete('/trading-logs/batch', { data: { ids } }),
  
  // 获取日志统计
  getLogStatistics: (params = {}) => request.get('/trading-logs/statistics', params),
  
  // 复制交易日志
  copyLog: (id) => request.post(`/trading-logs/${id}/copy`)
}

export default tradingLogApi
