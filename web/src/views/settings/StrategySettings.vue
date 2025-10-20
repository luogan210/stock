<template>
  <div class="strategy-settings">
    <t-card title="选股策略管理" class="strategy-card">
      <!-- 操作栏 -->
      <div class="action-section">
        <t-button theme="primary" @click="addStrategy">
          <t-icon name="add" />
          新增策略
        </t-button>
        <t-button theme="default" @click="resetToDefault">
          <t-icon name="refresh" />
          重置为默认
        </t-button>
      </div>

      <!-- 策略列表 -->
      <div class="strategy-list">
        <div 
          v-for="strategy in strategies" 
          :key="strategy.id" 
          class="strategy-item"
          :class="{ disabled: !strategy.enabled }"
        >
          <div class="strategy-info">
            <div class="strategy-header">
              <h4>{{ strategy.name }}</h4>
              <t-tag :theme="getCategoryTheme(strategy.category)">
                {{ getCategoryText(strategy.category) }}
              </t-tag>
            </div>
            <p class="strategy-desc">{{ strategy.description }}</p>
            <div class="strategy-meta">
              <span class="win-rate">胜率: {{ strategy.winRate }}</span>
              <span class="risk-level">风险: {{ getRiskLevelText(strategy.parameters.riskLevel) }}</span>
              <span class="timeframe">时间: {{ strategy.parameters.timeframe }}</span>
            </div>
          </div>
          
          <div class="strategy-actions">
            <t-button 
              theme="default" 
              variant="text" 
              size="small" 
              @click="editStrategy(strategy)"
            >
              编辑
            </t-button>
            <t-button 
              :theme="strategy.enabled ? 'warning' : 'success'" 
              variant="text" 
              size="small" 
              @click="toggleStrategy(strategy)"
            >
              {{ strategy.enabled ? '禁用' : '启用' }}
            </t-button>
            <t-button 
              theme="danger" 
              variant="text" 
              size="small" 
              @click="deleteStrategy(strategy)"
            >
              删除
            </t-button>
          </div>
        </div>
      </div>
    </t-card>

    <!-- 策略编辑弹窗 -->
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
        <t-form-item label="策略名称" name="name">
          <t-input v-model="editForm.name" placeholder="请输入策略名称" />
        </t-form-item>
        <t-form-item label="策略描述" name="description">
          <t-input v-model="editForm.description" placeholder="请输入策略描述" />
        </t-form-item>
        <t-form-item label="详细说明" name="detail">
          <t-textarea v-model="editForm.detail" placeholder="请输入详细说明" :maxlength="500" />
        </t-form-item>
        <t-form-item label="策略分类" name="category">
          <t-select v-model="editForm.category" placeholder="请选择分类">
            <t-option value="technical" label="技术分析" />
            <t-option value="fundamental" label="基本面分析" />
          </t-select>
        </t-form-item>
        <t-form-item label="胜率预期" name="winRate">
          <t-input v-model="editForm.winRate" placeholder="如：60-70%" />
        </t-form-item>
        <t-form-item label="适用场景" name="suitableFor">
          <t-textarea v-model="editForm.suitableFor" placeholder="请输入适用场景" />
        </t-form-item>
        <t-form-item label="时间框架" name="timeframe">
          <t-input v-model="editForm.timeframe" placeholder="如：日线、小时线" />
        </t-form-item>
        <t-form-item label="技术指标" name="indicators">
          <t-input v-model="editForm.indicators" placeholder="如：MA、MACD、RSI" />
        </t-form-item>
        <t-form-item label="风险等级" name="riskLevel">
          <t-select v-model="editForm.riskLevel" placeholder="请选择风险等级">
            <t-option value="low" label="低风险" />
            <t-option value="medium" label="中风险" />
            <t-option value="high" label="高风险" />
          </t-select>
        </t-form-item>
      </t-form>
    </t-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { MessagePlugin, DialogPlugin } from 'tdesign-vue-next'
import { 
  strategyConfig, 
  getEnabledStrategies,
  updateStrategyConfig,
  addStrategy as addStrategyUtil,
  deleteStrategy as deleteStrategyUtil,
  STRATEGY_CATEGORIES,
  RISK_LEVELS
} from '@/utils/strategyConfig'

// 策略列表
const strategies = ref([])

// 弹窗状态
const showEditDialog = ref(false)
const isEditMode = ref(false)
const editForm = reactive({
  id: '',
  name: '',
  description: '',
  detail: '',
  category: '',
  winRate: '',
  suitableFor: '',
  timeframe: '',
  indicators: '',
  riskLevel: ''
})

const editRules = {
  name: [{ required: true, message: '请输入策略名称', type: 'error' }],
  description: [{ required: true, message: '请输入策略描述', type: 'error' }],
  category: [{ required: true, message: '请选择策略分类', type: 'error' }]
}

// 计算属性
const editDialogTitle = computed(() => 
  isEditMode.value ? '编辑策略' : '新增策略'
)

// 方法
const getCategoryTheme = (category) => {
  const themes = {
    [STRATEGY_CATEGORIES.TECHNICAL]: 'primary',
    [STRATEGY_CATEGORIES.FUNDAMENTAL]: 'success'
  }
  return themes[category] || 'default'
}

const getCategoryText = (category) => {
  const texts = {
    [STRATEGY_CATEGORIES.TECHNICAL]: '技术分析',
    [STRATEGY_CATEGORIES.FUNDAMENTAL]: '基本面分析'
  }
  return texts[category] || '未知'
}

const getRiskLevelText = (riskLevel) => {
  const texts = {
    [RISK_LEVELS.LOW]: '低风险',
    [RISK_LEVELS.MEDIUM]: '中风险',
    [RISK_LEVELS.HIGH]: '高风险'
  }
  return texts[riskLevel] || '未知'
}

const loadStrategies = () => {
  strategies.value = [...strategyConfig]
}

const addStrategy = () => {
  resetEditForm()
  showEditDialog.value = true
  isEditMode.value = false
}

const editStrategy = (strategy) => {
  Object.assign(editForm, {
    id: strategy.id,
    name: strategy.name,
    description: strategy.description,
    detail: strategy.detail,
    category: strategy.category,
    winRate: strategy.winRate,
    suitableFor: strategy.suitableFor,
    timeframe: strategy.parameters.timeframe,
    indicators: strategy.parameters.indicators.join('、'),
    riskLevel: strategy.parameters.riskLevel
  })
  showEditDialog.value = true
  isEditMode.value = true
}

const toggleStrategy = (strategy) => {
  updateStrategyConfig(strategy.id, { enabled: !strategy.enabled })
  loadStrategies()
  MessagePlugin.success(`${strategy.enabled ? '禁用' : '启用'}成功`)
}

const deleteStrategy = (strategy) => {
  DialogPlugin.confirm({
    header: '确认删除',
    body: `确定要删除策略"${strategy.name}"吗？`,
    onConfirm: () => {
      deleteStrategyUtil(strategy.id)
      loadStrategies()
      MessagePlugin.success('删除成功')
    }
  })
}

const confirmEdit = () => {
  if (isEditMode.value) {
    // 编辑模式
    const updates = {
      name: editForm.name,
      description: editForm.description,
      detail: editForm.detail,
      category: editForm.category,
      winRate: editForm.winRate,
      suitableFor: editForm.suitableFor,
      parameters: {
        timeframe: editForm.timeframe,
        indicators: editForm.indicators.split('、'),
        riskLevel: editForm.riskLevel
      }
    }
    updateStrategyConfig(editForm.id, updates)
    MessagePlugin.success('编辑成功')
  } else {
    // 新增模式
    const newStrategy = {
      id: `strategy_${Date.now()}`,
      name: editForm.name,
      description: editForm.description,
      detail: editForm.detail,
      category: editForm.category,
      winRate: editForm.winRate,
      suitableFor: editForm.suitableFor,
      parameters: {
        timeframe: editForm.timeframe,
        indicators: editForm.indicators.split('、'),
        riskLevel: editForm.riskLevel
      },
      enabled: true
    }
    addStrategyUtil(newStrategy)
    MessagePlugin.success('新增成功')
  }
  
  loadStrategies()
  showEditDialog.value = false
}

const resetEditForm = () => {
  Object.assign(editForm, {
    id: '',
    name: '',
    description: '',
    detail: '',
    category: '',
    winRate: '',
    suitableFor: '',
    timeframe: '',
    indicators: '',
    riskLevel: ''
  })
}

const resetToDefault = () => {
  DialogPlugin.confirm({
    header: '确认重置',
    body: '确定要重置为默认策略配置吗？这将覆盖所有自定义设置。',
    onConfirm: () => {
      // 重新加载默认配置
      location.reload()
    }
  })
}

onMounted(() => {
  loadStrategies()
})
</script>

<style scoped>
.strategy-settings {
  padding: 16px;
}

.strategy-card {
  margin-bottom: 16px;
}

.action-section {
  margin-bottom: 20px;
  display: flex;
  gap: 12px;
}

.strategy-list {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.strategy-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 16px;
  border: 1px solid #e7e7e7;
  border-radius: 8px;
  background: #fff;
  transition: all 0.2s;
}

.strategy-item:hover {
  border-color: #0052d9;
  box-shadow: 0 2px 8px rgba(0, 82, 217, 0.1);
}

.strategy-item.disabled {
  opacity: 0.6;
  background: #f5f5f5;
}

.strategy-info {
  flex: 1;
}

.strategy-header {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 8px;
}

.strategy-header h4 {
  margin: 0;
  color: #0052d9;
  font-size: 16px;
  font-weight: 600;
}

.strategy-desc {
  margin: 0 0 8px 0;
  color: #666;
  font-size: 14px;
}

.strategy-meta {
  display: flex;
  gap: 16px;
  font-size: 12px;
  color: #999;
}

.win-rate {
  color: #52c41a;
  font-weight: 500;
}

.risk-level {
  color: #fa8c16;
  font-weight: 500;
}

.timeframe {
  color: #1890ff;
  font-weight: 500;
}

.strategy-actions {
  display: flex;
  gap: 8px;
}
</style>
