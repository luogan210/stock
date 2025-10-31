/**
 * 交易复盘 Store
 * 管理交易复盘相关的状态和操作
 */

import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import * as tradingReviewApi from '../api'

export const useTradingReviewStore = defineStore('tradingReview', () => {
  // 状态
  const reviews = ref([])
  const loading = ref(false)
  const reviewStats = ref({
    totalReviews: 0,
    weeklyReviews: 0,
    monthlyReviews: 0,
    avgWinRate: 0,
    totalProfit: 0,
    recentReviews: []
  })

  // 计算属性
  const getReviews = computed(() => reviews.value)
  const getReviewStats = computed(() => reviewStats.value)
  const isLoading = computed(() => loading.value)

  // 获取复盘列表
  const loadReviews = async (params = {}) => {
    try {
      loading.value = true
      const response = await tradingReviewApi.getReviews(params)

      if (response.code === 0) {
        reviews.value = response.data.list || []
        return response.data
      } else {
        throw new Error(response.message || '获取复盘列表失败')
      }
    } catch (error) {
      console.error('加载复盘列表失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 获取复盘详情
  const loadReview = async (id) => {
    try {
      loading.value = true
      const response = await tradingReviewApi.getReview(id)

      if (response.code === 0) {
        return response.data
      } else {
        throw new Error(response.message || '获取复盘详情失败')
      }
    } catch (error) {
      console.error('加载复盘详情失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 创建复盘
  const addReview = async (reviewData) => {
    try {
      loading.value = true
      const response = await tradingReviewApi.createReview(reviewData)

      if (response.code === 0) {
        reviews.value.unshift(response.data)
        return response.data
      } else {
        throw new Error(response.message || '创建复盘失败')
      }
    } catch (error) {
      console.error('创建复盘失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 更新复盘
  const updateReview = async (id, reviewData) => {
    try {
      loading.value = true
      const response = await tradingReviewApi.updateReview(id, reviewData)

      if (response.code === 0) {
        const index = reviews.value.findIndex(r => r.id === id)
        if (index !== -1) {
          reviews.value[index] = response.data
        }
        return response.data
      } else {
        throw new Error(response.message || '更新复盘失败')
      }
    } catch (error) {
      console.error('更新复盘失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 删除复盘
  const deleteReview = async (id) => {
    try {
      loading.value = true
      const response = await tradingReviewApi.deleteReview(id)

      if (response.code === 0) {
        const index = reviews.value.findIndex(r => r.id === id)
        if (index !== -1) {
          reviews.value.splice(index, 1)
        }
        return true
      } else {
        throw new Error(response.message || '删除复盘失败')
      }
    } catch (error) {
      console.error('删除复盘失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 加载复盘统计
  const loadReviewStats = async () => {
    try {
      loading.value = true
      const response = await currentApi.tradingReview.getReviewStats()

      if (response.code === 0) {
        reviewStats.value = response.data
        return response.data
      } else {
        throw new Error(response.message || '获取复盘统计失败')
      }
    } catch (error) {
      console.error('加载复盘统计失败:', error)
      throw error
    } finally {
      loading.value = false
    }
  }

  // 根据ID获取复盘
  const getReviewById = (id) => {
    return reviews.value.find(r => r.id === id)
  }

  // 根据周期筛选复盘
  const getReviewsByPeriod = (period) => {
    return reviews.value.filter(r => r.period === period)
  }

  // 根据状态筛选复盘
  const getReviewsByStatus = (status) => {
    return reviews.value.filter(r => r.status === status)
  }

  // 获取最近复盘
  const getRecentReviews = (limit = 5) => {
    return reviews.value.slice(0, limit)
  }

  // 清空数据
  const clearReviews = () => {
    reviews.value = []
    reviewStats.value = {
      totalReviews: 0,
      weeklyReviews: 0,
      monthlyReviews: 0,
      avgWinRate: 0,
      totalProfit: 0,
      recentReviews: []
    }
  }

  return {
    // 状态
    reviews,
    loading,
    reviewStats,

    // 计算属性
    getReviews,
    getReviewStats,
    isLoading,

    // 方法
    loadReviews,
    loadReview,
    addReview,
    updateReview,
    deleteReview,
    loadReviewStats,
    getReviewById,
    getReviewsByPeriod,
    getReviewsByStatus,
    getRecentReviews,
    clearReviews
  }
})
