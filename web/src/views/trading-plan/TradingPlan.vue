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
              <t-option value="active" label="进行中" />
              <t-option value="completed" label="已完成" />
              <t-option value="paused" label="已暂停" />
              <t-option value="cancelled" label="已取消" />
            </t-select>
          </t-form-item>
          <t-form-item label="方向" name="type">
            <t-select class="form-input-md" v-model="searchForm.type" placeholder="交易方向筛选">
              <t-option value="" label="全部方向" />
              <t-option value="buy" label="买多" />
              <t-option value="sell" label="买空" />
            </t-select>
          </t-form-item>
          <t-form-item label="创建时间" name="dateRange">
            <t-date-range-picker
              v-model="searchForm.dateRange"
              placeholder="选择创建时间范围"
            />
          </t-form-item>
          <t-form-item label="操作" name="actions">
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
      width="800px"
      :confirm-btn="null"
      :cancel-btn="null"
    >
      <div v-if="selectedPlan" class="plan-detail">
        <t-descriptions :data="getPlanDetailData(selectedPlan)" :column="2" />
        
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
import { PLAN_STATUS_TEXT, PLAN_STATUS_THEME, TRADING_TYPE_TEXT } from '@/utils/constants'
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
  { colKey: 'status', title: '状态', width: 100 },
  { colKey: 'createTime', title: '创建时间', width: 150 },
  { colKey: 'operation', title: '操作', width: 150 }
]

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

// 计算属性
const filteredPlans = computed(() => {
  const plans = filterData(tradingPlanStore.getPlans, searchForm)
  pagination.total = plans.length
  return plans
})

// 方法
const getStatusTheme = (status) => PLAN_STATUS_THEME[status] || 'default'
const getStatusText = (status) => PLAN_STATUS_TEXT[status] || status


const getActionOptions = (row) => {
  const options = []
  
  if (row.status === 'active') {
    options.push(
      { content: '暂停计划', value: 'pause' },
      { content: '完成计划', value: 'complete' }
    )
  } else if (row.status === 'paused') {
    options.push(
      { content: '恢复计划', value: 'resume' },
      { content: '取消计划', value: 'cancel' }
    )
  }
  
  options.push({ content: '删除计划', value: 'delete' })
  return options
}

const getPlanDetailData = (plan) => [
  { label: '计划名称', content: plan.name },
  { label: '股票代码', content: plan.stockCode },
  { label: '股票名称', content: plan.stockName },
  { label: '交易方向', content: TRADING_TYPE_TEXT[plan.type] || plan.type },
  { label: '计划买进价格', content: `¥${plan.targetPrice}` },
  { label: '计划数量', content: plan.quantity },
  { label: '当前状态', content: getStatusText(plan.status) },
  { label: '创建时间', content: plan.createTime },
  { label: '更新时间', content: plan.updateTime },
  { label: '备注', content: plan.remark || '无' }
]

const goToCreatePlan = () => {
  router.push('/trading-plan/create')
}

const handleSearch = () => {
  // 搜索逻辑已在计算属性中处理
}

const resetFilter = () => {
  searchForm.keyword = ''
  searchForm.status = ''
  searchForm.type = ''
  searchForm.dateRange = []
}

const handlePageChange = (pageInfo) => {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
}

const editPlan = (plan) => {
  router.push(`/trading-plan/edit/${plan.id}`)
}

const viewPlan = (plan) => {
  selectedPlan.value = plan
  showDetailDialog.value = true
}

const handleAction = (data) => {
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
      tradingPlanStore.deletePlan(row.id)
      MessagePlugin.success('计划已删除')
      break
  }
}


onMounted(() => {
  tradingPlanStore.loadPlans()
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
</style>
