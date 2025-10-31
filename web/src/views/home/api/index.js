// 首页相关 API 服务
import { request } from '@/services/api.js'

// 获取首页统计数据
export const getHomeStats = () => request.get('/home/stats')

// 获取当日交易次数
export const getTodayTradingCount = () => request.get('/home/today-trading')

// 获取当周交易次数
export const getWeekTradingCount = () => request.get('/home/week-trading')

// 获取当月交易次数
export const getMonthTradingCount = () => request.get('/home/month-trading')

// 获取快速操作数据
export const getQuickActions = () => request.get('/home/quick-actions')

// 获取最近交易
export const getRecentTrades = (limit = 5) => request.get('/home/recent-trades', { limit })

// 获取市场概览
export const getMarketOverview = () => request.get('/home/market-overview')

// 默认导出
export default {
  getHomeStats,
  getTodayTradingCount,
  getWeekTradingCount,
  getMonthTradingCount,
  getQuickActions,
  getRecentTrades,
  getMarketOverview
}
