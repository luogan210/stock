<template>
  <div class="trading-review">
    <t-card title="交易复盘" :bordered="false">
      <template #header>
        <div class="card-header">
          <span>交易复盘管理</span>
          <t-button theme="primary" @click="goToCreate">
            <template #icon>
              <t-icon name="add" />
            </template>
            新建复盘
          </t-button>
        </div>
      </template>

      <!-- 搜索筛选 -->
      <div class="search-section">
        <t-form
          ref="searchFormRef"
          :model="searchForm"
          layout="inline"
          @submit="handleSearch"
        >
          <t-form-item label="关键词" name="keyword">
            <t-input
              class="form-input-md"
              v-model="searchForm.keyword"
              placeholder="搜索复盘标题"
              clearable
            >
              <template #prefix-icon>
                <t-icon name="search" />
              </template>
            </t-input>
          </t-form-item>
          <t-form-item label="周期" name="period">
            <t-select class="form-input-md" v-model="searchForm.period" placeholder="周期筛选">
              <t-option value="" label="全部周期" />
              <t-option value="daily" label="日复盘" />
              <t-option value="weekly" label="周复盘" />
              <t-option value="monthly" label="月复盘" />
            </t-select>
          </t-form-item>
          <t-form-item label="状态" name="status">
            <t-select class="form-input-md" v-model="searchForm.status" placeholder="状态筛选">
              <t-option value="" label="全部状态" />
              <t-option value="completed" label="已完成" />
              <t-option value="draft" label="草稿" />
            </t-select>
          </t-form-item>
          <t-form-item label="时间范围" name="dateRange">
            <t-date-range-picker
              v-model="searchForm.dateRange"
              placeholder="选择时间范围"
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
              <t-button theme="default" @click="resetSearch">
                <template #icon>
                  <t-icon name="refresh" />
                </template>
                重置
              </t-button>
            </t-space>
          </t-form-item>
        </t-form>
      </div>

      <!-- 复盘列表 -->
      <t-table
        :data="filteredReviews"
        :columns="columns"
        :loading="isLoading"
        :pagination="pagination"
        :bordered="true"
        @page-change="handlePageChange"
        row-key="id"
      >
        <!-- 复盘周期 -->
        <template #period="{ row }">
          <t-tag :theme="getPeriodTheme(row.period)">
            {{ getPeriodText(row.period) }}
          </t-tag>
        </template>

        <!-- 日期范围 -->
        <template #dateRange="{ row }">
          <span class="date-range">{{ row.dateRange }}</span>
        </template>

        <!-- 盈亏情况 -->
        <template #profit="{ row }">
          <div class="profit-info">
            <div class="profit-amount" :class="getProfitClass(row.totalProfit)">
              {{ formatProfit(row.totalProfit).text }}
            </div>
            <div class="profit-rate">
              胜率: {{ row.winRate }}%
            </div>
          </div>
        </template>

        <!-- 交易统计 -->
        <template #trades="{ row }">
          <div class="trades-info">
            <div>买入: {{ row.buyCount }}次</div>
            <div>卖出: {{ row.sellCount }}次</div>
          </div>
        </template>

        <!-- 操作 -->
        <template #operation="{ row }">
          <t-button
            size="small"
            theme="primary"
            variant="text"
            @click="viewReview(row)"
          >
            查看
          </t-button>
          <t-button
            size="small"
            theme="default"
            variant="text"
            @click="editReview(row)"
          >
            编辑
          </t-button>
          <t-button
            size="small"
            theme="danger"
            variant="text"
            @click="deleteReview(row)"
          >
            删除
          </t-button>
        </template>
      </t-table>
    </t-card>

    <!-- 复盘详情弹窗 -->
    <t-dialog
      v-model:visible="showDetailDialog"
      :header="selectedReview?.title"
      width="800px"
      :footer="false"
    >
      <div v-if="selectedReview" class="review-detail">
        <t-descriptions :data="getReviewDetailData(selectedReview)" />
        
        <div class="review-content">
          <h4>复盘内容</h4>
          <p>{{ selectedReview.content }}</p>
        </div>
        
        <div class="review-lessons">
          <h4>经验教训</h4>
          <ul>
            <li v-for="(lesson, index) in selectedReview.lessons" :key="index">
              {{ lesson }}
            </li>
          </ul>
        </div>
        
        <div class="review-improvements">
          <h4>改进计划</h4>
          <ul>
            <li v-for="(improvement, index) in selectedReview.improvements" :key="index">
              {{ improvement }}
            </li>
          </ul>
        </div>
        
        <div class="review-plan">
          <h4>{{ getNextPlanLabel(selectedReview.period) }}</h4>
          <p>{{ selectedReview.nextPlan }}</p>
        </div>
      </div>
    </t-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useTradingReviewStore } from './store'
import { MessagePlugin, DialogPlugin } from 'tdesign-vue-next'
import { filterData, formatProfit } from '@/utils/helpers'

const router = useRouter()
const tradingReviewStore = useTradingReviewStore()

// 响应式数据
const searchFormRef = ref()
const showDetailDialog = ref(false)
const selectedReview = ref(null)

// 搜索表单
const searchForm = reactive({
  keyword: '',
  period: '',
  status: '',
  dateRange: null
})

// 分页
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

// 表格列配置
const columns = [
  { colKey: 'title', title: '复盘标题' },
  { colKey: 'period', title: '复盘周期' },
  { colKey: 'dateRange', title: '日期范围' },
  { colKey: 'trades', title: '交易统计' },
  { colKey: 'profit', title: '盈亏情况' },
  { colKey: 'createTime', title: '创建时间' },
  { colKey: 'operation', title: '操作' }
]

// 计算属性
const reviews = computed(() => tradingReviewStore.getReviews)
const isLoading = computed(() => tradingReviewStore.isLoading)

// 筛选后的复盘列表
const filteredReviews = computed(() => {
  return filterData(reviews.value, searchForm)
})

// 获取周期主题色
const getPeriodTheme = (period) => {
  const themes = {
    daily: 'success',
    weekly: 'primary',
    monthly: 'warning'
  }
  return themes[period] || 'default'
}

// 获取周期文本
const getPeriodText = (period) => {
  const texts = {
    daily: '日复盘',
    weekly: '周复盘',
    monthly: '月复盘'
  }
  return texts[period] || period
}

// 获取下期计划标签
const getNextPlanLabel = (period) => {
  const labels = {
    daily: '明日计划',
    weekly: '下周计划',
    monthly: '下月计划'
  }
  return labels[period] || '下期计划'
}

// 获取盈亏样式
const getProfitClass = (profit) => {
  return profit >= 0 ? 'profit-positive' : 'profit-negative'
}

// 获取复盘详情数据
const getReviewDetailData = (review) => {
  return [
    { label: '复盘标题', content: review.title },
    { label: '复盘周期', content: getPeriodText(review.period) },
    { label: '日期范围', content: review.dateRange },
    { label: '总交易次数', content: review.totalTrades },
    { label: '买入次数', content: review.buyCount },
    { label: '卖出次数', content: review.sellCount },
    { label: '盈利交易', content: review.winTrades },
    { label: '亏损交易', content: review.lossTrades },
    { label: '胜率', content: `${review.winRate}%` },
    { label: '总盈亏', content: formatProfit(review.totalProfit).text },
    { label: '平均盈利', content: formatProfit(review.avgProfit).text },
    { label: '平均亏损', content: formatProfit(review.avgLoss).text },
    { label: '最大亏损', content: formatProfit(review.maxLoss).text },
    { label: '创建时间', content: review.createTime },
    { label: '更新时间', content: review.updateTime }
  ]
}

// 搜索
const handleSearch = () => {
  pagination.current = 1
  loadReviews()
}

// 重置搜索
const resetSearch = () => {
  Object.assign(searchForm, {
    keyword: '',
    period: '',
    status: '',
    dateRange: null
  })
  pagination.current = 1
  loadReviews()
}

// 分页变化
const handlePageChange = (pageInfo) => {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
  loadReviews()
}

// 加载复盘列表
const loadReviews = async () => {
  try {
    const params = {
      page: pagination.current,
      pageSize: pagination.pageSize,
      ...searchForm
    }
    
    const result = await tradingReviewStore.loadReviews(params)
    pagination.total = result.total
  } catch (error) {
    MessagePlugin.error('加载复盘列表失败')
  }
}

// 查看复盘
const viewReview = (review) => {
  selectedReview.value = review
  showDetailDialog.value = true
}

// 编辑复盘
const editReview = (review) => {
  router.push(`/trading-review/edit/${review.id}`)
}

// 删除复盘
const deleteReview = (review) => {
  DialogPlugin.confirm({
    header: '确认删除',
    body: `确定要删除复盘"${review.title}"吗？`,
    onConfirm: async () => {
      try {
        await tradingReviewStore.deleteReview(review.id)
        MessagePlugin.success('删除成功')
        loadReviews()
      } catch (error) {
        MessagePlugin.error('删除失败')
      }
    }
  })
}

// 跳转到创建页面
const goToCreate = () => {
  router.push('/trading-review/create')
}

// 组件挂载时加载数据
onMounted(() => {
  loadReviews()
})
</script>

<style scoped>
.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-weight: 600;
}

.search-section {
  margin-bottom: 20px;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 6px;
}

.profit-info {
  text-align: center;
}

.profit-amount {
  font-size: 16px;
  font-weight: 600;
  margin-bottom: 4px;
}

.profit-positive {
  color: #00a870;
}

.profit-negative {
  color: #d54941;
}

.profit-rate {
  font-size: 12px;
  color: #666;
}

.trades-info {
  font-size: 12px;
  line-height: 1.4;
}

.review-detail {
  max-height: 600px;
  overflow-y: auto;
}

.review-content,
.review-lessons,
.review-improvements,
.review-plan {
  margin-top: 20px;
}

.review-content h4,
.review-lessons h4,
.review-improvements h4,
.review-plan h4 {
  margin-bottom: 8px;
  color: #333;
  font-size: 14px;
  font-weight: 600;
}

.review-lessons ul,
.review-improvements ul {
  margin: 0;
  padding-left: 20px;
}

.review-lessons li,
.review-improvements li {
  margin-bottom: 4px;
  line-height: 1.5;
}

.date-range {
  font-size: 12px;
  color: #666;
}
</style>
