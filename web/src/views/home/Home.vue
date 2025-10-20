<template>
  <div class="home">
    <t-card title="欢迎使用股票分析系统" class="welcome-card">
      <template #header>
        <div class="card-header">
          <t-icon name="chart" size="24px" />
          <span>系统概览</span>
        </div>
      </template>
      
      <div class="welcome-content">
        <t-row :gutter="16">
          <t-col :span="3">
            <t-card class="stat-card">
              <div class="stat-item">
                <div class="stat-value">{{ stockCount }}</div>
                <div class="stat-label">股票数量</div>
              </div>
            </t-card>
          </t-col>
          <t-col :span="3">
            <t-card class="stat-card">
              <div class="stat-item">
                <div class="stat-value">{{ watchCount }}</div>
                <div class="stat-label">关注股票</div>
              </div>
            </t-card>
          </t-col>
          <t-col :span="3">
            <t-card class="stat-card">
              <div class="stat-item">
                <div class="stat-value">{{ planCount }}</div>
                <div class="stat-label">交易计划</div>
              </div>
            </t-card>
          </t-col>
          <t-col :span="3">
            <t-card class="stat-card">
              <div class="stat-item">
                <div class="stat-value">{{ logCount }}</div>
                <div class="stat-label">交易日志</div>
              </div>
            </t-card>
          </t-col>
          <t-col :span="3">
            <t-card class="stat-card">
              <div class="stat-item">
                <div class="stat-value" :class="totalProfit >= 0 ? 'profit-positive' : 'profit-negative'">
                  {{ totalProfit >= 0 ? '+' : '' }}{{ totalProfit.toFixed(2) }}
                </div>
                <div class="stat-label">总盈亏</div>
              </div>
            </t-card>
          </t-col>
          <t-col :span="3">
            <t-card class="stat-card">
              <div class="stat-item">
                <div class="stat-value">{{ todayTradingCount }}</div>
                <div class="stat-label">当日交易</div>
              </div>
            </t-card>
          </t-col>
          <t-col :span="3">
            <t-card class="stat-card">
              <div class="stat-item">
                <div class="stat-value">{{ weekTradingCount }}</div>
                <div class="stat-label">当周交易</div>
              </div>
            </t-card>
          </t-col>
          <t-col :span="3">
            <t-card class="stat-card">
              <div class="stat-item">
                <div class="stat-value">{{ monthTradingCount }}</div>
                <div class="stat-label">当月交易</div>
              </div>
            </t-card>
          </t-col>
        </t-row>
        
        <div class="quick-actions">
          <h3>快速操作</h3>
          <t-space size="small">
            <t-button theme="primary" @click="goToTradingPlan">
              <template #icon>
                <t-icon name="money" />
              </template>
              交易计划
            </t-button>
            <t-button theme="primary" @click="goToTradingLog">
              <template #icon>
                <t-icon name="file" />
              </template>
              交易日志
            </t-button>
          </t-space>
        </div>
      </div>
    </t-card>
  </div>
</template>

<script setup>
import { computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { useStockStore } from '@/stores/stock'
import { useHomeStore } from './store'
import { getDateRange } from '@/utils/helpers'
import { currentApi } from '@/services/api'

const router = useRouter()
const userStore = useUserStore()
const stockStore = useStockStore()
const homeStore = useHomeStore()

const stockCount = computed(() => stockStore.getStockList.length)
const watchCount = computed(() => stockStore.getWatchList.length)
const planCount = computed(() => homeStore.getHomeStats.planCount)
const activePlanCount = computed(() => homeStore.getHomeStats.activePlanCount || 0)
const logCount = computed(() => homeStore.getHomeStats.logCount)
const totalProfit = computed(() => homeStore.getHomeStats.totalProfit)

// 计算当日交易次数
const todayTradingCount = computed(() => homeStore.getHomeStats.todayTradingCount)

// 计算当周交易次数
const weekTradingCount = computed(() => homeStore.getHomeStats.weekTradingCount)

// 计算当月交易次数
const monthTradingCount = computed(() => homeStore.getHomeStats.monthTradingCount)

// 加载首页统计数据
const loadHomeStats = async () => {
  try {
    await homeStore.loadHomeStats()
  } catch (error) {
    console.error('加载首页统计数据失败:', error)
  }
}

// 组件挂载时加载数据
onMounted(() => {
  loadHomeStats()
})

const goToTradingPlan = () => {
  router.push('/trading-plan')
}

const goToTradingLog = () => {
  router.push('/trading-log')
}
</script>

<style scoped>
.home {
}

.welcome-card {
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
}

.welcome-content {
  padding: 20px 0;
}

.stat-card {
  text-align: center;
  border: 1px solid var(--td-border-level-1-color);
}

.stat-item {
  padding: 20px;
}

.stat-value {
  font-size: 32px;
  font-weight: bold;
  color: var(--td-brand-color);
  margin-bottom: 8px;
}

.stat-label {
  color: var(--td-text-color-secondary);
  font-size: 14px;
}

.quick-actions {
  margin-top: 30px;
}

.quick-actions h3 {
  margin-bottom: 16px;
  color: var(--td-text-color-primary);
}

.profit-positive {
  color: var(--td-success-color);
}

.profit-negative {
  color: var(--td-error-color);
}
</style>
