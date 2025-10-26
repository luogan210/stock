// 股票模块模拟 API 服务
import { mockResponse } from '@/utils/helpers'

// 模拟股票数据
const mockStocks = [
  {
    id: '1',
    code: '000001',
    name: '平安银行',
    region: 'china',
    currency: 'CNY',
    category: 'main_board',
    enabled: true,
    remark: '银行股'
  },
  {
    id: '2',
    code: '000002',
    name: '万科A',
    region: 'china',
    currency: 'CNY',
    category: 'main_board',
    enabled: true,
    remark: '地产股'
  },
  {
    id: '3',
    code: '00700',
    name: '腾讯控股',
    region: 'hongkong',
    currency: 'HKD',
    category: 'hk_main',
    enabled: true,
    remark: '科技股'
  },
  {
    id: '4',
    code: 'AAPL',
    name: '苹果公司',
    region: 'usa',
    currency: 'USD',
    category: 'us_nasdaq',
    enabled: true,
    remark: '科技股'
  }
]

// 模拟 API 服务
export const mockStockApi = {
  // 获取股票列表
  getStocks: (params = {}) => {
    return mockResponse(() => {
      let stocks = [...mockStocks]
      
      // 模拟搜索过滤
      if (params.keyword) {
        const keyword = params.keyword.toLowerCase()
        stocks = stocks.filter(stock => 
          stock.code.toLowerCase().includes(keyword) ||
          stock.name.toLowerCase().includes(keyword)
        )
      }
      
      // 模拟地区过滤
      if (params.region) {
        stocks = stocks.filter(stock => stock.region === params.region)
      }
      
      // 模拟分类过滤
      if (params.category) {
        stocks = stocks.filter(stock => stock.category === params.category)
      }
      
      return {
        items: stocks,
        total: stocks.length
      }
    })
  },

  // 获取股票详情
  getStockDetail: (id) => {
    return mockResponse(() => {
      const stock = mockStocks.find(s => s.id === id)
      if (!stock) {
        throw new Error('股票不存在')
      }
      return stock
    })
  },

  // 创建股票
  createStock: (data) => {
    return mockResponse(() => {
      const newStock = {
        id: String(mockStocks.length + 1),
        ...data,
        enabled: data.enabled !== false
      }
      mockStocks.push(newStock)
      return newStock
    })
  },

  // 更新股票
  updateStock: (id, data) => {
    return mockResponse(() => {
      const index = mockStocks.findIndex(s => s.id === id)
      if (index === -1) {
        throw new Error('股票不存在')
      }
      mockStocks[index] = { ...mockStocks[index], ...data }
      return mockStocks[index]
    })
  },

  // 删除股票
  deleteStock: (id) => {
    return mockResponse(() => {
      const index = mockStocks.findIndex(s => s.id === id)
      if (index === -1) {
        throw new Error('股票不存在')
      }
      mockStocks.splice(index, 1)
      return { id }
    })
  }
}

export default mockStockApi
