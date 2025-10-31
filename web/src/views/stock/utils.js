import { stockCategoryConfig } from './config'
// 获取股票分类文本
export const getStockCategoryText = (category) => {
    return stockCategoryConfig.find(item => item.id === category)?.name || '未知'    
}