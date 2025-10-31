import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { stockApi } from '../api'
import { MessagePlugin } from 'tdesign-vue-next'

export const useStockStore = defineStore('stock', () => {
  // 状态
  const stocks = ref([])
  const loading = ref(false)
  const currentStock = ref(null)

  // 计算属性
  const getStocks = computed(() => stocks.value || [])
  const isLoading = computed(() => loading.value)

  // 获取股票列表
  const loadStocks = async (params = {}) => {
    loading.value = true
    try {
      const response = await stockApi.getStocks(params)
      if (response.data && response.data.list) {
        stocks.value = response.data.list
      } else {
        stocks.value = []
      }
    } catch (error) {
      console.error('获取股票列表失败:', error)
      MessagePlugin.error('获取股票列表失败: ' + (error.message || '请重试'))
      stocks.value = []
    } finally {
      loading.value = false
    }
  }

  // 创建股票
  const addStock = async (stockData) => {
    try {
      const response = await stockApi.createStock(stockData)
      if (response.data && response.data.id) {
        // 重新加载列表
        await loadStocks()
        return response.data
      }
    } catch (error) {
      console.error('创建股票失败:', error)
      throw new Error(error.message || '创建股票失败')
    }
  }

  // 更新股票
  const updateStockById = async (id, stockData) => {
    try {
      const response = await stockApi.updateStock(id, stockData)
      if (response.data) {
        // 重新加载列表
        await loadStocks()
        return response.data
      }
    } catch (error) {
      console.error('更新股票失败:', error)
      throw new Error(error.message || '更新股票失败')
    }
  }

  // 删除股票
  const removeStock = async (id) => {
    try {
      await stockApi.deleteStock(id)
      // 重新加载列表
      await loadStocks()
    } catch (error) {
      console.error('删除股票失败:', error)
      throw new Error(error.message || '删除股票失败')
    }
  }

  // 获取股票详情
  const getStockById = async (id) => {
    try {
      const response = await stockApi.getStockDetail(id)
      if (response.data) {
        currentStock.value = response.data
        return response.data
      }
    } catch (error) {
      console.error('获取股票详情失败:', error)
      throw new Error(error.message || '获取股票详情失败')
    }
  }

  // 根据ID查找股票
  const findStockById = (id) => {
    return (stocks.value || []).find(stock => stock.id === id)
  }

  // 重置状态
  const resetState = () => {
    stocks.value = []
    currentStock.value = null
    loading.value = false
  }
  loadStocks()
  return {
    // 状态
    stocks,
    loading,
    currentStock,

    // 计算属性
    getStocks,
    isLoading,

    // 方法
    loadStocks,
    addStock,
    updateStockById,
    removeStock,
    getStockById,
    findStockById,
    resetState
  }
})

