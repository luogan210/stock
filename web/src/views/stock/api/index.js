/**
 * 股票管理 API 服务
 * 真实环境下的 API 接口
 */

import { request } from '@/services/api.js'

// 股票管理相关 API
export const stockApi = {
  // 获取股票列表
  getStocks: (params = {}) => request.get('/stocks/getList', params),

  // 获取股票详情
  getStockDetail: (id) => request.get(`/stocks/getDetail/${id}`, { }),

  // 创建股票
  createStock: (data) => request.post('/stocks/create', data),

  // 更新股票
  updateStock: (id, data) => request.put(`/stocks/update/${id}`, data),

  // 删除股票
  deleteStock: (id) => request.delete(`/stocks/delete/${id}`)
}

export default stockApi