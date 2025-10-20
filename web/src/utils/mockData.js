// 模拟股票数据
export const mockStockData = [
  {
    code: '000001',
    name: '平安银行',
    price: 12.50,
    change: 0.15,
    changePercent: 1.22,
    volume: 1250000,
    marketCap: 242000000000,
    pe: 5.2,
    pb: 0.8
  },
  {
    code: '000002',
    name: '万科A',
    price: 18.30,
    change: -0.20,
    changePercent: -1.08,
    volume: 890000,
    marketCap: 203000000000,
    pe: 8.5,
    pb: 1.2
  },
  {
    code: '600036',
    name: '招商银行',
    price: 35.80,
    change: 0.45,
    changePercent: 1.27,
    volume: 2100000,
    marketCap: 950000000000,
    pe: 4.8,
    pb: 0.9
  },
  {
    code: '600519',
    name: '贵州茅台',
    price: 1850.00,
    change: 25.00,
    changePercent: 1.37,
    volume: 150000,
    marketCap: 2320000000000,
    pe: 35.2,
    pb: 8.5
  },
  {
    code: '000858',
    name: '五粮液',
    price: 168.50,
    change: 2.30,
    changePercent: 1.38,
    volume: 320000,
    marketCap: 654000000000,
    pe: 28.5,
    pb: 6.2
  }
]

// 模拟用户活动数据
export const mockUserActivities = [
  {
    id: 1,
    time: '2024-01-15 14:30',
    action: '查看了股票',
    target: '000001 平安银行',
    type: 'view'
  },
  {
    id: 2,
    time: '2024-01-15 13:45',
    action: '添加了股票到关注列表',
    target: '600036 招商银行',
    type: 'watch'
  },
  {
    id: 3,
    time: '2024-01-15 10:20',
    action: '搜索了关键词',
    target: '银行',
    type: 'search'
  },
  {
    id: 4,
    time: '2024-01-15 09:15',
    action: '查看了股票',
    target: '600519 贵州茅台',
    type: 'view'
  }
]

// 模拟系统设置
export const defaultSettings = {
  theme: 'light',
  language: 'zh-CN',
  autoRefresh: true,
  soundAlert: false,
  dataSource: 'sina',
  refreshInterval: '30',
  dataCache: true,
  historyDays: 30
}

