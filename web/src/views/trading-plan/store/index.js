// 交易计划页面 Store
import { defineStore } from 'pinia'
import { generateId } from '@/utils/helpers'
import * as tradingPlanApi from '../api' 

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
    
    async loadPlans(params = {}) {
      this.setLoading(true)
      try {
        // 使用真实 API
        const response = await tradingPlanApi.getPlans(params)
        if (response.code === 0) {
          this.plans = response.data.list || []
          // 返回分页信息
          return {
            list: response.data.list || [],
            total: response.data.total || 0,
            page: response.data.page || 1,
            pageSize: response.data.pageSize || 10
          }
        }
      } catch (error) {
        console.error('加载交易计划失败:', error)
        // 降级到本地存储
        this.loadFromLocalStorage()
        return {
          list: this.plans,
          total: this.plans.length,
          page: 1,
          pageSize: 10
        }
      } finally {
        this.setLoading(false)
      }
    },
    
    async addPlan(plan) {
      try {
        const response = await tradingPlanApi.createPlan(plan)
        if (response.code === 0) {
          this.plans.unshift(response.data)
          this.saveToLocalStorage()
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('创建交易计划失败:', error)
        // 不再降级到本地存储，直接抛出错误
        throw error
      }
    },
    
    async updatePlan(planId, updates) {
      try {
        const response = await tradingPlanApi.updatePlan(planId, updates)
        if (response.code === 0) {
          const index = this.plans.findIndex(plan => plan.id === planId)
          if (index !== -1) {
            this.plans[index] = { ...this.plans[index], ...updates }
            this.saveToLocalStorage()
          }
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('更新交易计划失败:', error)
        // 降级到本地更新
        const index = this.plans.findIndex(plan => plan.id === planId)
        if (index !== -1) {
          this.plans[index] = { ...this.plans[index], ...updates }
          this.saveToLocalStorage()
        }
        throw error
      }
    },
    
    async getPlanById(planId) {
      try {
        const response = await tradingPlanApi.getPlanDetail(planId)
        if (response.code === 0) {
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('获取交易计划详情失败:', error)
        // 降级到本地查找
        const plan = this.plans.find(plan => plan.id === planId)
        if (plan) {
          return plan
        }
        throw error
      }
    },
    
    async updatePlanStatus(planId, status) {
      try {
        const response = await tradingPlanApi.updatePlanStatus(planId, status)
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
    
    async deletePlan(planId) {
      try {
        const response = await tradingPlanApi.deletePlan(planId)
        if (response.code === 0) {
          this.plans = this.plans.filter(plan => plan.id !== planId)
          this.saveToLocalStorage()
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('删除交易计划失败:', error)
        // 降级到本地存储
        this.plans = this.plans.filter(plan => plan.id !== planId)
        this.saveToLocalStorage()
        throw error
      }
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
