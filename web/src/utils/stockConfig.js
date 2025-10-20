/**
 * 股票数据配置 - 轻量级前端维护
 * 统一管理所有股票的基础信息
 */

// 股票市场
export const STOCK_MARKETS = {
  A_SHARE: 'a_share',              // A股
  HK_STOCK: 'hk_stock',            // 港股
  US_STOCK: 'us_stock',            // 美股
  OTHER: 'other'                   // 其他市场
}

// 股票分类
export const STOCK_CATEGORIES = {
  MAIN_BOARD: 'main_board',        // 主板
  SME_BOARD: 'sme_board',          // 中小板
  GEM_BOARD: 'gem_board',          // 创业板
  STAR_BOARD: 'star_board',        // 科创板
  NEW_THIRD_BOARD: 'new_third_board', // 新三板
  HK_MAIN: 'hk_main',              // 港股主板
  HK_GEM: 'hk_gem',                // 港股创业板
  US_NYSE: 'us_nyse',              // 纽交所
  US_NASDAQ: 'us_nasdaq',          // 纳斯达克
  US_AMEX: 'us_amex'               // 美交所
}

// 货币类型
export const CURRENCY_TYPES = {
  CNY: 'cny',                      // 人民币
  HKD: 'hkd',                      // 港币
  USD: 'usd',                      // 美元
  EUR: 'eur',                      // 欧元
  JPY: 'jpy'                       // 日元
}

// 行业分类
export const INDUSTRY_CATEGORIES = {
  TECHNOLOGY: 'technology',         // 科技
  FINANCE: 'finance',               // 金融
  HEALTHCARE: 'healthcare',         // 医疗
  CONSUMER: 'consumer',             // 消费
  INDUSTRY: 'industry',             // 工业
  ENERGY: 'energy',                 // 能源
  MATERIALS: 'materials',          // 材料
  UTILITIES: 'utilities',          // 公用事业
  REAL_ESTATE: 'real_estate',      // 房地产
  COMMUNICATION: 'communication'   // 通信
}

// 股票数据配置
export const stockConfig = [


  
  {
    id: '300750',
    code: '300750',
    name: '宁德时代',
    category: STOCK_CATEGORIES.GEM_BOARD,
    industry: INDUSTRY_CATEGORIES.TECHNOLOGY,
    market: '深圳',
    listingDate: '2018-06-11',
    marketCap: 1200000000000,
    pe: 45.6,
    pb: 8.5,
    dividend: 0.01,
    volatility: 0.45,
    liquidity: 'high',
    riskLevel: 'high',
    description: '全球领先的动力电池制造商',
    tags: ['新能源', '电池', '科技'],
    enabled: true
  },
  
  {
    id: '688981',
    code: '688981',
    name: '中芯国际',
    category: STOCK_CATEGORIES.STAR_BOARD,
    industry: INDUSTRY_CATEGORIES.TECHNOLOGY,
    market: '上海',
    listingDate: '2020-07-16',
    marketCap: 320000000000,
    pe: 28.5,
    pb: 2.8,
    dividend: 0.005,
    volatility: 0.50,
    liquidity: 'medium',
    riskLevel: 'high',
    description: '中国领先的半导体制造企业',
    tags: ['半导体', '芯片', '科技'],
    enabled: true
  },
  
  {
    id: '600036',
    code: '600036',
    name: '招商银行',
    category: STOCK_CATEGORIES.MAIN_BOARD,
    industry: INDUSTRY_CATEGORIES.FINANCE,
    market: '上海',
    listingDate: '2002-04-09',
    marketCap: 380000000000,
    pe: 5.8,
    pb: 0.9,
    dividend: 0.12,
    volatility: 0.22,
    liquidity: 'high',
    riskLevel: 'low',
    description: '中国领先的股份制商业银行',
    tags: ['银行', '金融', '零售'],
    enabled: true
  },
  
  

  
  // 港股示例
  {
    id: '00700',
    code: '00700',
    name: '腾讯控股',
    category: STOCK_CATEGORIES.HK_MAIN,
    industry: INDUSTRY_CATEGORIES.TECHNOLOGY,
    market: '香港',
    marketType: STOCK_MARKETS.HK_STOCK,
    currency: CURRENCY_TYPES.HKD,
    listingDate: '2004-06-16',
    marketCap: 3500000000000,  // 3.5万亿港币
    pe: 18.5,
    pb: 2.8,
    dividend: 0.008,
    volatility: 0.40,
    liquidity: 'high',
    riskLevel: 'medium',
    description: '中国领先的互联网科技公司，社交和游戏业务龙头',
    tags: ['互联网', '科技', '社交', '游戏'],
    enabled: true
  },
  
  {
    id: '00941',
    code: '00941',
    name: '中国移动',
    category: STOCK_CATEGORIES.HK_MAIN,
    industry: INDUSTRY_CATEGORIES.COMMUNICATION,
    market: '香港',
    marketType: STOCK_MARKETS.HK_STOCK,
    currency: CURRENCY_TYPES.HKD,
    listingDate: '1997-10-23',
    marketCap: 1200000000000,  // 1.2万亿港币
    pe: 8.2,
    pb: 0.9,
    dividend: 0.065,
    volatility: 0.20,
    liquidity: 'high',
    riskLevel: 'low',
    description: '中国最大的移动通信运营商',
    tags: ['通信', '运营商', '国企'],
    enabled: true
  },
  
  // 美股示例
  {
    id: 'AAPL',
    code: 'AAPL',
    name: '苹果公司',
    category: STOCK_CATEGORIES.US_NASDAQ,
    industry: INDUSTRY_CATEGORIES.TECHNOLOGY,
    market: '纳斯达克',
    marketType: STOCK_MARKETS.US_STOCK,
    currency: CURRENCY_TYPES.USD,
    listingDate: '1980-12-12',
    marketCap: 2800000000000,  // 2.8万亿美元
    pe: 28.5,
    pb: 5.2,
    dividend: 0.004,
    volatility: 0.25,
    liquidity: 'high',
    riskLevel: 'medium',
    description: '全球领先的科技公司，iPhone和Mac制造商',
    tags: ['科技', '消费电子', '品牌'],
    enabled: true
  },
  
  {
    id: 'TSLA',
    code: 'TSLA',
    name: '特斯拉',
    category: STOCK_CATEGORIES.US_NASDAQ,
    industry: INDUSTRY_CATEGORIES.TECHNOLOGY,
    market: '纳斯达克',
    marketType: STOCK_MARKETS.US_STOCK,
    currency: CURRENCY_TYPES.USD,
    listingDate: '2010-06-29',
    marketCap: 800000000000,   // 8000亿美元
    pe: 65.8,
    pb: 12.5,
    dividend: 0,
    volatility: 0.60,
    liquidity: 'high',
    riskLevel: 'high',
    description: '全球领先的电动汽车和清洁能源公司',
    tags: ['电动车', '新能源', '科技'],
    enabled: true
  },
  
  {
    id: 'MSFT',
    code: 'MSFT',
    name: '微软',
    category: STOCK_CATEGORIES.US_NASDAQ,
    industry: INDUSTRY_CATEGORIES.TECHNOLOGY,
    market: '纳斯达克',
    marketType: STOCK_MARKETS.US_STOCK,
    currency: CURRENCY_TYPES.USD,
    listingDate: '1986-03-13',
    marketCap: 2500000000000,  // 2.5万亿美元
    pe: 32.5,
    pb: 8.8,
    dividend: 0.008,
    volatility: 0.22,
    liquidity: 'high',
    riskLevel: 'medium',
    description: '全球领先的软件和云计算公司',
    tags: ['软件', '云计算', '科技'],
    enabled: true
  }
]

// 获取所有启用的股票
export const getEnabledStocks = () => {
  return stockConfig.filter(stock => stock.enabled)
}

// 根据分类获取股票
export const getStocksByCategory = (category) => {
  return stockConfig.filter(stock => 
    stock.category === category && stock.enabled
  )
}

// 根据行业获取股票
export const getStocksByIndustry = (industry) => {
  return stockConfig.filter(stock => 
    stock.industry === industry && stock.enabled
  )
}

// 根据风险等级获取股票
export const getStocksByRiskLevel = (riskLevel) => {
  return stockConfig.filter(stock => 
    stock.riskLevel === riskLevel && stock.enabled
  )
}

// 根据市场类型获取股票
export const getStocksByMarket = (marketType) => {
  return stockConfig.filter(stock => 
    stock.marketType === marketType && stock.enabled
  )
}

// 根据货币类型获取股票
export const getStocksByCurrency = (currency) => {
  return stockConfig.filter(stock => 
    stock.currency === currency && stock.enabled
  )
}

// 获取股票详情
export const getStockById = (id) => {
  return stockConfig.find(stock => stock.id === id)
}

// 根据股票代码获取股票
export const getStockByCode = (code) => {
  return stockConfig.find(stock => stock.code === code)
}

// 搜索股票
export const searchStocks = (keyword) => {
  const lowerKeyword = keyword.toLowerCase()
  return stockConfig.filter(stock => 
    stock.enabled && (
      stock.code.includes(keyword) ||
      stock.name.includes(keyword) ||
      stock.tags.some(tag => tag.includes(keyword))
    )
  )
}

// 获取股票名称
export const getStockName = (code) => {
  const stock = getStockByCode(code)
  return stock ? stock.name : '未知股票'
}

// 获取股票信息
export const getStockInfo = (code) => {
  const stock = getStockByCode(code)
  if (!stock) return null
  
  return {
    code: stock.code,
    name: stock.name,
    category: stock.category,
    industry: stock.industry,
    market: stock.market,
    marketCap: stock.marketCap,
    pe: stock.pe,
    pb: stock.pb,
    dividend: stock.dividend,
    volatility: stock.volatility,
    liquidity: stock.liquidity,
    riskLevel: stock.riskLevel,
    description: stock.description,
    tags: stock.tags
  }
}

// 更新股票配置
export const updateStockConfig = (id, updates) => {
  const index = stockConfig.findIndex(stock => stock.id === id)
  if (index !== -1) {
    stockConfig[index] = { ...stockConfig[index], ...updates }
    return true
  }
  return false
}

// 添加新股票
export const addStock = (stock) => {
  const newStock = {
    id: stock.id || `stock_${Date.now()}`,
    enabled: true,
    ...stock
  }
  stockConfig.push(newStock)
  return newStock
}

// 删除股票（软删除）
export const deleteStock = (id) => {
  return updateStockConfig(id, { enabled: false })
}

// 获取市场文本
export const getMarketText = (marketType) => {
  const texts = {
    [STOCK_MARKETS.A_SHARE]: 'A股',
    [STOCK_MARKETS.HK_STOCK]: '港股',
    [STOCK_MARKETS.US_STOCK]: '美股',
    [STOCK_MARKETS.OTHER]: '其他'
  }
  return texts[marketType] || '未知'
}

// 获取货币文本
export const getCurrencyText = (currency) => {
  const texts = {
    [CURRENCY_TYPES.CNY]: '人民币',
    [CURRENCY_TYPES.HKD]: '港币',
    [CURRENCY_TYPES.USD]: '美元',
    [CURRENCY_TYPES.EUR]: '欧元',
    [CURRENCY_TYPES.JPY]: '日元'
  }
  return texts[currency] || '未知'
}

// 获取分类文本
export const getCategoryText = (category) => {
  const texts = {
    [STOCK_CATEGORIES.MAIN_BOARD]: '主板',
    [STOCK_CATEGORIES.SME_BOARD]: '中小板',
    [STOCK_CATEGORIES.GEM_BOARD]: '创业板',
    [STOCK_CATEGORIES.STAR_BOARD]: '科创板',
    [STOCK_CATEGORIES.NEW_THIRD_BOARD]: '新三板',
    [STOCK_CATEGORIES.HK_MAIN]: '港股主板',
    [STOCK_CATEGORIES.HK_GEM]: '港股创业板',
    [STOCK_CATEGORIES.US_NYSE]: '纽交所',
    [STOCK_CATEGORIES.US_NASDAQ]: '纳斯达克',
    [STOCK_CATEGORIES.US_AMEX]: '美交所'
  }
  return texts[category] || '未知'
}

// 获取行业文本
export const getIndustryText = (industry) => {
  const texts = {
    [INDUSTRY_CATEGORIES.TECHNOLOGY]: '科技',
    [INDUSTRY_CATEGORIES.FINANCE]: '金融',
    [INDUSTRY_CATEGORIES.HEALTHCARE]: '医疗',
    [INDUSTRY_CATEGORIES.CONSUMER]: '消费',
    [INDUSTRY_CATEGORIES.INDUSTRY]: '工业',
    [INDUSTRY_CATEGORIES.ENERGY]: '能源',
    [INDUSTRY_CATEGORIES.MATERIALS]: '材料',
    [INDUSTRY_CATEGORIES.UTILITIES]: '公用事业',
    [INDUSTRY_CATEGORIES.REAL_ESTATE]: '房地产',
    [INDUSTRY_CATEGORIES.COMMUNICATION]: '通信'
  }
  return texts[industry] || '未知'
}

// 获取风险等级文本
export const getRiskLevelText = (riskLevel) => {
  const texts = {
    low: '低风险',
    medium: '中风险',
    high: '高风险'
  }
  return texts[riskLevel] || '未知'
}

// 获取流动性文本
export const getLiquidityText = (liquidity) => {
  const texts = {
    high: '高流动性',
    medium: '中流动性',
    low: '低流动性'
  }
  return texts[liquidity] || '未知'
}

// 格式化市值
export const formatMarketCap = (marketCap) => {
  if (marketCap >= 1000000000000) {
    return `${(marketCap / 1000000000000).toFixed(1)}万亿`
  } else if (marketCap >= 100000000) {
    return `${(marketCap / 100000000).toFixed(1)}亿`
  } else if (marketCap >= 10000) {
    return `${(marketCap / 10000).toFixed(1)}万`
  }
  return marketCap.toString()
}

// 格式化数字
export const formatNumber = (num, decimals = 2) => {
  if (num >= 100000000) {
    return `${(num / 100000000).toFixed(decimals)}亿`
  } else if (num >= 10000) {
    return `${(num / 10000).toFixed(decimals)}万`
  }
  return num.toFixed(decimals)
}
