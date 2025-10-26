/**
 * 风险等级计算器
 * 根据选股策略和交易策略计算综合风险等级
 */

import { RISK_LEVELS } from './strategyConfig'
import { getStrategyById } from './strategyConfig'
import { getTradingStrategyById } from './tradingStrategyConfig'

// 风险等级权重映射
const RISK_WEIGHTS = {
  [RISK_LEVELS.LOW]: 1,
  [RISK_LEVELS.MEDIUM]: 2,
  [RISK_LEVELS.HIGH]: 3
}

// 风险等级文本映射
export const RISK_LEVEL_TEXT = {
  [RISK_LEVELS.LOW]: '低风险',
  [RISK_LEVELS.MEDIUM]: '中风险',
  [RISK_LEVELS.HIGH]: '高风险'
}

// 风险等级颜色映射
export const RISK_LEVEL_COLORS = {
  [RISK_LEVELS.LOW]: '#52c41a',    // 绿色
  [RISK_LEVELS.MEDIUM]: '#faad14', // 橙色
  [RISK_LEVELS.HIGH]: '#ff4d4f'    // 红色
}

/**
 * 根据选股策略和交易策略计算综合风险等级
 * @param {string} strategyId - 选股策略ID
 * @param {string} tradingStrategyId - 交易策略ID
 * @returns {object} 风险等级信息
 */
export const calculateRiskLevel = (strategyId, tradingStrategyId) => {
  // 获取策略详情
  const strategy = getStrategyById(strategyId)
  const tradingStrategy = getTradingStrategyById(tradingStrategyId)
  
  if (!strategy || !tradingStrategy) {
    return {
      level: RISK_LEVELS.MEDIUM,
      text: RISK_LEVEL_TEXT[RISK_LEVELS.MEDIUM],
      color: RISK_LEVEL_COLORS[RISK_LEVELS.MEDIUM],
      score: 2,
      details: {
        strategyRisk: null,
        tradingStrategyRisk: null,
        combinedScore: 2,
        explanation: '策略信息不完整，默认中风险'
      }
    }
  }

  // 获取各策略的风险等级
  const strategyRisk = strategy.parameters?.riskLevel || RISK_LEVELS.MEDIUM
  const tradingStrategyRisk = tradingStrategy.parameters?.riskLevel || RISK_LEVELS.MEDIUM

  // 计算风险权重分数
  const strategyWeight = RISK_WEIGHTS[strategyRisk] || 2
  const tradingStrategyWeight = RISK_WEIGHTS[tradingStrategyRisk] || 2

  // 计算综合风险分数 (选股策略权重40%，交易策略权重60%)
  const combinedScore = (strategyWeight * 0.4) + (tradingStrategyWeight * 0.6)

  // 确定最终风险等级
  let finalRiskLevel
  if (combinedScore <= 1.5) {
    finalRiskLevel = RISK_LEVELS.LOW
  } else if (combinedScore <= 2.5) {
    finalRiskLevel = RISK_LEVELS.MEDIUM
  } else {
    finalRiskLevel = RISK_LEVELS.HIGH
  }

  // 生成风险说明
  const explanation = generateRiskExplanation(strategy, tradingStrategy, strategyRisk, tradingStrategyRisk, finalRiskLevel)

  return {
    level: finalRiskLevel,
    text: RISK_LEVEL_TEXT[finalRiskLevel],
    color: RISK_LEVEL_COLORS[finalRiskLevel],
    score: Math.round(combinedScore * 10) / 10, // 保留1位小数
    details: {
      strategyRisk,
      tradingStrategyRisk,
      combinedScore: Math.round(combinedScore * 10) / 10,
      explanation
    }
  }
}

/**
 * 生成风险说明
 */
const generateRiskExplanation = (strategy, tradingStrategy, strategyRisk, tradingStrategyRisk, finalRisk) => {
  const strategyName = strategy.name
  const tradingStrategyName = tradingStrategy.name
  const strategyRiskText = RISK_LEVEL_TEXT[strategyRisk]
  const tradingStrategyRiskText = RISK_LEVEL_TEXT[tradingStrategyRisk]
  const finalRiskText = RISK_LEVEL_TEXT[finalRisk]

  let explanation = `选股策略"${strategyName}"为${strategyRiskText}，交易策略"${tradingStrategyName}"为${tradingStrategyRiskText}。`

  // 添加具体的风险分析
  if (finalRisk === RISK_LEVELS.LOW) {
    explanation += '综合评估为低风险，适合稳健型投资者。建议严格执行止损止盈策略。'
  } else if (finalRisk === RISK_LEVELS.MEDIUM) {
    explanation += '综合评估为中风险，适合有一定经验的投资者。需要密切关注市场变化，及时调整策略。'
  } else {
    explanation += '综合评估为高风险，适合有丰富经验的投资者。建议控制仓位，设置严格的风控措施。'
  }

  return explanation
}

/**
 * 获取风险等级建议
 * @param {string} riskLevel - 风险等级
 * @returns {object} 风险建议
 */
export const getRiskSuggestions = (riskLevel) => {
  const suggestions = {
    [RISK_LEVELS.LOW]: {
      positionSize: '建议仓位：30-50%',
      stopLoss: '建议止损：3-5%',
      takeProfit: '建议止盈：8-15%',
      timeframe: '持仓周期：中长期',
      tips: [
        '适合新手投资者',
        '严格执行止损策略',
        '分批建仓降低风险',
        '关注基本面变化'
      ]
    },
    [RISK_LEVELS.MEDIUM]: {
      positionSize: '建议仓位：20-40%',
      stopLoss: '建议止损：5-8%',
      takeProfit: '建议止盈：15-25%',
      timeframe: '持仓周期：短中期',
      tips: [
        '需要一定交易经验',
        '密切关注技术指标',
        '灵活调整策略',
        '控制情绪波动'
      ]
    },
    [RISK_LEVELS.HIGH]: {
      positionSize: '建议仓位：10-30%',
      stopLoss: '建议止损：8-12%',
      takeProfit: '建议止盈：25-50%',
      timeframe: '持仓周期：短期',
      tips: [
        '需要丰富交易经验',
        '严格控制仓位',
        '快速止损止盈',
        '避免重仓操作',
        '保持冷静心态'
      ]
    }
  }

  return suggestions[riskLevel] || suggestions[RISK_LEVELS.MEDIUM]
}

/**
 * 验证风险等级是否合理
 * @param {number} targetPrice - 目标价格
 * @param {number} stopLoss - 止损价格
 * @param {number} takeProfit - 止盈价格
 * @param {string} riskLevel - 风险等级
 * @returns {object} 验证结果
 */
export const validateRiskParameters = (targetPrice, stopLoss, takeProfit, riskLevel) => {
  if (!targetPrice || !stopLoss || !takeProfit) {
    return {
      valid: false,
      message: '价格参数不完整'
    }
  }

  const stopLossRatio = Math.abs((targetPrice - stopLoss) / targetPrice) * 100
  const takeProfitRatio = Math.abs((takeProfit - targetPrice) / targetPrice) * 100

  const suggestions = getRiskSuggestions(riskLevel)
  
  // 解析建议的止损止盈范围
  const stopLossRange = suggestions.stopLoss.match(/(\d+)-(\d+)%/)
  const takeProfitRange = suggestions.takeProfit.match(/(\d+)-(\d+)%/)

  let warnings = []

  if (stopLossRange) {
    const minStopLoss = parseFloat(stopLossRange[1])
    const maxStopLoss = parseFloat(stopLossRange[2])
    
    if (stopLossRatio < minStopLoss) {
      warnings.push(`止损比例${stopLossRatio.toFixed(1)}%偏小，建议${minStopLoss}-${maxStopLoss}%`)
    } else if (stopLossRatio > maxStopLoss) {
      warnings.push(`止损比例${stopLossRatio.toFixed(1)}%偏大，建议${minStopLoss}-${maxStopLoss}%`)
    }
  }

  if (takeProfitRange) {
    const minTakeProfit = parseFloat(takeProfitRange[1])
    const maxTakeProfit = parseFloat(takeProfitRange[2])
    
    if (takeProfitRatio < minTakeProfit) {
      warnings.push(`止盈比例${takeProfitRatio.toFixed(1)}%偏小，建议${minTakeProfit}-${maxTakeProfit}%`)
    } else if (takeProfitRatio > maxTakeProfit) {
      warnings.push(`止盈比例${takeProfitRatio.toFixed(1)}%偏大，建议${minTakeProfit}-${maxTakeProfit}%`)
    }
  }

  return {
    valid: warnings.length === 0,
    warnings,
    stopLossRatio: stopLossRatio.toFixed(1),
    takeProfitRatio: takeProfitRatio.toFixed(1)
  }
}
