/**
 * 统一 API 服务入口
 * 集中管理所有业务 API，避免循环依赖
 */

// 导入各个页面的 API 服务
import { tradingPlanApi } from '@/views/trading-plan/api'
import { tradingLogApi } from '@/views/trading-log/api'
import { tradingReviewApi } from '@/views/trading-review/api'
import { homeApi } from '@/views/home/api'

// 导入模拟 API 服务
import { mockTradingPlanApi } from '@/views/trading-plan/api/mock'
import { mockTradingLogApi } from '@/views/trading-log/api/mock'
import { mockTradingReviewApi } from '@/views/trading-review/api/mock'
import { mockHomeApi } from '@/views/home/api/mock'

// 根据环境选择使用真实 API 还是模拟 API
const isDevelopment = process.env.NODE_ENV === 'development'

// 统一导出所有 API 服务
export const api = {
  // 真实 API 服务
  tradingPlan: tradingPlanApi,
  tradingLog: tradingLogApi,
  tradingReview: tradingReviewApi,
  home: homeApi,
  
  // 模拟 API 服务
  mock: {
    tradingPlan: mockTradingPlanApi,
    tradingLog: mockTradingLogApi,
    tradingReview: mockTradingReviewApi,
    home: mockHomeApi
  }
}

// 根据环境自动选择 API
// 注意：为对接后端，交易计划、交易日志、交易复盘在开发环境也强制使用真实接口
export const currentApi = isDevelopment ? {
  tradingPlan: tradingPlanApi,
  tradingLog: tradingLogApi,
  tradingReview: tradingReviewApi,
  home: mockHomeApi
} : {
  tradingPlan: tradingPlanApi,
  tradingLog: tradingLogApi,
  tradingReview: tradingReviewApi,
  home: homeApi
}

// 兼容性导出（保持向后兼容）
export { tradingPlanApi, tradingLogApi, tradingReviewApi, homeApi }
export { mockTradingPlanApi, mockTradingLogApi, mockTradingReviewApi, mockHomeApi }

export default api
