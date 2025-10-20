// 通用工具函数

/**
 * 格式化日期
 * @param {Date|string} date - 日期
 * @param {string} format - 格式
 * @returns {string} 格式化后的日期字符串
 */
export function formatDate(date, format = 'YYYY-MM-DD HH:mm:ss') {
  if (!date) return ''
  
  const d = new Date(date)
  if (isNaN(d.getTime())) return ''
  
  const year = d.getFullYear()
  const month = String(d.getMonth() + 1).padStart(2, '0')
  const day = String(d.getDate()).padStart(2, '0')
  const hours = String(d.getHours()).padStart(2, '0')
  const minutes = String(d.getMinutes()).padStart(2, '0')
  const seconds = String(d.getSeconds()).padStart(2, '0')
  
  return format
    .replace('YYYY', year)
    .replace('MM', month)
    .replace('DD', day)
    .replace('HH', hours)
    .replace('mm', minutes)
    .replace('ss', seconds)
}

/**
 * 格式化货币
 * @param {number} amount - 金额
 * @param {string} currency - 货币符号
 * @returns {string} 格式化后的货币字符串
 */
export function formatCurrency(amount, currency = '¥') {
  if (amount === null || amount === undefined) return '-'
  return `${currency}${Number(amount).toFixed(2)}`
}

/**
 * 格式化盈亏显示
 * @param {number} profit - 盈亏金额
 * @returns {object} 包含显示文本和样式的对象
 */
export function formatProfit(profit) {
  if (profit === null || profit === undefined) {
    return { text: '-', class: 'profit-neutral' }
  }
  
  const num = Number(profit)
  if (num > 0) {
    return { text: `+${num.toFixed(2)}`, class: 'profit-positive' }
  } else if (num < 0) {
    return { text: num.toFixed(2), class: 'profit-negative' }
  } else {
    return { text: '0.00', class: 'profit-neutral' }
  }
}

/**
 * 获取日期范围
 * @param {string} type - 类型: 'today', 'week', 'month'
 * @returns {Array} 日期范围数组
 */
export function getDateRange(type) {
  const now = new Date()
  
  switch (type) {
    case 'today':
      const today = new Date(now)
      today.setHours(0, 0, 0, 0)
      const endOfToday = new Date(now)
      endOfToday.setHours(23, 59, 59, 999)
      return [today, endOfToday]
      
    case 'week':
      const startOfWeek = new Date(now)
      startOfWeek.setDate(now.getDate() - now.getDay())
      startOfWeek.setHours(0, 0, 0, 0)
      const endOfWeek = new Date(startOfWeek)
      endOfWeek.setDate(startOfWeek.getDate() + 6)
      endOfWeek.setHours(23, 59, 59, 999)
      return [startOfWeek, endOfWeek]
      
    case 'month':
      const startOfMonth = new Date(now.getFullYear(), now.getMonth(), 1)
      const endOfMonth = new Date(now.getFullYear(), now.getMonth() + 1, 0, 23, 59, 59, 999)
      return [startOfMonth, endOfMonth]
      
    default:
      return []
  }
}

/**
 * 过滤数据
 * @param {Array} data - 数据数组
 * @param {Object} filters - 过滤条件
 * @returns {Array} 过滤后的数据
 */
export function filterData(data, filters) {
  return data.filter(item => {
    // 关键词搜索
    if (filters.keyword) {
      const keyword = filters.keyword.toLowerCase()
      const searchFields = ['name', 'title', 'stockCode', 'stockName', 'content']
      const hasMatch = searchFields.some(field => {
        const value = item[field]
        return value && value.toString().toLowerCase().includes(keyword)
      })
      if (!hasMatch) return false
    }
    
    // 状态过滤
    if (filters.status && item.status !== filters.status) {
      return false
    }
    
    // 类型过滤
    if (filters.type && item.type !== filters.type) {
      return false
    }
    
    // 日期范围过滤
    if (filters.dateRange && filters.dateRange.length === 2) {
      const [startDate, endDate] = filters.dateRange
      const itemDate = new Date(item.tradingTime || item.createTime)
      if (itemDate < startDate || itemDate > endDate) {
        return false
      }
    }
    
    return true
  })
}

/**
 * 生成唯一ID
 * @returns {string} 唯一ID
 */
export function generateId() {
  return Date.now().toString() + Math.random().toString(36).substr(2, 9)
}

/**
 * 深拷贝对象
 * @param {any} obj - 要拷贝的对象
 * @returns {any} 拷贝后的对象
 */
export function deepClone(obj) {
  if (obj === null || typeof obj !== 'object') return obj
  if (obj instanceof Date) return new Date(obj.getTime())
  if (obj instanceof Array) return obj.map(item => deepClone(item))
  if (typeof obj === 'object') {
    const clonedObj = {}
    for (const key in obj) {
      if (obj.hasOwnProperty(key)) {
        clonedObj[key] = deepClone(obj[key])
      }
    }
    return clonedObj
  }
}

/**
 * 防抖函数
 * @param {Function} func - 要防抖的函数
 * @param {number} wait - 等待时间
 * @returns {Function} 防抖后的函数
 */
export function debounce(func, wait) {
  let timeout
  return function executedFunction(...args) {
    const later = () => {
      clearTimeout(timeout)
      func(...args)
    }
    clearTimeout(timeout)
    timeout = setTimeout(later, wait)
  }
}

/**
 * 节流函数
 * @param {Function} func - 要节流的函数
 * @param {number} limit - 限制时间
 * @returns {Function} 节流后的函数
 */
export function throttle(func, limit) {
  let inThrottle
  return function executedFunction(...args) {
    if (!inThrottle) {
      func.apply(this, args)
      inThrottle = true
      setTimeout(() => inThrottle = false, limit)
    }
  }
}
