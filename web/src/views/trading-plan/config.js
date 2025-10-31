export const PLAN_STATUS = [
    { id: 'pending', text: '未执行', theme: 'default' },
    { id: 'executing', text: '执行中', theme: 'primary' },
    { id: 'completed', text: '已完成', theme: 'success' }
]

export function getTradingPlanStatusText(status) {
    return PLAN_STATUS.find(item => item.id === status)?.text || status
}

export function getTradingPlanStatusTheme(status) {
    return PLAN_STATUS.find(item => item.id === status)?.theme || 'default'
}
