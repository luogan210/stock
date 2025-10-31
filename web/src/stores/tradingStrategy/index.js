/**
 * 交易策略 Store - 简化版
 * 只负责交易策略数据存储
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { tradingStrategyConfig } from '@/utils/tradingStrategyConfig'

export const useTradingStrategyStore = defineStore('tradingStrategy', () => {
  // 交易策略数据
  const tradingStrategies = ref([...tradingStrategyConfig])

  // 添加交易策略
  const addStrategy = (strategy) => {
    const newStrategy = {
      id: `trading_strategy_${Date.now()}`,
      enabled: true,
      ...strategy
    }
    tradingStrategies.value.push(newStrategy)
    return newStrategy
  }

  // 更新交易策略
  const updateStrategy = (id, updates) => {
    const index = tradingStrategies.value.findIndex(strategy => strategy.id === id)
    if (index !== -1) {
      tradingStrategies.value[index] = { ...tradingStrategies.value[index], ...updates }
      return true
    }
    return false
  }

  // 删除交易策略
  const removeStrategy = (id) => {
    const index = tradingStrategies.value.findIndex(strategy => strategy.id === id)
    if (index !== -1) {
      tradingStrategies.value.splice(index, 1)
      return true
    }
    return false
  }

  // 获取交易策略
  const getStrategyById = (id) => {
    return tradingStrategies.value.find(strategy => strategy.id === id)
  }

  // 获取启用的策略
  const getEnabledStrategies = () => {
    return tradingStrategies.value.filter(strategy => strategy.enabled)
  }

  // 根据分类获取策略
  const getStrategiesByCategory = (category) => {
    return tradingStrategies.value.filter(strategy => strategy.category === category)
  }

  return {
    // 数据
    tradingStrategies,

    // 方法
    addStrategy,
    updateStrategy,
    removeStrategy,
    getStrategyById,
    getEnabledStrategies,
    getStrategiesByCategory
  }
})
