// 交易计划相关模拟 API 服务
import { generateId } from '@/utils/helpers'

// 模拟延迟
const delay = (ms = 300) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟数据存储
let mockTradingPlans = [
  {
    id: '1',
    name: '平安银行买多计划',
    type: 'buy',
    stockCode: '000001',
    stockName: '平安银行',
    targetPrice: 12.50,
    quantity: 1000,
    stopLoss: 11.50,
    takeProfit: 13.50,
    status: 'active',
    progress: 65,
    strategy: 'trend_following',
    riskLevel: 'medium',
    description: '基于技术分析的买多计划',
    remark: '关注银行板块整体走势',
    createTime: '2024-01-15 09:30:00',
    updateTime: '2024-01-15 14:30:00'
  },
  {
    id: '2',
    name: '贵州茅台买空计划',
    type: 'sell',
    stockCode: '600519',
    stockName: '贵州茅台',
    targetPrice: 1850.00,
    quantity: 100,
    stopLoss: 1900.00,
    takeProfit: 1800.00,
    status: 'completed',
    progress: 100,
    strategy: 'breakout',
    riskLevel: 'high',
    description: '获利了结买空计划',
    remark: '已达到预期收益目标',
    createTime: '2024-01-14 10:00:00',
    updateTime: '2024-01-15 15:00:00'
  }
]

// 模拟 API 响应
const createMockResponse = (data, message = '操作成功') => ({
  code: 0,
  message,
  data,
  timestamp: Date.now()
})

// 模拟错误响应
const createMockError = (message = '操作失败', code = 500) => ({
  code,
  message,
  data: null,
  timestamp: Date.now()
})

// 交易计划相关模拟 API
export const mockTradingPlanApi = {
  getPlans: async (params = {}) => {
    await delay()
    let plans = [...mockTradingPlans]
    
    if (params.status) {
      plans = plans.filter(plan => plan.status === params.status)
    }
    
    if (params.type) {
      plans = plans.filter(plan => plan.type === params.type)
    }
    
    if (params.keyword) {
      plans = plans.filter(plan => 
        plan.name.includes(params.keyword) ||
        plan.stockName.includes(params.keyword) ||
        plan.stockCode.includes(params.keyword)
      )
    }
    
    return createMockResponse({
      list: plans,
      total: plans.length,
      page: params.page || 1,
      pageSize: params.pageSize || 10
    })
  },
  
  getPlanDetail: async (id) => {
    await delay()
    const plan = mockTradingPlans.find(p => p.id === id)
    if (!plan) {
      return createMockError('交易计划不存在', 404)
    }
    return createMockResponse(plan)
  },
  
  createPlan: async (data) => {
    await delay()
    const newPlan = {
      ...data,
      id: generateId(),
      createTime: new Date().toLocaleString(),
      updateTime: new Date().toLocaleString()
    }
    mockTradingPlans.unshift(newPlan)
    return createMockResponse(newPlan, '交易计划创建成功')
  },
  
  updatePlan: async (id, data) => {
    await delay()
    const index = mockTradingPlans.findIndex(p => p.id === id)
    if (index === -1) {
      return createMockError('交易计划不存在', 404)
    }
    mockTradingPlans[index] = {
      ...mockTradingPlans[index],
      ...data,
      updateTime: new Date().toLocaleString()
    }
    return createMockResponse(mockTradingPlans[index], '交易计划更新成功')
  },
  
  deletePlan: async (id) => {
    await delay()
    const index = mockTradingPlans.findIndex(p => p.id === id)
    if (index === -1) {
      return createMockError('交易计划不存在', 404)
    }
    mockTradingPlans.splice(index, 1)
    return createMockResponse(null, '交易计划删除成功')
  },
  
  updatePlanStatus: async (id, status) => {
    await delay()
    const plan = mockTradingPlans.find(p => p.id === id)
    if (!plan) {
      return createMockError('交易计划不存在', 404)
    }
    plan.status = status
    plan.updateTime = new Date().toLocaleString()
    if (status === 'completed') {
      plan.progress = 100
    }
    return createMockResponse(plan, '计划状态更新成功')
  },
  
  getActivePlans: async () => {
    await delay()
    const activePlans = mockTradingPlans.filter(plan => plan.status === 'active')
    return createMockResponse(activePlans)
  },
  
  getCompletedPlans: async () => {
    await delay()
    const completedPlans = mockTradingPlans.filter(plan => plan.status === 'completed')
    return createMockResponse(completedPlans)
  },
  
  batchUpdatePlanStatus: async (ids, status) => {
    await delay()
    const updatedPlans = []
    ids.forEach(id => {
      const plan = mockTradingPlans.find(p => p.id === id)
      if (plan) {
        plan.status = status
        plan.updateTime = new Date().toLocaleString()
        if (status === 'completed') {
          plan.progress = 100
        }
        updatedPlans.push(plan)
      }
    })
    return createMockResponse(updatedPlans, '批量更新成功')
  },
  
  copyPlan: async (id) => {
    await delay()
    const originalPlan = mockTradingPlans.find(p => p.id === id)
    if (!originalPlan) {
      return createMockError('交易计划不存在', 404)
    }
    
    const copiedPlan = {
      ...originalPlan,
      id: generateId(),
      name: `${originalPlan.name} - 副本`,
      status: 'active',
      progress: 0,
      createTime: new Date().toLocaleString(),
      updateTime: new Date().toLocaleString()
    }
    mockTradingPlans.unshift(copiedPlan)
    return createMockResponse(copiedPlan, '交易计划复制成功')
  },
  
  exportPlans: async (params = {}) => {
    await delay()
    let plans = [...mockTradingPlans]
    
    if (params.status) {
      plans = plans.filter(plan => plan.status === params.status)
    }
    
    if (params.type) {
      plans = plans.filter(plan => plan.type === params.type)
    }
    
    // 模拟导出数据
    const exportData = {
      plans,
      exportTime: new Date().toISOString(),
      totalCount: plans.length
    }
    
    return createMockResponse(exportData, '导出成功')
  }
}

export default mockTradingPlanApi
