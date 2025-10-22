/**
 * 股票管理 API 服务
 * 真实环境下的 API 接口
 */

import { request } from '@/services/apiService'
import { mockStockApi } from './mock'

// 根据环境选择使用真实 API 还是模拟 API
const isDevelopment = process.env.NODE_ENV === 'development'
const useMockApi = false // 强制使用真实 API，不使用模拟数据

// 真实 API 服务
const realStockApi = {
  // 获取股票列表
  getStocks: (params = {}) => request.get('/stocks/getList', params),
  
  // 获取股票详情
  getStockDetail: (id) => request.get(`/stocks/getDetail/${id}`),
  
  // 创建股票
  createStock: (data) => request.post('/stocks/create', data),
  
  // 更新股票
  updateStock: (id, data) => request.put(`/stocks/update/${id}`, data),
  
  // 删除股票
  deleteStock: (id) => request.delete(`/stocks/delete/${id}`)
}

// 导出当前使用的 API
const currentApi = useMockApi ? mockStockApi : realStockApi

// 股票管理相关 API
export const stockApi = {
  // 获取股票列表
  getStocks: (params = {}) => currentApi.getStocks(params),
  
  // 获取股票详情
  getStockDetail: (id) => currentApi.getStockDetail(id),
  
  // 创建股票
  createStock: (data) => currentApi.createStock(data),
  
  // 更新股票
  updateStock: (id, data) => currentApi.updateStock(id, data),
  
  // 删除股票
  deleteStock: (id) => currentApi.deleteStock(id)
}

// 导出 API 方法（保持向后兼容）
export const getStockList = stockApi.getStocks
export const createStock = stockApi.createStock
export const updateStock = stockApi.updateStock
export const deleteStock = stockApi.deleteStock
export const getStockDetail = stockApi.getStockDetail

export default stockApi
