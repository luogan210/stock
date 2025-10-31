/**
 * 交易复盘 API 服务
 * 真实环境下的 API 接口
 */

import { request } from '@/services/api.js'

// 获取复盘列表
export const getReviews = (params = {}) => request.get('/reviews/getList', params)

// 获取复盘详情
export const getReview = (id) => request.get(`/reviews/getDetail/${id}`)

// 创建复盘
export const createReview = (data) => request.post('/reviews/create', data)

// 更新复盘
export const updateReview = (id, data) => request.put(`/reviews/update/${id}`, data)

// 删除复盘
export const deleteReview = (id) => request.delete(`/reviews/delete/${id}`)

// 默认导出
export default {
  getReviews,
  getReview,
  createReview,
  updateReview,
  deleteReview
}
