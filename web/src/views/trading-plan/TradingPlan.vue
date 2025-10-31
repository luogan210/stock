<template>
  <div class="trading-plan">
    <t-card title="交易计划" class="plan-card">
      <template #header>
        <div class="card-header">
          <t-icon name="chart" size="24px" />
          <span>我的交易计划</span>
          <t-button theme="primary" @click="goToCreatePlan" class="create-btn" style="pointer-events: auto;">
            <template #icon>
              <t-icon name="add" />
            </template>
            新建计划
          </t-button>
        </div>
      </template>
      
      <div class="plan-content">
        <!-- 筛选和搜索 -->
        <t-form :model="searchForm" @submit="handleSearch" layout="inline" label-width="80px" class="search-form">
          <t-form-item label="关键词" name="keyword">
            <t-input
              class="form-input-md"
              v-model="searchForm.keyword"
              placeholder="搜索股票代码或名称"
              clearable
            >
              <template #prefix-icon>
                <t-icon name="search" />
              </template>
            </t-input>
          </t-form-item>
          <t-form-item label="状态" name="status">
            <t-select class="form-input-md" v-model="searchForm.status" placeholder="状态筛选">
              <t-option value="" label="全部状态" />
              <t-option v-for="status in PLAN_STATUS" :key="status" :value="status" :label="getTradingPlanStatusText(status)" />
            </t-select>
          </t-form-item>
          <t-form-item label="方向" name="type">
            <t-select class="form-input-md" v-model="searchForm.type" placeholder="交易方向筛选">
              <t-option value="" label="全部方向" />
              <t-option v-for="type in TRADING_TYPES" :key="type" :value="type" :label="type.text" />
            </t-select>
          </t-form-item>
          <t-form-item label="创建时间" name="dateRange">
            <t-date-range-picker
              v-model="searchForm.dateRange"
              placeholder="选择创建时间范围"
            />
          </t-form-item>
          <t-form-item name="actions" label-width="0">
            <t-space size="small">
              <t-button theme="primary" type="submit">
                <template #icon>
                  <t-icon name="search" />
                </template>
                搜索
              </t-button>
              <t-button theme="default" @click="resetFilter">
                <template #icon>
                  <t-icon name="refresh" />
                </template>
                重置
              </t-button>
            </t-space>
          </t-form-item>
        </t-form>
        
        <!-- 计划列表 -->
        <div class="plan-list">
          <t-table
            :data="filteredPlans"
            :columns="columns"
            :bordered="true"
            row-key="id"
            :pagination="pagination"
            @page-change="handlePageChange"
          >
            <template #status="{ row }">
              <t-tag
                :theme="getStatusTheme(row.status)"
                variant="light"
              >
                {{ getStatusText(row.status) }}
              </t-tag>
            </template>
            
            <template #type="{ row }">
              <t-tag
                :theme="row.type === 'buy' ? 'success' : 'warning'"
                variant="light"
              >
                {{ row.type === 'buy' ? '买多' : '买空' }}
              </t-tag>
            </template>
            
            <template #stock="{ row }">
              <div class="stock-info">
                <span class="stock-name">{{ row.stockName }}</span>
                <span class="stock-code">({{ row.stockCode }})</span>
              </div>
            </template>
            
            <template #targetPrice="{ row }">
              <span class="price-text">¥{{ row.targetPrice }}</span>
            </template>
            
            <template #stopLoss="{ row }">
              <span class="price-text">{{ row.stopLoss ? `¥${row.stopLoss}` : '-' }}</span>
            </template>
            
            <template #riskLevel="{ row }">
              <t-tag
                :color="getRiskLevelColor(row.riskLevel)"
                variant="light-outline"
                size="small"
              >
                {{ getRiskLevelText(row.riskLevel) }}
              </t-tag>
            </template>
            
            <template #operation="{ row }">
              <t-button
                size="small"
                theme="primary"
                variant="text"
                @click.stop="editPlan(row)"
              >
                编辑
              </t-button>
              <t-button
                size="small"
                theme="default"
                variant="text"
                @click.stop="viewPlan(row)"
              >
                查看
              </t-button>
              <t-dropdown
                :options="getActionOptions(row)"
                @click="handleAction"
              >
                <t-button size="small" variant="text">
                  更多
                  <template #suffix-icon>
                    <t-icon name="chevron-down" />
                  </template>
                </t-button>
              </t-dropdown>
            </template>
          </t-table>
        </div>
      </div>
    </t-card>
    
    <!-- 计划详情对话框 -->
    <t-dialog
      v-model:visible="showDetailDialog"
      :header="selectedPlan?.name || '计划详情'"
      width="600px"
      :confirm-btn="null"
      :cancel-btn="null"
    >
      <div v-if="selectedPlan" class="plan-detail">
        <div class="detail-grid">
          <div v-for="item in getPlanDetailData(selectedPlan)" :key="item.label" class="detail-item">
            <span class="detail-label">{{ item.label }}:</span>
            <span class="detail-content" v-if="!item.isRiskLevel">{{ item.content }}</span>
            <t-tag
              v-else-if="item.isRiskLevel && item.content !== '未设置'"
              :color="getRiskLevelColor(selectedPlan.riskLevel)"
              variant="light-outline"
              size="small"
            >
              {{ getRiskLevelText(selectedPlan.riskLevel) }}
            </t-tag>
            <span v-else class="detail-content">{{ item.content }}</span>
          </div>
        </div>
        
        <!-- 选股策略说明部分 -->
        <div v-if="selectedPlan.strategy || selectedPlan.tradingStrategy" class="strategy-section">
          <h4 class="strategy-title">选股策略说明</h4>
          <div class="strategy-content">
            <div v-if="selectedPlan.strategy" class="strategy-item">
              <h5>选股策略：{{ getStrategyName(selectedPlan.strategy) }}</h5>
              <p>{{ getStrategyDescription(selectedPlan.strategy) }}</p>
            </div>
            <div v-if="selectedPlan.tradingStrategy" class="strategy-item">
              <h5>交易策略：{{ getTradingStrategyName(selectedPlan.tradingStrategy) }}</h5>
              <p>{{ getTradingStrategyDescription(selectedPlan.tradingStrategy) }}</p>
            </div>
            <div v-if="selectedPlan.description" class="strategy-item">
              <h5>详细说明</h5>
              <p>{{ selectedPlan.description }}</p>
            </div>
          </div>
        </div>
        
        <div class="detail-actions">
          <t-space size="small">
            <t-button theme="primary" @click="editPlan(selectedPlan)">
              <template #icon>
                <t-icon name="edit" />
              </template>
              编辑计划
            </t-button>
            <t-button theme="default" @click="showDetailDialog = false">
              关闭
            </t-button>
          </t-space>
        </div>
      </div>
    </t-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTradingPlanStore } from './store'
import { MessagePlugin } from 'tdesign-vue-next'
import { PLAN_STATUS, getTradingPlanStatusText, getTradingPlanStatusTheme } from './config'
import { TRADING_TYPES, RISK_LEVELS } from '@/utils/constants'
import { filterData } from '@/utils/helpers'

const router = useRouter()
const tradingPlanStore = useTradingPlanStore()

// 搜索表单数据
const searchForm = reactive({
  keyword: '',
  status: '',
  type: '',
  dateRange: []
})

const showDetailDialog = ref(false)
const selectedPlan = ref(null)

const columns = [
  { colKey: 'name', title: '计划名称', width: 180 },
  { colKey: 'stock', title: '股票', width: 150 },
  { colKey: 'type', title: '交易方向', width: 80 },
  { colKey: 'targetPrice', title: '计划买进价格', width: 100 },
  { colKey: 'stopLoss', title: '止损价格', width: 100 },
  { colKey: 'quantity', title: '数量', width: 80 },
  { colKey: 'riskLevel', title: '风险等级', width: 100 },
  { colKey: 'status', title: '状态', width: 100 },
  { colKey: 'createdAt', title: '创建时间', width: 150 },
  { colKey: 'operation', title: '操作', width: 150 }
]

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

// 计算属性
const filteredPlans = computed(() => {
  return tradingPlanStore.getPlans || []
})

// 方法
const getStatusTheme = (status) => getTradingPlanStatusTheme(status) || 'default'
const getStatusText = (status) => getTradingPlanStatusText(status) || status


const getActionOptions = (row) => {
  const options = []
  
  if (row.status === 'active') {
    options.push(
      { content: '暂停计划', value: 'pause', row: row },
      { content: '完成计划', value: 'complete', row: row }
    )
  } else if (row.status === 'paused') {
    options.push(
      { content: '恢复计划', value: 'resume', row: row },
      { content: '取消计划', value: 'cancel', row: row }
    )
  }
  
  options.push({ content: '删除计划', value: 'delete', row: row })
  return options
}

const getPlanDetailData = (plan) => {
  return [
    { label: '计划名称', content: plan.name },
    { label: '股票代码', content: plan.stockCode },
    { label: '股票名称', content: plan.stockName },
    { label: '交易方向', content: getTradingTypeText(plan.type) || plan.type },
    { label: '计划买进价格', content: `¥${plan.targetPrice}` },
    { label: '计划数量', content: plan.quantity },
    { label: '止损价格', content: plan.stopLoss ? `¥${plan.stopLoss}` : '未设置' },
    { label: '止盈价格', content: plan.takeProfit ? `¥${plan.takeProfit}` : '未设置' },
    { label: '选股策略', content: plan.strategy || '未设置' },
    { label: '交易策略', content: plan.tradingStrategy || '未设置' },
    { label: '风险等级', content: plan.riskLevel || '未设置', isRiskLevel: true },
    { label: '策略说明', content: plan.description || '无' },
    { label: '当前状态', content: getStatusText(plan.status) },
    { label: '创建时间', content: plan.createdAt },
    { label: '更新时间', content: plan.updatedAt },
    { label: '备注', content: plan.remark || '无' }
  ]
}

const goToCreatePlan = () => {
  router.push('/trading-plan/create')
}

// 获取策略名称
const getStrategyName = (strategyId) => {
  const strategyNames = {
    'technical_analysis': '技术分析',
    'fundamental_analysis': '基本面分析',
    'momentum': '动量策略',
    'value': '价值投资',
    'growth': '成长投资',
    'bollinger_reversal': '布林反转',
    'mean_reversion': '均值回归',
    'breakout': '突破策略'
  }
  return strategyNames[strategyId] || strategyId
}

// 获取风险等级文本
const getRiskLevelText = (riskLevel) => {
  const texts = {
    'low': '低风险',
    'medium': '中风险',
    'high': '高风险'
  }
  return texts[riskLevel] || '未知'
}

// 获取风险等级颜色
const getRiskLevelColor = (riskLevel) => {
  const colors = {
    'low': '#52c41a',    // 绿色
    'medium': '#faad14', // 橙色
    'high': '#ff4d4f'    // 红色
  }
  return colors[riskLevel] || '#d9d9d9'
}

// 获取策略描述
const getStrategyDescription = (strategyId) => {
  const strategyDescriptions = {
    'technical_analysis': '通过分析股票的技术指标、图表形态和价格走势来选股，重点关注K线形态、均线系统、成交量等技术指标。',
    'fundamental_analysis': '基于公司的财务状况、行业前景、管理层能力等基本面因素进行选股，寻找具有长期投资价值的优质企业。',
    'momentum': '利用股票价格的动量效应，选择近期表现强势的股票，跟随市场趋势进行投资。',
    'value': '寻找被市场低估的优质股票，通过分析公司的内在价值与市场价格的差异来发现投资机会。',
    'growth': '重点关注具有高成长潜力的公司，通过分析公司的盈利增长、市场份额扩张等成长性指标来选股。',
    'bollinger_reversal': '利用布林带指标识别股价的超买超卖状态，在价格触及布林带边界时进行反向操作。',
    'mean_reversion': '基于均值回归理论，选择偏离均值的股票，预期价格会回归到均值水平。',
    'breakout': '寻找突破重要阻力位或支撑位的股票，跟随突破方向进行交易。'
  }
  return strategyDescriptions[strategyId] || '暂无详细描述'
}

// 获取交易策略名称
const getTradingStrategyName = (strategyId) => {
  const tradingStrategyNames = {
    'scalping': '剥头皮',
    'day_trading': '日内交易',
    'swing_trading': '波段交易',
    'position_trading': '持仓交易'
  }
  return tradingStrategyNames[strategyId] || strategyId
}

// 获取交易策略描述
const getTradingStrategyDescription = (strategyId) => {
  const tradingStrategyDescriptions = {
    'scalping': '利用极短时间内的价格波动进行频繁交易，追求小幅但稳定的利润，适合高流动性的市场环境。',
    'day_trading': '在同一个交易日内完成买卖操作，不持仓过夜，通过捕捉日内价格波动来获取利润。',
    'swing_trading': '持有股票数天到数周，利用中期的价格波动来获取利润，适合有一定趋势性的市场。',
    'position_trading': '长期持有股票，基于基本面分析进行投资决策，适合价值投资和成长投资策略。'
  }
  return tradingStrategyDescriptions[strategyId] || '暂无详细描述'
}


const resetFilter = () => {
  searchForm.keyword = ''
  searchForm.status = ''
  searchForm.type = ''
  searchForm.dateRange = []
  pagination.current = 1 // 重置到第一页
  loadPlans() // 重新加载数据
}

const handlePageChange = (pageInfo) => {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
  loadPlans()
}

const editPlan = (plan) => {
  router.push(`/trading-plan/edit/${plan.id}`)
}

const viewPlan = (plan) => {
  selectedPlan.value = plan
  showDetailDialog.value = true
}

const handleAction = async (data) => {
  const { value, row } = data
  switch (value) {
    case 'view':
      viewPlan(row)
      break
    case 'edit':
      editPlan(row)
      break
    case 'pause':
      tradingPlanStore.updatePlanStatus(row.id, 'paused')
      MessagePlugin.success('计划已暂停')
      break
    case 'resume':
      tradingPlanStore.updatePlanStatus(row.id, 'active')
      MessagePlugin.success('计划已恢复')
      break
    case 'complete':
      tradingPlanStore.updatePlanStatus(row.id, 'completed')
      MessagePlugin.success('计划已完成')
      break
    case 'cancel':
      tradingPlanStore.updatePlanStatus(row.id, 'cancelled')
      MessagePlugin.success('计划已取消')
      break
    case 'delete':
      try {
        await tradingPlanStore.deletePlan(row.id)
        MessagePlugin.success('计划已删除')
      } catch (error) {
        MessagePlugin.error('删除计划失败: ' + error.message)
      }
      break
  }
}


// 加载计划数据
const loadPlans = async () => {
  const params = {
    keyword: searchForm.keyword,
    status: searchForm.status,
    type: searchForm.type,
    startDate: searchForm.dateRange?.[0] || '',
    endDate: searchForm.dateRange?.[1] || '',
    page: pagination.current,
    pageSize: pagination.pageSize
  }
  
  try {
    const result = await tradingPlanStore.loadPlans(params)
    if (result && result.total !== undefined) {
      pagination.total = result.total
    }
  } catch (error) {
    console.error('加载交易计划失败:', error)
  }
}

// 搜索处理
const handleSearch = () => {
  pagination.current = 1 // 重置到第一页
  loadPlans()
}

onMounted(() => {
  loadPlans()
})
</script>

<style scoped>

.plan-card {
  border: none !important;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
  position: relative;
  z-index: 1;
}

.create-btn {
  margin-left: auto;
  position: relative;
  z-index: 10;
}

/* 股票信息样式 */
.stock-info {
  display: flex;
  align-items: center;
  gap: 4px;
}

.stock-name {
  font-weight: 500;
  color: var(--td-text-color-primary);
}

.stock-code {
  font-size: 12px;
  color: var(--td-text-color-secondary);
}

/* 价格文本样式 */
.price-text {
  font-weight: 500;
  color: var(--td-text-color-primary);
}

/* 详情弹窗样式 */
.detail-grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 8px 16px;
  margin-bottom: 16px;
}

.detail-item {
  display: flex;
  flex-direction: column;
  padding: 6px 8px;
  background: var(--td-bg-color-container-hover);
  border-radius: 4px;
}

.detail-label {
  font-weight: 500;
  color: var(--td-text-color-secondary);
  font-size: 12px;
  margin-bottom: 2px;
}

.detail-content {
  color: var(--td-text-color-primary);
  font-size: 14px;
  font-weight: 500;
}

/* 策略说明部分样式 */
.strategy-section {
  margin-top: 20px;
  padding: 16px;
  background: var(--td-bg-color-container);
  border-radius: 8px;
  border: 1px solid var(--td-border-level-1-color);
}

.strategy-title {
  margin: 0 0 12px 0;
  font-size: 16px;
  font-weight: 600;
  color: var(--td-text-color-primary);
  border-bottom: 2px solid var(--td-brand-color);
  padding-bottom: 8px;
}

.strategy-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.strategy-item {
  padding: 12px;
  background: var(--td-bg-color-container-hover);
  border-radius: 6px;
  border-left: 3px solid var(--td-brand-color);
}

.strategy-item h5 {
  margin: 0 0 8px 0;
  font-size: 14px;
  font-weight: 600;
  color: var(--td-text-color-primary);
}

.strategy-item p {
  margin: 0;
  font-size: 13px;
  line-height: 1.5;
  color: var(--td-text-color-secondary);
}
</style>
