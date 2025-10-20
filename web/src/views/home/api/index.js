// 首页相关 API 服务
import { request } from '@/services/apiService'

// 首页相关 API
export const homeApi = {
  // 获取首页统计数据
  getHomeStats: () => request.get('/home/stats'),
  
  // 获取当日交易次数
  getTodayTradingCount: () => request.get('/home/today-trading'),
  
  // 获取当周交易次数
  getWeekTradingCount: () => request.get('/home/week-trading'),
  
  // 获取当月交易次数
  getMonthTradingCount: () => request.get('/home/month-trading'),
  
  // 获取快速操作数据
  getQuickActions: () => request.get('/home/quick-actions'),
  
  // 获取最近交易
  getRecentTrades: (limit = 5) => request.get('/home/recent-trades', { limit }),
  
  // 获取市场概览
  getMarketOverview: () => request.get('/home/market-overview')
}

export default homeApi
