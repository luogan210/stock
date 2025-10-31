/**
 * 通用 API 工具类
 * 只提供 HTTP 请求工具方法，不包含业务逻辑
 *
 * 职责：
 * 1. 提供统一的 HTTP 请求方法
 * 2. 统一处理 API 基础路径
 * 3. 不涉及具体业务逻辑
 */

import http from './http'

// 基础 API 配置
const API_BASE = '/api'

/**
 * 通用请求方法
 * 封装了常用的 HTTP 请求方法，统一处理 API 基础路径
 * 所有业务逻辑的接口都应该通过这个方法来发送请求
 */
export const request = {
  // 允许传入 axios config，便于使用 showBizError / silent 等开关
  get: (url, params = {}, config = {}) => {
    const merged = { ...(config || {}) };
    merged.params = { ...(config?.params || {}), ...(params || {}) };
    return http.get(`${API_BASE}${url}`, merged);
  },
  post: (url, data = {}, config = {}) => http.post(`${API_BASE}${url}`, data, config),
  put: (url, data = {}, config = {}) => http.put(`${API_BASE}${url}`, data, config),
  delete: (url, data = {}, config = {}) => {
    const merged = { ...(config || {}), data };
    return http.delete(`${API_BASE}${url}`, merged);
  },
  patch: (url, data = {}, config = {}) => http.patch(`${API_BASE}${url}`, data, config)
}

export default request
