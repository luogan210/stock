/**
 * 交易复盘 API 服务
 * 真实环境下的 API 接口
 */

import { request } from '@/services/apiService'

export const tradingReviewApi = {
  // 获取复盘列表
  getReviews: (params = {}) => {
    return request.get('/reviews/getList', params)
  },

  // 获取复盘详情
  getReview: (id) => {
    return request.get(`/reviews/getDetail/${id}`)
  },

  // 创建复盘
  createReview: (data) => {
    return request.post('/reviews/create', data)
  },

  // 更新复盘
  updateReview: (id, data) => {
    return request.put(`/reviews/update/${id}`, data)
  },

  // 删除复盘
  deleteReview: (id) => {
    return request.delete(`/reviews/delete/${id}`)
  }
}

export default tradingReviewApi
