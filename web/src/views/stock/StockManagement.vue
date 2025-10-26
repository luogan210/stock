<template>
  <div class="stock-management">
    <t-card class="stock-card">
      <template #header>
        <div class="card-header">
          <t-icon name="chart" size="20px" />
          <span>股票管理</span>
          <t-button theme="primary" @click="addStock" class="create-btn">
            <template #icon>
              <t-icon name="add" />
            </template>
            新增股票
          </t-button>
        </div>
      </template>
      <!-- 搜索栏 -->
      <div class="search-section">
        <t-form :model="searchForm" layout="inline" @submit="handleSearch">
          <t-form-item label="股票代码/名称" name="keyword">
            <t-input v-model="searchForm.keyword" placeholder="请输入股票代码或名称" clearable />
          </t-form-item>
          <t-form-item label="地区" name="region">
            <t-select v-model="searchForm.region" placeholder="请选择地区" clearable>
              <t-option value="china" label="中国" />
              <t-option value="hongkong" label="香港" />
              <t-option value="usa" label="美国" />
            </t-select>
          </t-form-item>
          <t-form-item label-width="0">
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


      <!-- 股票列表 -->
      <t-table
        :data="filteredStocks"
        :columns="columns"
        :pagination="pagination"
        :bordered="true"
        @page-change="handlePageChange"
        row-key="id"
      >
        <!-- 股票代码和名称 -->
        <template #code="{ row }">
          <div class="stock-info">
            <span class="stock-name">{{ row.name }}</span>
            <span class="stock-code">({{ row.code }})</span>
          </div>
        </template>

        <!-- 地区 -->
        <template #region="{ row }">
          <t-tag :theme="getRegionTheme(row.region)">
            {{ getRegionText(row.region) }}
          </t-tag>
        </template>

        <!-- 分类 -->
        <template #category="{ row }">
          <t-tag :theme="getCategoryTheme(row.category)">
            {{ getCategoryText(row.category) }}
          </t-tag>
        </template>


        <!-- 操作 -->
        <template #operation="{ row }">
          <t-button
            size="small"
            theme="primary"
            variant="text"
            @click.stop="viewStock(row)"
          >
            查看
          </t-button>
          <t-button
            size="small"
            theme="default"
            variant="text"
            @click.stop="editStock(row)"
          >
            编辑
          </t-button>
          <t-button
            size="small"
            theme="danger"
            variant="text"
            @click.stop="deleteStock(row)"
          >
            删除
          </t-button>
        </template>
      </t-table>
    </t-card>

    <!-- 股票详情弹窗 -->
    <t-dialog
      v-model:visible="showDetailDialog"
      :header="selectedStock?.name || '股票详情'"
      width="600px"
      :confirm-btn="null"
      :cancel-btn="null"
    >
      <div v-if="selectedStock" class="stock-detail">
        <div class="detail-grid">
          <div v-for="item in getStockDetailData(selectedStock)" :key="item.label" class="detail-item">
            <span class="detail-label">{{ item.label }}:</span>
            <span class="detail-content">{{ item.content }}</span>
          </div>
        </div>
      </div>
    </t-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { MessagePlugin, DialogPlugin } from 'tdesign-vue-next'
import { useStockStore } from "./store/index.js";

const router = useRouter()
const stockStore = useStockStore()

// 搜索表单
const searchForm = reactive({
  keyword: '',
  region: ''
})

// 分页
const pagination = reactive({
  current: 1,
  pageSize: 10,
  total: 0
})

// 弹窗状态
const showDetailDialog = ref(false)
const selectedStock = ref(null)

// 表格列定义
const columns = [
  { colKey: 'code', title: '股票信息', width: 150 },
  { colKey: 'region', title: '地区', width: 80 },
  { colKey: 'category', title: '分类', width: 100 },
  { colKey: 'enabled', title: '状态', width: 80 },
  { colKey: 'operation', title: '操作', width: 200, fixed: 'right' }
]

// 计算属性
const filteredStocks = computed(() => {
  let stocks = (stockStore.getStocks || []).filter(stock => stock.enabled)

  if (searchForm.keyword) {
    const keyword = searchForm.keyword.toLowerCase()
    stocks = stocks.filter(stock =>
      stock.code.toLowerCase().includes(keyword) ||
      stock.name.toLowerCase().includes(keyword)
    )
  }

  if (searchForm.region) {
    stocks = stocks.filter(stock => stock.region === searchForm.region)
  }

  pagination.total = stocks.length
  const start = (pagination.current - 1) * pagination.pageSize
  const end = start + pagination.pageSize
  return stocks.slice(start, end)
})

// 工具函数
const getRegionText = (region) => {
  const texts = {
    'china': '中国',
    'hongkong': '香港',
    'usa': '美国'
  }
  return texts[region] || '未知'
}

const getCategoryText = (category) => {
  const texts = {
    'main_board': '主板',
    'hk_main': '港股主板',
    'us_nasdaq': '纳斯达克'
  }
  return texts[category] || '未知'
}

const getRiskLevelText = (riskLevel) => {
  const texts = {
    'low': '低风险',
    'medium': '中风险',
    'high': '高风险'
  }
  return texts[riskLevel] || '未知'
}

const getRegionTheme = (region) => {
  const themes = {
    'china': 'primary',
    'hongkong': 'success',
    'usa': 'warning'
  }
  return themes[region] || 'default'
}

const getCategoryTheme = (category) => {
  const themes = {
    'main_board': 'primary',
    'hk_main': 'success',
    'us_nasdaq': 'warning'
  }
  return themes[category] || 'default'
}


const getStockDetailData = (stock) => {
  return [
    { label: '股票代码', content: stock.code },
    { label: '股票名称', content: stock.name },
    { label: '地区', content: getRegionText(stock.region) },
    { label: '股票分类', content: getCategoryText(stock.category) },
    { label: '状态', content: stock.enabled ? '启用' : '禁用' },
    { label: '备注', content: stock.remark || '无' }
  ]
}

// 事件处理
const handleSearch = () => {
  pagination.current = 1
  loadStocks()
}

const resetSearch = () => {
  searchForm.keyword = ''
  searchForm.region = ''
  pagination.current = 1
  loadStocks()
}

const handlePageChange = (pageInfo) => {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
}

const viewStock = (stock) => {
  selectedStock.value = stock
  showDetailDialog.value = true
}

const editStock = (stock) => {
  router.push(`/stock/edit/${stock.id}`)
}

const deleteStock = async (stock) => {
  const dialogInstance = DialogPlugin.confirm({
    header: '确认删除',
    body: `确定要删除股票"${stock.name}"吗？`,
    onConfirm: async () => {
      try {
        await stockStore.removeStock(stock.id)
        MessagePlugin.success('删除成功')
        // 确保弹窗关闭
        dialogInstance.close()
      } catch (error) {
        MessagePlugin.error('删除失败: ' + error.message)
        // 即使删除失败也要关闭确认弹窗
        dialogInstance.close()
      }
    },
    onCancel: () => {
      // 取消时关闭弹窗
      dialogInstance.close()
    }
  })
}

const addStock = () => {
  router.push('/stock/create')
}

// 加载股票数据
const loadStocks = async () => {
  const params = {
    keyword: searchForm.keyword,
    region: searchForm.region
  }
  await stockStore.loadStocks(params)
}

// 组件挂载时加载数据
onMounted(() => {
  loadStocks()
})
</script>

<style scoped>


.stock-card {
  border: none !important;
  margin-bottom: 20px;
}

.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-weight: 600;
}

.create-btn {
  margin-left: auto;
}

.search-section {
  margin-bottom: 20px;
  /* padding: 16px; */
  /* background: #f8f9fa; */
  border-radius: 6px;
}

.action-section {
  margin-bottom: 20px;
  display: flex;
  gap: 12px;
}

.stock-code-name {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stock-code {
  font-size: 12px;
  color: var(--td-text-color-secondary);
}

.stock-name {
  font-size: 12px;
  color: var(--td-text-color-secondary);
}

.market-cap {
  font-weight: 500;
  color: var(--td-success-color);
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

</style>
