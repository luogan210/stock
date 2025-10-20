/**
 * 选股策略配置 - 轻量级前端维护
 * 统一管理所有选股策略的配置信息
 */

// 策略分类
export const STRATEGY_CATEGORIES = {
  TECHNICAL: 'technical',
  FUNDAMENTAL: 'fundamental'
}

// 风险等级
export const RISK_LEVELS = {
  LOW: 'low',
  MEDIUM: 'medium', 
  HIGH: 'high'
}

// 选股策略配置数据
export const strategyConfig = [
  // 技术分析策略
  {
    id: 'trend_following',
    name: '趋势跟踪',
    category: STRATEGY_CATEGORIES.TECHNICAL,
    description: '跟随市场趋势选择股票',
    detail: '通过分析价格趋势、成交量等指标，选择处于上升趋势的股票',
    parameters: {
      timeframe: '日线',
      indicators: ['MA', 'MACD', 'RSI'],
      riskLevel: RISK_LEVELS.MEDIUM
    },
    winRate: '60-70%',
    suitableFor: '趋势明显的市场环境',
    enabled: true
  },
  
  {
    id: 'breakout',
    name: '突破策略',
    category: STRATEGY_CATEGORIES.TECHNICAL,
    description: '选择突破关键阻力位的股票',
    detail: '寻找突破重要技术位、成交量放大的股票',
    parameters: {
      timeframe: '日线',
      indicators: ['支撑阻力位', '成交量'],
      riskLevel: RISK_LEVELS.HIGH
    },
    winRate: '55-65%',
    suitableFor: '震荡突破的市场环境',
    enabled: true
  },
  
  {
    id: 'reversal',
    name: '反转策略',
    category: STRATEGY_CATEGORIES.TECHNICAL,
    description: '选择超跌反弹的股票',
    detail: '寻找超卖后可能出现反弹的股票',
    parameters: {
      timeframe: '日线',
      indicators: ['RSI', 'KDJ', '布林带'],
      riskLevel: RISK_LEVELS.HIGH
    },
    winRate: '50-60%',
    suitableFor: '超跌反弹的市场环境',
    enabled: true
  },
  
  {
    id: 'bollinger_reversal',
    name: '布林带反转',
    category: STRATEGY_CATEGORIES.TECHNICAL,
    description: '利用布林带上下轨进行反向操作',
    detail: '在小时级别的布林上轨和下轨之间反向操作',
    parameters: {
      timeframe: '小时线',
      indicators: ['布林带', 'RSI'],
      riskLevel: RISK_LEVELS.MEDIUM
    },
    winRate: '60-75%',
    suitableFor: '震荡市场，高波动股票',
    enabled: true
  },
  
  {
    id: 'momentum',
    name: '动量策略',
    category: STRATEGY_CATEGORIES.TECHNICAL,
    description: '选择近期表现强势的股票',
    detail: '基于价格动量和相对强度选择股票',
    parameters: {
      timeframe: '日线',
      indicators: ['ROC', 'RSI', 'MACD'],
      riskLevel: RISK_LEVELS.MEDIUM
    },
    winRate: '55-70%',
    suitableFor: '强势上涨的市场环境',
    enabled: true
  }
]

// 获取所有启用的策略
export const getEnabledStrategies = () => {
  return strategyConfig.filter(strategy => strategy.enabled)
}

// 根据分类获取策略
export const getStrategiesByCategory = (category) => {
  return strategyConfig.filter(strategy => 
    strategy.category === category && strategy.enabled
  )
}

// 获取策略详情
export const getStrategyById = (id) => {
  return strategyConfig.find(strategy => strategy.id === id)
}

// 获取策略名称
export const getStrategyName = (id) => {
  const strategy = getStrategyById(id)
  return strategy ? strategy.name : '未知策略'
}

// 获取策略描述
export const getStrategyDescription = (id) => {
  const strategy = getStrategyById(id)
  return strategy ? strategy.description : '未知策略'
}

// 获取策略详细说明（用于弹窗显示）
export const getStrategyDetailContent = (id) => {
  const strategy = getStrategyById(id)
  if (!strategy) return null
  
  return {
    name: strategy.name,
    description: strategy.description,
    detail: strategy.detail,
    parameters: strategy.parameters,
    winRate: strategy.winRate,
    suitableFor: strategy.suitableFor
  }
}

// 获取策略选项（用于下拉选择）
export const getStrategyOptions = () => {
  return getEnabledStrategies().map(strategy => ({
    value: strategy.id,
    label: `${strategy.name} - ${strategy.description}`
  }))
}

// 更新策略配置（用于管理功能）
export const updateStrategyConfig = (id, updates) => {
  const index = strategyConfig.findIndex(strategy => strategy.id === id)
  if (index !== -1) {
    strategyConfig[index] = { ...strategyConfig[index], ...updates }
    return true
  }
  return false
}

// 添加新策略
export const addStrategy = (strategy) => {
  const newStrategy = {
    id: strategy.id || `strategy_${Date.now()}`,
    enabled: true,
    ...strategy
  }
  strategyConfig.push(newStrategy)
  return newStrategy
}

// 删除策略（软删除）
export const deleteStrategy = (id) => {
  return updateStrategyConfig(id, { enabled: false })
}
