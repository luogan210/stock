// 交易日志状态
export const LOG_STATUS = [
    { id: 'pending', text: '进行中', theme: 'warning' },
    { id: 'completed', text: '已经结束', theme: 'success' }
]

// 交易状态
export const TRADING_STATUS = [
    { id: 'pending', text: '待执行', theme: 'warning' },
    { id: 'executing', text: '执行中', theme: 'primary' },
    { id: 'completed', text: '已完成', theme: 'success' },
    { id: 'cancelled', text: '已取消', theme: 'default' },
    { id: 'failed', text: '执行失败', theme: 'danger' }
]

export function getTradingLogStatusText(status) {
    return LOG_STATUS.find(item => item.id === status)?.text || status
}

export function getTradingLogStatusTheme(status) {
    return LOG_STATUS.find(item => item.id === status)?.theme || 'default'
}

export function getTradingStatusText(status) {
    return TRADING_STATUS.find(item => item.id === status)?.text || status
}

export function getTradingStatusTheme(status) {
    return TRADING_STATUS.find(item => item.id === status)?.theme || 'default'
}