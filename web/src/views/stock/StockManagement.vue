<template>
  <div class="stock-management">
    <t-card title="股票管理" class="stock-card">
      <!-- 搜索栏 -->
      <div class="search-section">
        <t-form :model="searchForm" layout="inline" @submit="handleSearch">
          <t-form-item label="股票代码/名称" name="keyword">
            <t-input v-model="searchForm.keyword" placeholder="请输入股票代码或名称" clearable />
          </t-form-item>
          <t-form-item label="市场类型" name="marketType">
            <t-select v-model="searchForm.marketType" placeholder="请选择市场" clearable>
              <t-option value="a_share" label="A股" />
              <t-option value="hk_stock" label="港股" />
              <t-option value="us_stock" label="美股" />
              <t-option value="other" label="其他" />
            </t-select>
          </t-form-item>
          <t-form-item label="股票分类" name="category">
            <t-select v-model="searchForm.category" placeholder="请选择分类" clearable>
              <t-option value="main_board" label="主板" />
              <t-option value="sme_board" label="中小板" />
              <t-option value="gem_board" label="创业板" />
              <t-option value="star_board" label="科创板" />
              <t-option value="new_third_board" label="新三板" />
              <t-option value="hk_main" label="港股主板" />
              <t-option value="hk_gem" label="港股创业板" />
              <t-option value="us_nyse" label="纽交所" />
              <t-option value="us_nasdaq" label="纳斯达克" />
              <t-option value="us_amex" label="美交所" />
            </t-select>
          </t-form-item>
          <t-form-item label="行业分类" name="industry">
            <t-select v-model="searchForm.industry" placeholder="请选择行业" clearable>
              <t-option value="technology" label="科技" />
              <t-option value="finance" label="金融" />
              <t-option value="healthcare" label="医疗" />
              <t-option value="consumer" label="消费" />
              <t-option value="industry" label="工业" />
              <t-option value="energy" label="能源" />
              <t-option value="materials" label="材料" />
              <t-option value="utilities" label="公用事业" />
              <t-option value="real_estate" label="房地产" />
              <t-option value="communication" label="通信" />
            </t-select>
          </t-form-item>
          <t-form-item label="风险等级" name="riskLevel">
            <t-select v-model="searchForm.riskLevel" placeholder="请选择风险等级" clearable>
              <t-option value="low" label="低风险" />
              <t-option value="medium" label="中风险" />
              <t-option value="high" label="高风险" />
            </t-select>
          </t-form-item>
          <t-form-item>
            <t-button theme="primary" type="submit">搜索</t-button>
            <t-button @click="resetSearch">重置</t-button>
          </t-form-item>
        </t-form>
      </div>

      <!-- 操作栏 -->
      <div class="action-section">
        <t-button theme="primary" @click="addStock">
          <t-icon name="add" />
          新增股票
        </t-button>
        <t-button theme="default" @click="importStocks">
          <t-icon name="upload" />
          批量导入
        </t-button>
        <t-button theme="default" @click="exportStocks">
          <t-icon name="download" />
          导出数据
        </t-button>
      </div>

      <!-- 股票列表 -->
      <t-table
        :data="filteredStocks"
        :columns="columns"
        :pagination="pagination"
        @page-change="handlePageChange"
        row-key="id"
      >
        <!-- 股票代码和名称 -->
        <template #code="{ row }">
          <div class="stock-code-name">
            <div class="stock-code">{{ row.code }}</div>
            <div class="stock-name">{{ row.name }}</div>
          </div>
        </template>

        <!-- 市场类型 -->
        <template #marketType="{ row }">
          <t-tag :theme="getMarketTheme(row.marketType)">
            {{ getMarketText(row.marketType) }}
          </t-tag>
        </template>

        <!-- 分类 -->
        <template #category="{ row }">
          <t-tag :theme="getCategoryTheme(row.category)">
            {{ getCategoryText(row.category) }}
          </t-tag>
        </template>

        <!-- 行业 -->
        <template #industry="{ row }">
          <span>{{ getIndustryText(row.industry) }}</span>
        </template>

        <!-- 市值 -->
        <template #marketCap="{ row }">
          <span class="market-cap">{{ formatMarketCap(row.marketCap) }}</span>
        </template>

        <!-- 市盈率 -->
        <template #pe="{ row }">
          <span :class="getPETheme(row.pe)">{{ row.pe }}</span>
        </template>

        <!-- 风险等级 -->
        <template #riskLevel="{ row }">
          <t-tag :theme="getRiskTheme(row.riskLevel)">
            {{ getRiskLevelText(row.riskLevel) }}
          </t-tag>
        </template>

        <!-- 标签 -->
        <template #tags="{ row }">
          <div class="stock-tags">
            <t-tag 
              v-for="tag in row.tags.slice(0, 2)" 
              :key="tag" 
              size="small"
              theme="default"
            >
              {{ tag }}
            </t-tag>
            <t-tag v-if="row.tags.length > 2" size="small" theme="default">
              +{{ row.tags.length - 2 }}
            </t-tag>
          </div>
        </template>

        <!-- 操作 -->
        <template #operation="{ row }">
          <t-button theme="default" variant="text" size="small" @click="viewStock(row)">
            查看
          </t-button>
          <t-button theme="default" variant="text" size="small" @click="editStock(row)">
            编辑
          </t-button>
          <t-button 
            :theme="row.enabled ? 'warning' : 'success'" 
            variant="text" 
            size="small" 
            @click="toggleStock(row)"
          >
            {{ row.enabled ? '禁用' : '启用' }}
          </t-button>
          <t-button theme="danger" variant="text" size="small" @click="deleteStock(row)">
            删除
          </t-button>
        </template>
      </t-table>
    </t-card>

    <!-- 股票详情弹窗 -->
    <t-dialog
      v-model:visible="showDetailDialog"
      :header="dialogTitle"
      width="800px"
      :footer="false"
    >
      <div v-if="currentStock" class="stock-detail">
        <t-descriptions :data="getStockDetailData()" :column="2" />
      </div>
    </t-dialog>

    <!-- 股票编辑弹窗 -->
    <t-dialog
      v-model:visible="showEditDialog"
      :header="editDialogTitle"
      width="600px"
      @confirm="confirmEdit"
    >
      <t-form
        ref="editFormRef"
        :model="editForm"
        :rules="editRules"
        label-width="100px"
      >
        <t-form-item label="股票代码" name="code">
          <t-input v-model="editForm.code" placeholder="请输入股票代码" />
        </t-form-item>
        <t-form-item label="股票名称" name="name">
          <t-input v-model="editForm.name" placeholder="请输入股票名称" />
        </t-form-item>
        <t-form-item label="股票分类" name="category">
          <t-select v-model="editForm.category" placeholder="请选择分类">
            <t-option value="main_board" label="主板" />
            <t-option value="sme_board" label="中小板" />
            <t-option value="gem_board" label="创业板" />
            <t-option value="star_board" label="科创板" />
            <t-option value="new_third_board" label="新三板" />
          </t-select>
        </t-form-item>
        <t-form-item label="行业分类" name="industry">
          <t-select v-model="editForm.industry" placeholder="请选择行业">
            <t-option value="technology" label="科技" />
            <t-option value="finance" label="金融" />
            <t-option value="healthcare" label="医疗" />
            <t-option value="consumer" label="消费" />
            <t-option value="industry" label="工业" />
            <t-option value="energy" label="能源" />
            <t-option value="materials" label="材料" />
            <t-option value="utilities" label="公用事业" />
            <t-option value="real_estate" label="房地产" />
            <t-option value="communication" label="通信" />
          </t-select>
        </t-form-item>
        <t-form-item label="上市日期" name="listingDate">
          <t-date-picker v-model="editForm.listingDate" placeholder="请选择上市日期" />
        </t-form-item>
        <t-form-item label="市值（元）" name="marketCap">
          <t-input-number v-model="editForm.marketCap" placeholder="请输入市值" :min="0" />
        </t-form-item>
        <t-form-item label="市盈率" name="pe">
          <t-input-number v-model="editForm.pe" placeholder="请输入市盈率" :min="0" :precision="2" />
        </t-form-item>
        <t-form-item label="市净率" name="pb">
          <t-input-number v-model="editForm.pb" placeholder="请输入市净率" :min="0" :precision="2" />
        </t-form-item>
        <t-form-item label="股息率" name="dividend">
          <t-input-number v-model="editForm.dividend" placeholder="请输入股息率" :min="0" :precision="3" />
        </t-form-item>
        <t-form-item label="波动率" name="volatility">
          <t-input-number v-model="editForm.volatility" placeholder="请输入波动率" :min="0" :max="1" :precision="2" />
        </t-form-item>
        <t-form-item label="流动性" name="liquidity">
          <t-select v-model="editForm.liquidity" placeholder="请选择流动性">
            <t-option value="high" label="高流动性" />
            <t-option value="medium" label="中流动性" />
            <t-option value="low" label="低流动性" />
          </t-select>
        </t-form-item>
        <t-form-item label="风险等级" name="riskLevel">
          <t-select v-model="editForm.riskLevel" placeholder="请选择风险等级">
            <t-option value="low" label="低风险" />
            <t-option value="medium" label="中风险" />
            <t-option value="high" label="高风险" />
          </t-select>
        </t-form-item>
        <t-form-item label="股票描述" name="description">
          <t-textarea v-model="editForm.description" placeholder="请输入股票描述" :maxlength="200" />
        </t-form-item>
        <t-form-item label="标签" name="tags">
          <t-tag-input v-model="editForm.tags" placeholder="请输入标签，按回车添加" />
        </t-form-item>
      </t-form>
    </t-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { MessagePlugin, DialogPlugin } from 'tdesign-vue-next'
import { 
  stockConfig, 
  getEnabledStocks,
  updateStockConfig,
  addStock as addStockUtil,
  deleteStock as deleteStockUtil,
  searchStocks,
  getCategoryText,
  getIndustryText,
  getRiskLevelText,
  getLiquidityText,
  getMarketText,
  getCurrencyText,
  formatMarketCap,
  STOCK_CATEGORIES,
  INDUSTRY_CATEGORIES,
  STOCK_MARKETS,
  CURRENCY_TYPES
} from '@/utils/stockConfig'

// 搜索表单
const searchForm = reactive({
  keyword: '',
  marketType: '',
  category: '',
  industry: '',
  riskLevel: ''
})

// 分页
const pagination = reactive({
  current: 1,
  pageSize: 20,
  total: 0
})

// 表格列定义
const columns = [
  { colKey: 'code', title: '股票信息', width: 150 },
  { colKey: 'marketType', title: '市场', width: 80 },
  { colKey: 'category', title: '分类', width: 100 },
  { colKey: 'industry', title: '行业', width: 100 },
  { colKey: 'marketCap', title: '市值', width: 120 },
  { colKey: 'pe', title: '市盈率', width: 100 },
  { colKey: 'pb', title: '市净率', width: 100 },
  { colKey: 'dividend', title: '股息率', width: 100 },
  { colKey: 'riskLevel', title: '风险等级', width: 100 },
  { colKey: 'tags', title: '标签', width: 150 },
  { colKey: 'operation', title: '操作', width: 200, fixed: 'right' }
]

// 弹窗状态
const showDetailDialog = ref(false)
const showEditDialog = ref(false)
const currentStock = ref(null)
const isEditMode = ref(false)

// 编辑表单
const editForm = reactive({
  id: '',
  code: '',
  name: '',
  category: '',
  industry: '',
  market: '',
  listingDate: '',
  marketCap: null,
  pe: null,
  pb: null,
  dividend: null,
  volatility: null,
  liquidity: '',
  riskLevel: '',
  description: '',
  tags: []
})

const editRules = {
  code: [{ required: true, message: '请输入股票代码', type: 'error' }],
  name: [{ required: true, message: '请输入股票名称', type: 'error' }],
  category: [{ required: true, message: '请选择股票分类', type: 'error' }],
  industry: [{ required: true, message: '请选择行业分类', type: 'error' }]
}

// 计算属性
const allStocks = computed(() => stockConfig)
const filteredStocks = computed(() => {
  let stocks = allStocks.value
  
  if (searchForm.keyword) {
    stocks = searchStocks(searchForm.keyword)
  }
  
  if (searchForm.marketType) {
    stocks = stocks.filter(s => s.marketType === searchForm.marketType)
  }
  
  if (searchForm.category) {
    stocks = stocks.filter(s => s.category === searchForm.category)
  }
  
  if (searchForm.industry) {
    stocks = stocks.filter(s => s.industry === searchForm.industry)
  }
  
  if (searchForm.riskLevel) {
    stocks = stocks.filter(s => s.riskLevel === searchForm.riskLevel)
  }
  
  pagination.total = stocks.length
  const start = (pagination.current - 1) * pagination.pageSize
  const end = start + pagination.pageSize
  return stocks.slice(start, end)
})

const dialogTitle = computed(() => 
  isEditMode.value ? '编辑股票' : '股票详情'
)

const editDialogTitle = computed(() => 
  isEditMode.value ? '编辑股票' : '新增股票'
)

// 方法
const getMarketTheme = (marketType) => {
  const themes = {
    [STOCK_MARKETS.A_SHARE]: 'primary',
    [STOCK_MARKETS.HK_STOCK]: 'success',
    [STOCK_MARKETS.US_STOCK]: 'warning',
    [STOCK_MARKETS.OTHER]: 'default'
  }
  return themes[marketType] || 'default'
}

const getCategoryTheme = (category) => {
  const themes = {
    [STOCK_CATEGORIES.MAIN_BOARD]: 'primary',
    [STOCK_CATEGORIES.SME_BOARD]: 'success',
    [STOCK_CATEGORIES.GEM_BOARD]: 'warning',
    [STOCK_CATEGORIES.STAR_BOARD]: 'danger',
    [STOCK_CATEGORIES.NEW_THIRD_BOARD]: 'default',
    [STOCK_CATEGORIES.HK_MAIN]: 'success',
    [STOCK_CATEGORIES.HK_GEM]: 'warning',
    [STOCK_CATEGORIES.US_NYSE]: 'primary',
    [STOCK_CATEGORIES.US_NASDAQ]: 'warning',
    [STOCK_CATEGORIES.US_AMEX]: 'danger'
  }
  return themes[category] || 'default'
}

const getRiskTheme = (riskLevel) => {
  const themes = {
    low: 'success',
    medium: 'warning',
    high: 'danger'
  }
  return themes[riskLevel] || 'default'
}

const getPETheme = (pe) => {
  if (pe < 10) return 'pe-low'
  if (pe < 20) return 'pe-medium'
  return 'pe-high'
}

const getStockDetailData = () => {
  if (!currentStock.value) return []
  
  return [
    { label: '股票代码', content: currentStock.value.code },
    { label: '股票名称', content: currentStock.value.name },
    { label: '股票分类', content: getCategoryText(currentStock.value.category) },
    { label: '行业分类', content: getIndustryText(currentStock.value.industry) },
    { label: '交易市场', content: currentStock.value.market },
    { label: '上市日期', content: currentStock.value.listingDate },
    { label: '市值', content: formatMarketCap(currentStock.value.marketCap) },
    { label: '市盈率', content: currentStock.value.pe },
    { label: '市净率', content: currentStock.value.pb },
    { label: '股息率', content: `${(currentStock.value.dividend * 100).toFixed(2)}%` },
    { label: '波动率', content: `${(currentStock.value.volatility * 100).toFixed(2)}%` },
    { label: '流动性', content: getLiquidityText(currentStock.value.liquidity) },
    { label: '风险等级', content: getRiskLevelText(currentStock.value.riskLevel) },
    { label: '股票描述', content: currentStock.value.description },
    { label: '标签', content: currentStock.value.tags.join('、') }
  ]
}

const handleSearch = () => {
  pagination.current = 1
}

const resetSearch = () => {
  searchForm.keyword = ''
  searchForm.marketType = ''
  searchForm.category = ''
  searchForm.industry = ''
  searchForm.riskLevel = ''
  pagination.current = 1
}

const handlePageChange = (pageInfo) => {
  pagination.current = pageInfo.current
  pagination.pageSize = pageInfo.pageSize
}

const viewStock = (stock) => {
  currentStock.value = stock
  showDetailDialog.value = true
  isEditMode.value = false
}

const addStock = () => {
  resetEditForm()
  showEditDialog.value = true
  isEditMode.value = false
}

const editStock = (stock) => {
  currentStock.value = stock
  Object.assign(editForm, {
    id: stock.id,
    code: stock.code,
    name: stock.name,
    category: stock.category,
    industry: stock.industry,
    market: stock.market,
    listingDate: stock.listingDate,
    marketCap: stock.marketCap,
    pe: stock.pe,
    pb: stock.pb,
    dividend: stock.dividend,
    volatility: stock.volatility,
    liquidity: stock.liquidity,
    riskLevel: stock.riskLevel,
    description: stock.description,
    tags: [...stock.tags]
  })
  showEditDialog.value = true
  isEditMode.value = true
}

const toggleStock = (stock) => {
  updateStockConfig(stock.id, { enabled: !stock.enabled })
  MessagePlugin.success(`${stock.enabled ? '禁用' : '启用'}成功`)
}

const deleteStock = (stock) => {
  DialogPlugin.confirm({
    header: '确认删除',
    body: `确定要删除股票"${stock.name}"吗？`,
    onConfirm: () => {
      deleteStockUtil(stock.id)
      MessagePlugin.success('删除成功')
    }
  })
}

const confirmEdit = () => {
  if (isEditMode.value) {
    // 编辑模式
    const updates = {
      code: editForm.code,
      name: editForm.name,
      category: editForm.category,
      industry: editForm.industry,
      market: editForm.market,
      listingDate: editForm.listingDate,
      marketCap: editForm.marketCap,
      pe: editForm.pe,
      pb: editForm.pb,
      dividend: editForm.dividend,
      volatility: editForm.volatility,
      liquidity: editForm.liquidity,
      riskLevel: editForm.riskLevel,
      description: editForm.description,
      tags: editForm.tags
    }
    updateStockConfig(editForm.id, updates)
    MessagePlugin.success('编辑成功')
  } else {
    // 新增模式
    const newStock = {
      id: `stock_${Date.now()}`,
      code: editForm.code,
      name: editForm.name,
      category: editForm.category,
      industry: editForm.industry,
      market: editForm.market,
      listingDate: editForm.listingDate,
      marketCap: editForm.marketCap,
      pe: editForm.pe,
      pb: editForm.pb,
      dividend: editForm.dividend,
      volatility: editForm.volatility,
      liquidity: editForm.liquidity,
      riskLevel: editForm.riskLevel,
      description: editForm.description,
      tags: editForm.tags,
      enabled: true
    }
    addStockUtil(newStock)
    MessagePlugin.success('新增成功')
  }
  
  showEditDialog.value = false
}

const resetEditForm = () => {
  Object.assign(editForm, {
    id: '',
    code: '',
    name: '',
    category: '',
    industry: '',
    market: '',
    listingDate: '',
    marketCap: null,
    pe: null,
    pb: null,
    dividend: null,
    volatility: null,
    liquidity: '',
    riskLevel: '',
    description: '',
    tags: []
  })
}

const importStocks = () => {
  MessagePlugin.info('批量导入功能开发中...')
}

const exportStocks = () => {
  MessagePlugin.info('导出数据功能开发中...')
}

onMounted(() => {
  // 初始化数据
})
</script>

<style scoped>
.stock-management {
  padding: 16px;
}

.stock-card {
  margin-bottom: 16px;
}

.search-section {
  margin-bottom: 16px;
  padding: 16px;
  background: #f8f9fa;
  border-radius: 6px;
}

.action-section {
  margin-bottom: 16px;
  display: flex;
  gap: 12px;
}

.stock-code-name {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.stock-code {
  font-weight: 600;
  color: #0052d9;
  font-size: 14px;
}

.stock-name {
  font-size: 12px;
  color: #666;
}

.market-cap {
  font-weight: 500;
  color: #52c41a;
}

.pe-low {
  color: #52c41a;
  font-weight: 500;
}

.pe-medium {
  color: #fa8c16;
  font-weight: 500;
}

.pe-high {
  color: #ff4d4f;
  font-weight: 500;
}

.stock-tags {
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.stock-detail {
  padding: 16px 0;
}
</style>
