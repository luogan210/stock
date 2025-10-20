// 常量定义文件

// 交易方向
export const TRADING_TYPES = {
  BUY: 'buy',
  SELL: 'sell'
}

// 交易方向文本映射
export const TRADING_TYPE_TEXT = {
  [TRADING_TYPES.BUY]: '买多',
  [TRADING_TYPES.SELL]: '买空'
}

// 交易计划状态
export const PLAN_STATUS = {
  ACTIVE: 'active',
  COMPLETED: 'completed',
  PAUSED: 'paused',
  CANCELLED: 'cancelled'
}

// 交易计划状态文本映射
export const PLAN_STATUS_TEXT = {
  [PLAN_STATUS.ACTIVE]: '进行中',
  [PLAN_STATUS.COMPLETED]: '已完成',
  [PLAN_STATUS.PAUSED]: '已暂停',
  [PLAN_STATUS.CANCELLED]: '已取消'
}

// 交易计划状态主题映射
export const PLAN_STATUS_THEME = {
  [PLAN_STATUS.ACTIVE]: 'primary',
  [PLAN_STATUS.COMPLETED]: 'success',
  [PLAN_STATUS.PAUSED]: 'warning',
  [PLAN_STATUS.CANCELLED]: 'danger'
}

// 交易日志状态
export const LOG_STATUS = {
  PENDING: 'pending',
  COMPLETED: 'completed'
}

// 交易日志状态文本映射
export const LOG_STATUS_TEXT = {
  [LOG_STATUS.PENDING]: '进行中',
  [LOG_STATUS.COMPLETED]: '已经结束'
}

// 交易日志状态主题映射
export const LOG_STATUS_THEME = {
  [LOG_STATUS.PENDING]: 'warning',
  [LOG_STATUS.COMPLETED]: 'success'
}

// 风险等级
export const RISK_LEVELS = {
  LOW: 'low',
  MEDIUM: 'medium',
  HIGH: 'high'
}

// 风险等级文本映射
export const RISK_LEVEL_TEXT = {
  [RISK_LEVELS.LOW]: '低风险',
  [RISK_LEVELS.MEDIUM]: '中风险',
  [RISK_LEVELS.HIGH]: '高风险'
}

// 选股策略
export const STRATEGIES = {
  TREND_FOLLOWING: 'trend_following',
  BREAKOUT: 'breakout',
  REVERSAL: 'reversal',
  MOMENTUM: 'momentum'
}

// 选股策略文本映射
export const STRATEGY_TEXT = {
  [STRATEGIES.TREND_FOLLOWING]: '趋势跟踪',
  [STRATEGIES.BREAKOUT]: '突破策略',
  [STRATEGIES.REVERSAL]: '反转策略',
  [STRATEGIES.MOMENTUM]: '动量策略'
}

// 交易策略
export const TRADING_STRATEGIES = {
  SCALPING: 'scalping',
  DAY_TRADING: 'daytrading',
  SWING: 'swing',
  POSITION: 'position',
  ARBITRAGE: 'arbitrage'
}

// 交易策略文本映射
export const TRADING_STRATEGY_TEXT = {
  [TRADING_STRATEGIES.SCALPING]: '剥头皮',
  [TRADING_STRATEGIES.DAY_TRADING]: '日内交易',
  [TRADING_STRATEGIES.SWING]: '波段交易',
  [TRADING_STRATEGIES.POSITION]: '持仓交易',
  [TRADING_STRATEGIES.ARBITRAGE]: '套利交易'
}

// 股票数据库
export const STOCK_DATABASE = {
  '000001': '平安银行',
  '000002': '万科A',
  '600036': '招商银行',
  '600519': '贵州茅台',
  '000858': '五粮液',
  '002415': '海康威视',
  '300059': '东方财富',
  '000725': '京东方A'
}

// 分页配置
export const PAGINATION_CONFIG = {
  DEFAULT_PAGE_SIZE: 10,
  PAGE_SIZE_OPTIONS: [10, 20, 50, 100]
}
