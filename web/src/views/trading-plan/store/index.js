// 交易计划页面 Store
import { defineStore } from 'pinia'
import { generateId } from '@/utils/helpers'
import { currentApi } from '@/services/api'

export const useTradingPlanStore = defineStore('tradingPlan', {
  state: () => ({
    plans: [],
    isLoading: false
  }),
  
  getters: {
    getPlans: (state) => state.plans,
    isPlanLoading: (state) => state.isLoading,
    getActivePlans: (state) => state.plans.filter(plan => plan.status === 'active'),
    getCompletedPlans: (state) => state.plans.filter(plan => plan.status === 'completed')
  },
  
  actions: {
    setLoading(loading) {
      this.isLoading = loading
    },
    
    async loadPlans() {
      this.setLoading(true)
      try {
        // 使用当前环境的 API
        const response = await currentApi.tradingPlan.getPlans()
        if (response.code === 0) {
          this.plans = response.data.list || []
        }
      } catch (error) {
        console.error('加载交易计划失败:', error)
        // 降级到本地存储
        this.loadFromLocalStorage()
      } finally {
        this.setLoading(false)
      }
    },
    
    async addPlan(plan) {
      try {
        const response = await currentApi.tradingPlan.createPlan(plan)
        if (response.code === 0) {
          this.plans.unshift(response.data)
          this.saveToLocalStorage()
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('创建交易计划失败:', error)
        // 降级到本地存储
        const newPlan = {
          ...plan,
          id: plan.id || generateId(),
          createTime: plan.createTime || new Date().toLocaleString(),
          updateTime: new Date().toLocaleString()
        }
        this.plans.unshift(newPlan)
        this.saveToLocalStorage()
        return newPlan
      }
    },
    
    updatePlan(planId, updates) {
      const index = this.plans.findIndex(plan => plan.id === planId)
      if (index !== -1) {
        this.plans[index] = { ...this.plans[index], ...updates }
        this.saveToLocalStorage()
      }
    },
    
    async updatePlanStatus(planId, status) {
      try {
        const response = await currentApi.tradingPlan.updatePlanStatus(planId, status)
        if (response.code === 0) {
          const plan = this.plans.find(plan => plan.id === planId)
          if (plan) {
            plan.status = status
            plan.updateTime = new Date().toLocaleString()
            if (status === 'completed') {
              plan.progress = 100
            }
            this.saveToLocalStorage()
          }
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('更新计划状态失败:', error)
        // 降级到本地更新
        const plan = this.plans.find(plan => plan.id === planId)
        if (plan) {
          plan.status = status
          plan.updateTime = new Date().toLocaleString()
          if (status === 'completed') {
            plan.progress = 100
          }
          this.saveToLocalStorage()
        }
      }
    },
    
    deletePlan(planId) {
      this.plans = this.plans.filter(plan => plan.id !== planId)
      this.saveToLocalStorage()
    },
    
    saveToLocalStorage() {
      localStorage.setItem('tradingPlans', JSON.stringify(this.plans))
    },
    
    loadFromLocalStorage() {
      const saved = localStorage.getItem('tradingPlans')
      if (saved) {
        try {
          this.plans = JSON.parse(saved)
        } catch (error) {
          console.error('加载本地交易计划失败:', error)
        }
      }
    }
  }
})
