// 股票信息 Store - 从股票管理模块获取数据
import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { useStockStore as useStockManagementStore } from '@/views/stock/store'

export const useStockStore = defineStore('stock', () => {
  // 使用股票管理模块的store
  const stockManagementStore = useStockManagementStore()
  
  // 本地状态
  const watchList = ref([])
  const selectedStock = ref(null)
  
  // 计算属性 - 从股票管理模块获取数据
  const stockList = computed(() => {
    const stocks = stockManagementStore.getStocks || []
    return stocks.map(stock => ({
      id: stock.id,
      code: stock.code,
      name: stock.name,
      region: stock.region,
      currency: stock.currency,
      category: stock.category,
      enabled: stock.enabled,
      remark: stock.remark,
      // 添加一些模拟的实时数据（实际项目中这些数据应该来自实时API）
      price: Math.random() * 100 + 10,
      change: (Math.random() - 0.5) * 10,
      changePercent: (Math.random() - 0.5) * 20,
      volume: Math.floor(Math.random() * 1000000),
      marketCap: Math.floor(Math.random() * 1000000000),
      updateTime: new Date().toISOString()
    }))
  })
  
  // 计算属性
  const getStockList = computed(() => stockList.value)
  const getWatchList = computed(() => watchList.value)
  const getSelectedStock = computed(() => selectedStock.value)
  
  // 方法
  const getStockByCode = (code) => {
    return stockList.value.find(stock => stock.code === code)
  }
  
  const addToWatchList = (stock) => {
    if (!watchList.value.find(item => item.code === stock.code)) {
      watchList.value.push(stock)
    }
  }
  
  const removeFromWatchList = (code) => {
    watchList.value = watchList.value.filter(item => item.code !== code)
  }
  
  const setSelectedStock = (stock) => {
    selectedStock.value = stock
  }
  
  const updateStockPrice = (code, priceData) => {
    const stock = stockList.value.find(s => s.code === code)
    if (stock) {
      Object.assign(stock, priceData)
      stock.updateTime = new Date().toISOString()
    }
  }
  
  // 加载股票数据
  const loadStocks = async () => {
    await stockManagementStore.loadStocks()
  }
  
  return {
    // 状态
    stockList,
    watchList,
    selectedStock,
    
    // 计算属性
    getStockList,
    getWatchList,
    getSelectedStock,
    
    // 方法
    getStockByCode,
    addToWatchList,
    removeFromWatchList,
    setSelectedStock,
    updateStockPrice,
    loadStocks
  }
})
