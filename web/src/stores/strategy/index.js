/**
 * 选股策略 Store - 简化版
 * 只负责选股策略数据存储
 */
import { defineStore } from 'pinia'
import { ref } from 'vue'
import { strategyConfig } from '@/utils/strategyConfig'

export const useStrategyStore = defineStore('strategy', () => {
  // 选股策略数据
  const strategies = ref([...strategyConfig])

  // 添加选股策略
  const addStrategy = (strategy) => {
    const newStrategy = {
      id: `strategy_${Date.now()}`,
      enabled: true,
      ...strategy
    }
    strategies.value.push(newStrategy)
    return newStrategy
  }

  // 更新选股策略
  const updateStrategy = (id, updates) => {
    const index = strategies.value.findIndex(strategy => strategy.id === id)
    if (index !== -1) {
      strategies.value[index] = { ...strategies.value[index], ...updates }
      return true
    }
    return false
  }

  // 删除选股策略
  const removeStrategy = (id) => {
    const index = strategies.value.findIndex(strategy => strategy.id === id)
    if (index !== -1) {
      strategies.value.splice(index, 1)
      return true
    }
    return false
  }

  // 获取选股策略
  const getStrategyById = (id) => {
    return strategies.value.find(strategy => strategy.id === id)
  }

  // 获取启用的策略
  const getEnabledStrategies = () => {
    return strategies.value.filter(strategy => strategy.enabled)
  }

  // 根据分类获取策略
  const getStrategiesByCategory = (category) => {
    return strategies.value.filter(strategy => strategy.category === category)
  }

  return {
    // 数据
    strategies,
    // 方法
    addStrategy,
    updateStrategy,
    removeStrategy,
    getStrategyById,
    getEnabledStrategies,
    getStrategiesByCategory
  }
})