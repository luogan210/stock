/**
 * 交易复盘 API 模拟服务
 * 开发环境下的模拟数据
 */

// 模拟复盘数据 - 支持日/周/月复盘
const mockReviews = [
  {
    id: '1',
    title: '2024年1月8日交易复盘',
    period: 'daily',
    tradingDate: '2024-01-08',
    dateRange: '2024-01-08',
    totalTrades: 5,
    winTrades: 3,
    lossTrades: 2,
    buyCount: 3,
    sellCount: 2,
    winRate: 60,
    totalProfit: 1200,
    maxLoss: -300,
    avgProfit: 400,
    avgLoss: -150,
    content: '今日主要操作了科技股，整体表现不错。早盘买入的AI概念股表现良好，但下午的半导体股操作失误，需要改进。',
    lessons: [
      'AI概念股在早盘表现活跃，适合短线操作',
      '半导体股在下午容易回调，需要谨慎',
      '单笔亏损控制在总资金的2%以内'
    ],
    improvements: [
      '加强对半导体板块的研究',
      '优化买入时机选择',
      '严格执行止损策略'
    ],
    nextPlan: '重点关注AI板块，控制仓位，严格执行止损',
    createTime: '2024-01-08 20:30:00',
    updateTime: '2024-01-08 20:30:00',
    status: 'completed'
  },
  {
    id: '2',
    title: '2024年1月第一周交易复盘',
    period: 'weekly',
    tradingDate: '2024-01-05',
    dateRange: '2024-01-01~2024-01-05',
    totalTrades: 15,
    winTrades: 9,
    lossTrades: 6,
    buyCount: 8,
    sellCount: 7,
    winRate: 60,
    totalProfit: 2500,
    maxLoss: -800,
    avgProfit: 277.78,
    avgLoss: -133.33,
    content: '本周交易表现良好，主要盈利来源于科技股的操作。需要改进的是止损策略，避免单笔亏损过大。',
    lessons: [
      '科技股在财报季表现活跃，适合短线操作',
      '止损设置需要更加严格，避免情绪化交易',
      '仓位管理需要优化，单笔交易不应超过总资金的5%'
    ],
    improvements: [
      '建立更严格的止损规则',
      '增加技术分析的学习',
      '优化仓位分配策略'
    ],
    nextPlan: '重点关注科技股财报，控制仓位，严格执行止损',
    createTime: '2024-01-05 20:30:00',
    updateTime: '2024-01-05 20:30:00',
    status: 'completed'
  },
  {
    id: '3',
    title: '2024年1月月度交易复盘',
    period: 'monthly',
    tradingDate: '2024-01-31',
    dateRange: '2024年1月',
    totalTrades: 45,
    winTrades: 28,
    lossTrades: 17,
    buyCount: 22,
    sellCount: 23,
    winRate: 62.22,
    totalProfit: 8500,
    maxLoss: -1200,
    avgProfit: 303.57,
    avgLoss: -70.59,
    content: '本月整体表现不错，实现了稳定盈利。主要成功因素是对科技股的准确把握和严格的止损纪律。',
    lessons: [
      '科技股是本月的主要盈利来源',
      '严格的止损纪律避免了大幅亏损',
      '仓位管理策略效果良好'
    ],
    improvements: [
      '继续优化科技股的操作策略',
      '加强其他板块的研究',
      '完善风险控制体系'
    ],
    nextPlan: '继续关注科技股，同时研究其他板块机会，完善交易系统',
    createTime: '2024-01-31 22:00:00',
    updateTime: '2024-01-31 22:00:00',
    status: 'completed'
  }
]

export const mockTradingReviewApi = {
  // 获取复盘列表
  getReviews: (params = {}) => {
    const { page = 1, pageSize = 10, period, status, keyword, tradingDate, dateRange } = params
    
    let filteredReviews = [...mockReviews]
    
    // 按周期筛选
    if (period) {
      filteredReviews = filteredReviews.filter(review => review.period === period)
    }
    
    // 按状态筛选
    if (status) {
      filteredReviews = filteredReviews.filter(review => review.status === status)
    }
    
    // 按关键词筛选
    if (keyword) {
      filteredReviews = filteredReviews.filter(review => 
        review.title.toLowerCase().includes(keyword.toLowerCase())
      )
    }
    
    // 按交易日期筛选
    if (tradingDate) {
      filteredReviews = filteredReviews.filter(review => review.tradingDate === tradingDate)
    }
    
    // 分页
    const start = (page - 1) * pageSize
    const end = start + pageSize
    const list = filteredReviews.slice(start, end)
    
    return Promise.resolve({
      code: 0,
      message: 'success',
      data: {
        list,
        total: filteredReviews.length,
        page,
        pageSize
      }
    })
  },

  // 获取复盘详情
  getReview: (id) => {
    const review = mockReviews.find(r => r.id === id)
    if (review) {
      return Promise.resolve({
        code: 0,
        message: 'success',
        data: review
      })
    } else {
      return Promise.resolve({
        code: 404,
        message: '复盘记录不存在',
        data: null
      })
    }
  },

  // 创建复盘
  createReview: (data) => {
    const newReview = {
      ...data,
      id: Date.now().toString(),
      createTime: new Date().toLocaleString(),
      updateTime: new Date().toLocaleString(),
      status: 'completed'
    }
    
    mockReviews.unshift(newReview)
    
    return Promise.resolve({
      code: 0,
      message: 'success',
      data: newReview
    })
  },

  // 更新复盘
  updateReview: (id, data) => {
    const index = mockReviews.findIndex(r => r.id === id)
    if (index !== -1) {
      mockReviews[index] = {
        ...mockReviews[index],
        ...data,
        updateTime: new Date().toLocaleString()
      }
      
      return Promise.resolve({
        code: 0,
        message: 'success',
        data: mockReviews[index]
      })
    } else {
      return Promise.resolve({
        code: 404,
        message: '复盘记录不存在',
        data: null
      })
    }
  },

  // 删除复盘
  deleteReview: (id) => {
    const index = mockReviews.findIndex(r => r.id === id)
    if (index !== -1) {
      mockReviews.splice(index, 1)
      return Promise.resolve({
        code: 0,
        message: 'success',
        data: null
      })
    } else {
      return Promise.resolve({
        code: 404,
        message: '复盘记录不存在',
        data: null
      })
    }
  },

  // 获取复盘统计
  getReviewStats: () => {
    const totalReviews = mockReviews.length
    const dailyReviews = mockReviews.filter(r => r.period === 'daily').length
    const weeklyReviews = mockReviews.filter(r => r.period === 'weekly').length
    const monthlyReviews = mockReviews.filter(r => r.period === 'monthly').length
    const avgWinRate = mockReviews.reduce((sum, r) => sum + r.winRate, 0) / totalReviews
    const totalProfit = mockReviews.reduce((sum, r) => sum + r.totalProfit, 0)
    
    return Promise.resolve({
      code: 0,
      message: 'success',
      data: {
        totalReviews,
        dailyReviews,
        weeklyReviews,
        monthlyReviews,
        avgWinRate: Math.round(avgWinRate * 100) / 100,
        totalProfit,
        recentReviews: mockReviews.slice(0, 3)
      }
    })
  }
}

export default mockTradingReviewApi
