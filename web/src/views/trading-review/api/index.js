/**
 * 交易复盘 API 服务
 * 真实环境下的 API 接口
 */

import { request } from '@/services/apiService'

export const tradingReviewApi = {
  // 获取复盘列表
  getReviews: (params = {}) => {
    return request.get('/trading-review', params)
  },

  // 获取复盘详情
  getReview: (id) => {
    return request.get(`/trading-review/${id}`)
  },

  // 创建复盘
  createReview: (data) => {
    return request.post('/trading-review', data)
  },

  // 更新复盘
  updateReview: (id, data) => {
    return request.put(`/trading-review/${id}`, data)
  },

  // 删除复盘
  deleteReview: (id) => {
    return request.delete(`/trading-review/${id}`)
  },

  // 获取复盘统计
  getReviewStats: () => {
    return request.get('/trading-review/stats')
  }
}

export default tradingReviewApi
