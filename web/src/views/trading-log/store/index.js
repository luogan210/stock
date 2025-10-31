// 交易日志页面 Store
import { defineStore } from 'pinia'
import { generateId } from '@/utils/helpers'
import * as tradingLogApi from '../api'

export const useTradingLogStore = defineStore('tradingLog', {
  state: () => ({
    logs: [],
    isLoading: false
  }),
  
  getters: {
    getLogs: (state) => state.logs,
    isLogLoading: (state) => state.isLoading,
    getSuccessLogs: (state) => state.logs.filter(log => log.status === 'completed'),
    getFailedLogs: (state) => state.logs.filter(log => log.status === 'failed'),
    getTotalProfit: (state) => state.logs.reduce((sum, log) => sum + (log.profit || 0), 0)
  },
  
  actions: {
    setLoading(loading) {
      this.isLoading = loading
    },
    
    async loadLogs(params = {}) {
      this.setLoading(true)
      try {
        // 使用真实 API
        const response = await tradingLogApi.getLogs(params)
        if (response.code === 0) {
          this.logs = response.data.list || []
          // 返回分页信息
          return {
            list: response.data.list || [],
            total: response.data.total || 0,
            page: response.data.page || 1,
            pageSize: response.data.pageSize || 10
          }
        }
      } catch (error) {
        console.error('加载交易日志失败:', error)
        // 降级到本地存储
        this.loadFromLocalStorage()
        return {
          list: this.logs,
          total: this.logs.length,
          page: 1,
          pageSize: 10
        }
      } finally {
        this.setLoading(false)
      }
    },
    
    async addLog(log) {
      try {
        const response = await tradingLogApi.createLog(log)
        if (response.code === 0) {
          this.logs.unshift(response.data)
          this.saveToLocalStorage()
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('创建交易日志失败:', error)
        throw error
      }
    },
    
    async updateLog(logId, updates) {
      try {
        const response = await tradingLogApi.updateLog(logId, updates)
        if (response.code === 0) {
          const index = this.logs.findIndex(log => log.id === logId)
          if (index !== -1) {
            this.logs[index] = { ...this.logs[index], ...response.data }
            this.saveToLocalStorage()
          }
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('更新交易日志失败:', error)
        // 降级到本地存储
        const index = this.logs.findIndex(log => log.id === logId)
        if (index !== -1) {
          this.logs[index] = { ...this.logs[index], ...updates }
          this.saveToLocalStorage()
        }
        throw error
      }
    },
    
    async getLogById(logId) {
      try {
        const response = await tradingLogApi.getLogDetail(logId)
        if (response.code === 0) {
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('获取交易日志失败:', error)
        // 降级到本地存储
        return this.logs.find(log => log.id === logId)
      }
    },
    
    async deleteLog(logId) {
      try {
        const response = await tradingLogApi.deleteLog(logId)
        if (response.code === 0) {
          this.logs = this.logs.filter(log => log.id !== logId)
          this.saveToLocalStorage()
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('删除交易日志失败:', error)
        // 降级到本地存储
        this.logs = this.logs.filter(log => log.id !== logId)
        this.saveToLocalStorage()
        throw error
      }
    },
    
    saveToLocalStorage() {
      localStorage.setItem('tradingLogs', JSON.stringify(this.logs))
    },
    
    loadFromLocalStorage() {
      const saved = localStorage.getItem('tradingLogs')
      if (saved) {
        try {
          this.logs = JSON.parse(saved)
        } catch (error) {
          console.error('加载本地交易日志失败:', error)
        }
      }
    }
  }
})
