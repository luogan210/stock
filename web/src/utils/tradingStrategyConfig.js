/**
 * 交易策略配置 - 轻量级前端维护
 * 统一管理所有交易策略的配置信息
 */

// 交易策略分类
export const TRADING_STRATEGY_CATEGORIES = {
  SCALPING: 'scalping',      // 超短线
  DAY_TRADING: 'day_trading', // 日内交易
  SWING: 'swing',            // 波段交易
  POSITION: 'position',      // 持仓交易
  ARBITRAGE: 'arbitrage'     // 套利交易
}

// 风险等级
export const RISK_LEVELS = {
  LOW: 'low',
  MEDIUM: 'medium',
  HIGH: 'high'
}

// 交易策略配置数据
export const tradingStrategyConfig = [
  {
    id: 'scalping',
    name: '剥头皮',
    category: TRADING_STRATEGY_CATEGORIES.SCALPING,
    description: '超短线交易，快速进出',
    detail: '利用极短时间内的价格波动进行交易，持仓时间通常为几分钟到几小时',
    parameters: {
      holdingTime: '几分钟到几小时',
      frequency: '高频',
      riskLevel: RISK_LEVELS.HIGH,
      capitalRequirement: '高资金要求',
      skillLevel: '需要专业交易技能'
    },
    winRate: '50-60%',
    suitableFor: '专业交易员，有充足时间盯盘',
    pros: ['快速获利', '资金周转快', '风险相对可控'],
    cons: ['需要专业技能', '手续费成本高', '心理压力大'],
    enabled: true
  },

  {
    id: 'daytrading',
    name: '日内交易',
    category: TRADING_STRATEGY_CATEGORIES.DAY_TRADING,
    description: '当日开平仓，不过夜',
    detail: '在同一个交易日内完成开仓和平仓，不持有过夜头寸',
    parameters: {
      holdingTime: '当日完成',
      frequency: '中频',
      riskLevel: RISK_LEVELS.MEDIUM,
      capitalRequirement: '中等资金要求',
      skillLevel: '需要一定交易经验'
    },
    winRate: '55-65%',
    suitableFor: '有时间的个人投资者，日内波动较大的市场',
    pros: ['避免隔夜风险', '资金利用率高', '适合上班族'],
    cons: ['需要盯盘时间', '可能错过趋势机会', '手续费成本'],
    enabled: true
  },

  {
    id: 'swing',
    name: '波段交易',
    category: TRADING_STRATEGY_CATEGORIES.SWING,
    description: '短期趋势，持仓数日',
    detail: '捕捉短期趋势，持仓时间通常为几天到几周',
    parameters: {
      holdingTime: '几天到几周',
      frequency: '低频',
      riskLevel: RISK_LEVELS.MEDIUM,
      capitalRequirement: '中等资金要求',
      skillLevel: '需要技术分析能力'
    },
    winRate: '60-70%',
    suitableFor: '有一定技术分析能力的投资者',
    pros: ['风险相对可控', '适合上班族', '技术分析友好'],
    cons: ['需要技术分析能力', '可能错过短期机会', '需要耐心'],
    enabled: true
  },

  {
    id: 'position',
    name: '持仓交易',
    category: TRADING_STRATEGY_CATEGORIES.POSITION,
    description: '中长期持仓，价值投资',
    detail: '基于基本面分析，长期持有优质股票',
    parameters: {
      holdingTime: '几个月到几年',
      frequency: '极低频',
      riskLevel: RISK_LEVELS.LOW,
      capitalRequirement: '资金要求灵活',
      skillLevel: '需要基本面分析能力'
    },
    winRate: '70-80%',
    suitableFor: '价值投资者，长期投资者',
    pros: ['风险较低', '适合长期投资', '手续费成本低'],
    cons: ['资金占用时间长', '需要基本面分析', '可能错过短期机会'],
    enabled: true
  },

  {
    id: 'arbitrage',
    name: '套利交易',
    category: TRADING_STRATEGY_CATEGORIES.ARBITRAGE,
    description: '利用价差获利',
    detail: '利用不同市场、不同时间或不同品种之间的价格差异进行无风险套利',
    parameters: {
      holdingTime: '几分钟到几天',
      frequency: '中高频',
      riskLevel: RISK_LEVELS.LOW,
      capitalRequirement: '高资金要求',
      skillLevel: '需要专业套利技能'
    },
    winRate: '80-90%',
    suitableFor: '专业投资者，机构投资者',
    pros: ['风险极低', '收益稳定', '专业性强'],
    cons: ['需要专业设备', '资金要求高', '机会有限'],
    enabled: true
  }
]

// 获取所有启用的交易策略
export const getEnabledTradingStrategies = () => {
  return tradingStrategyConfig.filter(strategy => strategy.enabled)
}

// 根据分类获取交易策略
export const getTradingStrategiesByCategory = (category) => {
  return tradingStrategyConfig.filter(strategy =>
    strategy.category === category && strategy.enabled
  )
}

// 获取交易策略详情
export const getTradingStrategyById = (id) => {
  return tradingStrategyConfig.find(strategy => strategy.id === id)
}

// 获取策略分类文本
export const getTradingStrategyCategoryText = (category) => {
  return tradingStrategyConfig.find(strategy => strategy.category === category)?.name || '未知' 
}
