// 交易日志页面 Store
import { defineStore } from 'pinia'
import { generateId } from '@/utils/helpers'
import { currentApi } from '@/services/api'

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
    
    async loadLogs() {
      this.setLoading(true)
      try {
        // 使用当前环境的 API
        const response = await currentApi.tradingLog.getLogs()
        if (response.code === 0) {
          this.logs = response.data.list || []
        }
      } catch (error) {
        console.error('加载交易日志失败:', error)
        // 降级到本地存储
        this.loadFromLocalStorage()
      } finally {
        this.setLoading(false)
      }
    },
    
    async addLog(log) {
      try {
        const response = await currentApi.tradingLog.createLog(log)
        if (response.code === 0) {
          this.logs.unshift(response.data)
          this.saveToLocalStorage()
          return response.data
        }
        throw new Error(response.message)
      } catch (error) {
        console.error('创建交易日志失败:', error)
        // 降级到本地存储
        const newLog = {
          ...log,
          id: log.id || generateId(),
          createTime: log.createTime || new Date().toLocaleString(),
          updateTime: new Date().toLocaleString()
        }
        this.logs.unshift(newLog)
        this.saveToLocalStorage()
        return newLog
      }
    },
    
    updateLog(logId, updates) {
      const index = this.logs.findIndex(log => log.id === logId)
      if (index !== -1) {
        this.logs[index] = { ...this.logs[index], ...updates }
        this.saveToLocalStorage()
      }
    },
    
    deleteLog(logId) {
      this.logs = this.logs.filter(log => log.id !== logId)
      this.saveToLocalStorage()
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
