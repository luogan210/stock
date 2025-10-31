// 交易日志相关 API 服务
import { request } from '@/services/api.js'

// 获取交易日志列表
export const getLogs = (params = {}) => request.get('/logs/getList', params)

// 获取交易日志详情
export const getLogDetail = (id) => request.get(`/logs/getDetail/${id}`)

// 创建交易日志
export const createLog = (data) => request.post('/logs/create', data)

// 更新交易日志
export const updateLog = (id, data) => request.put(`/logs/update/${id}`, data)

// 删除交易日志
export const deleteLog = (id) => request.delete(`/logs/delete/${id}`)

// 默认导出
export default {
  getLogs,
  getLogDetail,
  createLog,
  updateLog,
  deleteLog
}
