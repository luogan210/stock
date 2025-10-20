// Store 统一入口
// 导出所有 Store

// 用户相关 Store
export { useUserStore } from './user'

// 股票相关 Store
export { useStockStore } from './stock'

// 交易相关 Store
export { useTradingPlanStore } from '@/views/trading-plan/store'
export { useTradingLogStore } from '@/views/trading-log/store'
export { useTradingReviewStore } from '@/views/trading-review/store'
export { useHomeStore } from '@/views/home/store'

// 默认导出所有 Store（保持向后兼容）
export default {
  useUserStore: () => import('./user'),
  useStockStore: () => import('./stock'),
  useTradingPlanStore: () => import('@/views/trading-plan/store'),
  useTradingLogStore: () => import('@/views/trading-log/store'),
  useTradingReviewStore: () => import('@/views/trading-review/store'),
  useHomeStore: () => import('@/views/home/store')
}