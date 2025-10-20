// 首页相关模拟 API 服务
import { mockTradingLogApi } from '@/views/trading-log/api/mock'

// 模拟延迟
const delay = (ms = 300) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟 API 响应
const createMockResponse = (data, message = '操作成功') => ({
  code: 0,
  message,
  data,
  timestamp: Date.now()
})

// 首页相关模拟 API
export const mockHomeApi = {
  getHomeStats: async () => {
    await delay()
    const stats = {
      stockCount: 8, // 模拟股票数量
      watchCount: 5, // 模拟关注股票数量
      planCount: 2, // 模拟交易计划数量
      logCount: 2, // 模拟交易日志数量
      totalProfit: 0, // 模拟总盈亏
      todayTradingCount: 2,
      weekTradingCount: 8,
      monthTradingCount: 25
    }
    return createMockResponse(stats)
  },
  
  getTodayTradingCount: async () => {
    await delay()
    return createMockResponse({ count: 2 })
  },
  
  getWeekTradingCount: async () => {
    await delay()
    return createMockResponse({ count: 8 })
  },
  
  getMonthTradingCount: async () => {
    await delay()
    return createMockResponse({ count: 25 })
  },
  
  getQuickActions: async () => {
    await delay()
    const actions = [
      { id: 1, name: '创建交易计划', icon: 'add-circle', path: '/trading-plan/create' },
      { id: 2, name: '记录交易日志', icon: 'edit', path: '/trading-log/create' },
      { id: 3, name: '查看股票分析', icon: 'chart', path: '/stock-analysis' },
      { id: 4, name: '系统设置', icon: 'setting', path: '/settings' }
    ]
    return createMockResponse(actions)
  },
  
  getRecentTrades: async (limit = 5) => {
    await delay()
    const response = await mockTradingLogApi.getLogs({ pageSize: limit })
    if (response.code === 0) {
      return createMockResponse(response.data.list)
    }
    return createMockResponse([])
  },
  
  getMarketOverview: async () => {
    await delay()
    const overview = {
      marketStatus: 'open',
      totalStocks: 5000,
      risingStocks: 2800,
      fallingStocks: 2000,
      unchangedStocks: 200,
      marketIndex: {
        shanghai: { name: '上证指数', value: 3200.50, change: 15.20, changePercent: 0.48 },
        shenzhen: { name: '深证成指', value: 12000.30, change: -25.80, changePercent: -0.21 },
        chinext: { name: '创业板指', value: 2500.80, change: 8.50, changePercent: 0.34 }
      }
    }
    return createMockResponse(overview)
  }
}

export default mockHomeApi