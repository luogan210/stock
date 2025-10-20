// 股票信息 Store
import { defineStore } from 'pinia'
import { STOCK_DATABASE } from '@/utils/constants'

export const useStockStore = defineStore('stock', {
  state: () => ({
    stockList: Object.entries(STOCK_DATABASE).map(([code, name], index) => ({
      id: String(index + 1),
      code,
      name,
      price: Math.random() * 100 + 10,
      change: (Math.random() - 0.5) * 10,
      changePercent: (Math.random() - 0.5) * 20,
      volume: Math.floor(Math.random() * 1000000),
      marketCap: Math.floor(Math.random() * 1000000000),
      updateTime: new Date().toISOString()
    })),
    watchList: [],
    selectedStock: null
  }),
  
  getters: {
    getStockList: (state) => state.stockList,
    getWatchList: (state) => state.watchList,
    getSelectedStock: (state) => state.selectedStock,
    getStockByCode: (state) => (code) => state.stockList.find(stock => stock.code === code)
  },
  
  actions: {
    addToWatchList(stock) {
      if (!this.watchList.find(item => item.code === stock.code)) {
        this.watchList.push(stock)
      }
    },
    
    removeFromWatchList(code) {
      this.watchList = this.watchList.filter(item => item.code !== code)
    },
    
    setSelectedStock(stock) {
      this.selectedStock = stock
    },
    
    updateStockPrice(code, priceData) {
      const stock = this.stockList.find(s => s.code === code)
      if (stock) {
        Object.assign(stock, priceData)
        stock.updateTime = new Date().toISOString()
      }
    }
  }
})
