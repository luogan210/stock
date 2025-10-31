// 首页 Store
import { defineStore } from 'pinia'
import * as homeApi from '../api'

export const useHomeStore = defineStore('home', {
  state: () => ({
    homeStats: {
      stockCount: 0,
      watchCount: 0,
      planCount: 0,
      activePlanCount: 0,
      logCount: 0,
      totalProfit: 0,
      todayTradingCount: 0,
      weekTradingCount: 0,
      monthTradingCount: 0
    },
    isLoading: false,
    quickActions: [],
    recentTrades: [],
    marketOverview: null
  }),
  
  getters: {
    getHomeStats: (state) => state.homeStats,
    isHomeLoading: (state) => state.isLoading,
    getQuickActions: (state) => state.quickActions,
    getRecentTrades: (state) => state.recentTrades,
    getMarketOverview: (state) => state.marketOverview
  },
  
  actions: {
    setLoading(loading) {
      this.isLoading = loading
    },
    
    async loadHomeStats() {
      this.setLoading(true)
      try {
        const response = await homeApi.getHomeStats()
        if (response.code === 0) {
          this.homeStats = response.data
        }
      } catch (error) {
        console.error('加载首页统计数据失败:', error)
      } finally {
        this.setLoading(false)
      }
    },
    
    async loadQuickActions() {
      try {
        const response = await homeApi.getQuickActions()
        if (response.code === 0) {
          this.quickActions = response.data
        }
      } catch (error) {
        console.error('加载快速操作失败:', error)
      }
    },
    
    async loadRecentTrades(limit = 5) {
      try {
        const response = await homeApi.getRecentTrades(limit)
        if (response.code === 0) {
          this.recentTrades = response.data
        }
      } catch (error) {
        console.error('加载最近交易失败:', error)
      }
    },
    
    async loadMarketOverview() {
      try {
        const response = await homeApi.getMarketOverview()
        if (response.code === 0) {
          this.marketOverview = response.data
        }
      } catch (error) {
        console.error('加载市场概览失败:', error)
      }
    },
    
    async loadAllData() {
      await Promise.all([
        this.loadHomeStats(),
        this.loadQuickActions(),
        this.loadRecentTrades(),
        this.loadMarketOverview()
      ])
    }
  }
})
