<template>
  <div class="trading-log">
    <t-card title="交易日志" class="log-card">
      <template #header>
        <div class="card-header">
          <t-icon name="file" size="24px" />
          <span>我的交易日志</span>
          <t-button theme="primary" @click="goToCreateLog" class="create-btn">
            <template #icon>
              <t-icon name="add" />
            </template>
            新建日志
          </t-button>
        </div>
      </template>
      
      <div class="log-content">
        <!-- 筛选和搜索 -->
        <t-form :model="searchForm" @submit="handleSearch" layout="inline" label-width="80px" class="search-form">
          <t-form-item label="关键词" name="keyword">
            <t-input
              class="form-input-md"
              v-model="searchForm.keyword"
              placeholder="搜索股票代码、名称或日志内容"
              clearable
            >
              <template #prefix-icon>
                <t-icon name="search" />
              </template>
            </t-input>
          </t-form-item>
          <t-form-item label="方向" name="type">
            <t-select class="form-input-md" v-model="searchForm.type" placeholder="交易方向筛选">
              <t-option value="" label="全部方向" />
              <t-option value="buy" label="买多" />
              <t-option value="sell" label="买空" />
            </t-select>
          </t-form-item>
          <t-form-item label="状态" name="status">
            <t-select class="form-input-md" v-model="searchForm.status" placeholder="状态筛选">
              <t-option value="" label="全部状态" />
              <t-option value="pending" label="进行中" />
              <t-option value="completed" label="已经结束" />
            </t-select>
          </t-form-item>
          <t-form-item label="更新时间" name="dateRange">
            <t-date-range-picker
              v-model="searchForm.dateRange"
              placeholder="选择更新时间范围"
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
        
        <!-- 日志列表 -->
        <div class="log-list">
          <t-table
          :bordered="true"
            :data="filteredLogs"
            :columns="columns"
            row-key="id"
            :pagination="pagination"
            @page-change="handlePageChange"
          >
            <template #stock="{ row }">
              <div class="stock-info">
                <span class="stock-name">{{ row.stockName }}</span>
                <span class="stock-code">({{ row.stockCode }})</span>
              </div>
            </template>
            
            <template #type="{ row }">
              <t-tag
                :theme="row.type === 'buy' ? 'success' : 'warning'"
                variant="light"
              >
                {{ row.type === 'buy' ? '买多' : '买空' }}
              </t-tag>
            </template>
            
            <template #status="{ row }">
              <t-tag
                :theme="getStatusTheme(row.status)"
                variant="light"
              >
                {{ getStatusText(row.status) }}
              </t-tag>
            </template>
            
            <template #profit="{ row }">
              <span :class="getProfitClass(row.profit)">
                {{ formatProfit(row.profit).text }}
              </span>
            </template>
            
            <template #operation="{ row }">
              <t-button
                size="small"
                theme="primary"
                variant="text"
                @click.stop="editLog(row)"
              >
                编辑
              </t-button>
              <t-button
                size="small"
                theme="default"
                variant="text"
                @click.stop="viewLog(row)"
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
    
    <!-- 日志详情对话框 -->
    <t-dialog
      v-model:visible="showDetailDialog"
      :header="selectedLog?.title || '日志详情'"
      width="800px"
      :confirm-btn="null"
      :cancel-btn="null"
    >
      <div v-if="selectedLog" class="log-detail">
        <t-descriptions :data="getLogDetailData(selectedLog)" :column="2" />
        
        <div class="detail-actions">
          <t-space size="small">
            <t-button theme="primary" @click="editLog(selectedLog)">
              <template #icon>
                <t-icon name="edit" />
              </template>
              编辑日志
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
import { useTradingLogStore } from './store'
import { MessagePlugin } from 'tdesign-vue-next'
import { LOG_STATUS_TEXT, LOG_STATUS_THEME, TRADING_TYPE_TEXT } from '@/utils/constants'
import { filterData, formatProfit } from '@/utils/helpers'

const router = useRouter()
const tradingLogStore = useTradingLogStore()

// 搜索表单数据
const searchForm = reactive({
  keyword: '',
  type: '',
  status: '',
  dateRange: []
})

const showDetailDialog = ref(false)
const selectedLog = ref(null)

const columns = [
  { colKey: 'title', title: '日志标题', width: 180 },
  { colKey: 'stock', title: '股票', width: 150 },
  { colKey: 'type', title: '交易方向', width: 80 },
  { colKey: 'price', title: '成交价格', width: 100 },
  { colKey: 'quantity', title: '成交数量', width: 80 },
  { colKey: 'profit', title: '盈亏', width: 100 },
  { colKey: 'status', title: '状态', width: 80 },
  { colKey: 'planName', title: '执行交易计划', width: 120 },
  { colKey: 'tradingTime', title: '交易时间', width: 150 },
  { colKey: 'operation', title: '操作', width: 150 }
]

const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

// 计算属性
const filteredLogs = computed(() => {
  return tradingLogStore.getLogs || []
})

// 方法
const getStatusTheme = (status) => LOG_STATUS_THEME[status] || 'default'
const getStatusText = (status) => LOG_STATUS_TEXT[status] || status
const getProfitClass = (profit) => formatProfit(profit).class

const getActionOptions = (row) => {
  const options = []
  
  options.push({ 
    content: '删除日志', 
    value: 'delete',
    row: row  // 将row数据包含在选项中
  })
  return options
}

const getLogDetailData = (log) => [
  { label: '日志标题', content: log.title },
  { label: '股票代码', content: log.stockCode },
  { label: '股票名称', content: log.stockName },
  { label: '交易方向', content: TRADING_TYPE_TEXT[log.type] || log.type },
  { label: '成交价格', content: `¥${log.price}` },
  { label: '成交数量', content: log.quantity },
  { label: '盈亏金额', content: formatProfit(log.profit).text },
  { label: '交易状态', content: getStatusText(log.status) },
  { label: '执行交易计划', content: log.planName || '无关联计划' },
  { label: '交易时间', content: log.tradingTime },
  { label: '创建时间', content: log.createTime },
  { label: '更新时间', content: log.updateTime },
  { label: '日志内容', content: log.content || '无' },
  { label: '备注', content: log.remark || '无' }
]

const goToCreateLog = () => {
  router.push('/trading-log/create')
}


const resetFilter = () => {
  searchForm.keyword = ''
  searchForm.type = ''
  searchForm.status = ''
  searchForm.dateRange = []
  pagination.current = 1 // 重置到第一页
  loadLogs() // 重新加载数据
}

const handlePageChange = (pageInfo) => {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
  loadLogs()
}

const editLog = (log) => {
  router.push(`/trading-log/edit/${log.id}`)
}

const viewLog = (log) => {
  selectedLog.value = log
  showDetailDialog.value = true
}

const handleAction = async (data) => {
  const { value, row } = data
  switch (value) {
    case 'view':
      viewLog(row)
      break
    case 'edit':
      editLog(row)
      break
    case 'success':
      tradingLogStore.updateLogStatus(row.id, 'success')
      MessagePlugin.success('日志状态已更新为成功')
      break
    case 'failed':
      tradingLogStore.updateLogStatus(row.id, 'failed')
      MessagePlugin.success('日志状态已更新为失败')
      break
    case 'delete':
      try {
        await tradingLogStore.deleteLog(row.id)
        MessagePlugin.success('日志已删除')
      } catch (error) {
        MessagePlugin.error('删除日志失败: ' + error.message)
      }
      break
  }
}

const exportLogs = () => {
  const dataStr = JSON.stringify(filteredLogs.value, null, 2)
  const dataBlob = new Blob([dataStr], { type: 'application/json' })
  const url = URL.createObjectURL(dataBlob)
  const link = document.createElement('a')
  link.href = url
  link.download = 'trading-logs.json'
  link.click()
  URL.revokeObjectURL(url)
  MessagePlugin.success('日志已导出')
}

// 加载日志数据
const loadLogs = async () => {
  const params = {
    keyword: searchForm.keyword,
    type: searchForm.type,
    status: searchForm.status,
    startDate: searchForm.dateRange?.[0] || '',
    endDate: searchForm.dateRange?.[1] || '',
    page: pagination.current,
    pageSize: pagination.pageSize
  }
  
  try {
    const result = await tradingLogStore.loadLogs(params)
    if (result && result.total !== undefined) {
      pagination.total = result.total
    }
  } catch (error) {
    console.error('加载交易日志失败:', error)
  }
}

// 搜索处理
const handleSearch = () => {
  pagination.current = 1 // 重置到第一页
  loadLogs()
}

onMounted(() => {
  loadLogs()
})
</script>

<style scoped>

.log-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
}

.create-btn {
  margin-left: auto;
}


.profit-positive {
  color: var(--td-success-color);
  font-weight: bold;
}

.profit-negative {
  color: var(--td-error-color);
  font-weight: bold;
}

.profit-neutral {
  color: var(--td-text-color-secondary);
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
</style>
