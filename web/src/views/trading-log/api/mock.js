// 交易日志相关模拟 API 服务
import { generateId } from '@/utils/helpers'

// 模拟延迟
const delay = (ms = 300) => new Promise(resolve => setTimeout(resolve, ms))

// 模拟数据存储
let mockTradingLogs = [
  {
    id: '1',
    title: '平安银行买多交易',
    type: 'buy',
    stockCode: '000001',
    stockName: '平安银行',
    price: 12.50,
    quantity: 1000,
    profit: 150.00,
    status: 'completed',
    tradingTime: '2024-01-15 14:30:00',
    fee: 12.50,
    planName: '平安银行买多计划',
    strategy: 'daytrading',
    riskLevel: 'medium',
    content: '基于技术分析，在12.50价位买入1000股平安银行，预期短期上涨。',
    remark: '关注银行板块整体走势',
    createTime: '2024-01-15 14:35:00',
    updateTime: '2024-01-15 14:35:00'
  },
  {
    id: '2',
    title: '贵州茅台买空交易',
    type: 'sell',
    stockCode: '600519',
    stockName: '贵州茅台',
    price: 1850.00,
    quantity: 100,
    profit: -500.00,
    status: 'completed',
    tradingTime: '2024-01-14 10:00:00',
    fee: 185.00,
    planName: '贵州茅台买空计划',
    strategy: 'swing',
    riskLevel: 'high',
    content: '判断茅台价格过高，进行买空操作，但市场继续上涨导致亏损。',
    remark: '需要重新评估市场趋势',
    createTime: '2024-01-14 10:05:00',
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

// 交易日志相关模拟 API
export const mockTradingLogApi = {
  getLogs: async (params = {}) => {
    await delay()
    let logs = [...mockTradingLogs]
    
    if (params.status) {
      logs = logs.filter(log => log.status === params.status)
    }
    
    if (params.type) {
      logs = logs.filter(log => log.type === params.type)
    }
    
    if (params.keyword) {
      logs = logs.filter(log => 
        log.title.includes(params.keyword) ||
        log.stockName.includes(params.keyword) ||
        log.stockCode.includes(params.keyword) ||
        log.content.includes(params.keyword)
      )
    }
    
    return createMockResponse({
      list: logs,
      total: logs.length,
      page: params.page || 1,
      pageSize: params.pageSize || 10
    })
  },
  
  getLogDetail: async (id) => {
    await delay()
    const log = mockTradingLogs.find(l => l.id === id)
    if (!log) {
      return createMockError('交易日志不存在', 404)
    }
    return createMockResponse(log)
  },
  
  createLog: async (data) => {
    await delay()
    const newLog = {
      ...data,
      id: generateId(),
      createTime: new Date().toLocaleString(),
      updateTime: new Date().toLocaleString()
    }
    mockTradingLogs.unshift(newLog)
    return createMockResponse(newLog, '交易日志创建成功')
  },
  
  updateLog: async (id, data) => {
    await delay()
    const index = mockTradingLogs.findIndex(l => l.id === id)
    if (index === -1) {
      return createMockError('交易日志不存在', 404)
    }
    mockTradingLogs[index] = {
      ...mockTradingLogs[index],
      ...data,
      updateTime: new Date().toLocaleString()
    }
    return createMockResponse(mockTradingLogs[index], '交易日志更新成功')
  },
  
  deleteLog: async (id) => {
    await delay()
    const index = mockTradingLogs.findIndex(l => l.id === id)
    if (index === -1) {
      return createMockError('交易日志不存在', 404)
    }
    mockTradingLogs.splice(index, 1)
    return createMockResponse(null, '交易日志删除成功')
  },
  
  updateLogStatus: async (id, status) => {
    await delay()
    const log = mockTradingLogs.find(l => l.id === id)
    if (!log) {
      return createMockError('交易日志不存在', 404)
    }
    log.status = status
    log.updateTime = new Date().toLocaleString()
    return createMockResponse(log, '日志状态更新成功')
  },
  
  getSuccessLogs: async () => {
    await delay()
    const successLogs = mockTradingLogs.filter(log => log.status === 'completed' && log.profit > 0)
    return createMockResponse(successLogs)
  },
  
  getFailedLogs: async () => {
    await delay()
    const failedLogs = mockTradingLogs.filter(log => log.status === 'completed' && log.profit < 0)
    return createMockResponse(failedLogs)
  },
  
  getTotalProfit: async () => {
    await delay()
    const totalProfit = mockTradingLogs.reduce((sum, log) => sum + (log.profit || 0), 0)
    return createMockResponse({ totalProfit })
  },
  
  exportLogs: async (params = {}) => {
    await delay()
    let logs = [...mockTradingLogs]
    
    if (params.status) {
      logs = logs.filter(log => log.status === params.status)
    }
    
    if (params.type) {
      logs = logs.filter(log => log.type === params.type)
    }
    
    // 模拟导出数据
    const exportData = {
      logs,
      exportTime: new Date().toISOString(),
      totalCount: logs.length,
      totalProfit: logs.reduce((sum, log) => sum + (log.profit || 0), 0)
    }
    
    return createMockResponse(exportData, '导出成功')
  },
  
  batchDeleteLogs: async (ids) => {
    await delay()
    const deletedCount = ids.filter(id => {
      const index = mockTradingLogs.findIndex(l => l.id === id)
      if (index !== -1) {
        mockTradingLogs.splice(index, 1)
        return true
      }
      return false
    }).length
    
    return createMockResponse({ deletedCount }, `成功删除 ${deletedCount} 条日志`)
  },
  
  getLogStatistics: async (params = {}) => {
    await delay()
    const stats = {
      totalLogs: mockTradingLogs.length,
      successLogs: mockTradingLogs.filter(log => log.status === 'completed' && log.profit > 0).length,
      failedLogs: mockTradingLogs.filter(log => log.status === 'completed' && log.profit < 0).length,
      totalProfit: mockTradingLogs.reduce((sum, log) => sum + (log.profit || 0), 0),
      averageProfit: mockTradingLogs.length > 0 ? 
        mockTradingLogs.reduce((sum, log) => sum + (log.profit || 0), 0) / mockTradingLogs.length : 0
    }
    return createMockResponse(stats)
  },
  
  copyLog: async (id) => {
    await delay()
    const originalLog = mockTradingLogs.find(l => l.id === id)
    if (!originalLog) {
      return createMockError('交易日志不存在', 404)
    }
    
    const copiedLog = {
      ...originalLog,
      id: generateId(),
      title: `${originalLog.title} - 副本`,
      status: 'pending',
      profit: 0,
      createTime: new Date().toLocaleString(),
      updateTime: new Date().toLocaleString()
    }
    mockTradingLogs.unshift(copiedLog)
    return createMockResponse(copiedLog, '交易日志复制成功')
  }
}

export default mockTradingLogApi
