<template>
  <div class="trading-strategy-settings">
    <t-card title="交易策略管理" class="strategy-card">
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
              <span class="holding-time">持仓: {{ strategy.parameters.holdingTime }}</span>
            </div>
            <div class="strategy-details">
              <div class="pros-cons">
                <div class="pros">
                  <strong>优点：</strong>{{ strategy.pros.join('、') }}
                </div>
                <div class="cons">
                  <strong>缺点：</strong>{{ strategy.cons.join('、') }}
                </div>
              </div>
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
      width="700px"
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
            <t-option value="scalping" label="超短线" />
            <t-option value="day_trading" label="日内交易" />
            <t-option value="swing" label="波段交易" />
            <t-option value="position" label="持仓交易" />
            <t-option value="arbitrage" label="套利交易" />
          </t-select>
        </t-form-item>
        <t-form-item label="胜率预期" name="winRate">
          <t-input v-model="editForm.winRate" placeholder="如：60-70%" />
        </t-form-item>
        <t-form-item label="适用场景" name="suitableFor">
          <t-textarea v-model="editForm.suitableFor" placeholder="请输入适用场景" />
        </t-form-item>
        <t-form-item label="持仓时间" name="holdingTime">
          <t-input v-model="editForm.holdingTime" placeholder="如：几分钟到几小时" />
        </t-form-item>
        <t-form-item label="资金要求" name="capitalRequirement">
          <t-input v-model="editForm.capitalRequirement" placeholder="如：高资金要求" />
        </t-form-item>
        <t-form-item label="技能要求" name="skillLevel">
          <t-input v-model="editForm.skillLevel" placeholder="如：需要专业交易技能" />
        </t-form-item>
        <t-form-item label="风险等级" name="riskLevel">
          <t-select v-model="editForm.riskLevel" placeholder="请选择风险等级">
            <t-option value="low" label="低风险" />
            <t-option value="medium" label="中风险" />
            <t-option value="high" label="高风险" />
          </t-select>
        </t-form-item>
        <t-form-item label="优点" name="pros">
          <t-input v-model="editForm.pros" placeholder="如：快速获利、资金周转快" />
        </t-form-item>
        <t-form-item label="缺点" name="cons">
          <t-input v-model="editForm.cons" placeholder="如：需要专业技能、手续费成本高" />
        </t-form-item>
      </t-form>
    </t-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted } from 'vue'
import { MessagePlugin, DialogPlugin } from 'tdesign-vue-next'
import { 
  tradingStrategyConfig, 
  getEnabledTradingStrategies,
  updateTradingStrategyConfig,
  addTradingStrategy as addTradingStrategyUtil,
  deleteTradingStrategy as deleteTradingStrategyUtil,
  TRADING_STRATEGY_CATEGORIES,
  RISK_LEVELS,
  getTradingStrategyCategoryText,
  getRiskLevelText
} from '@/utils/tradingStrategyConfig'

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
  holdingTime: '',
  capitalRequirement: '',
  skillLevel: '',
  riskLevel: '',
  pros: '',
  cons: ''
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
    [TRADING_STRATEGY_CATEGORIES.SCALPING]: 'danger',
    [TRADING_STRATEGY_CATEGORIES.DAY_TRADING]: 'warning',
    [TRADING_STRATEGY_CATEGORIES.SWING]: 'primary',
    [TRADING_STRATEGY_CATEGORIES.POSITION]: 'success',
    [TRADING_STRATEGY_CATEGORIES.ARBITRAGE]: 'default'
  }
  return themes[category] || 'default'
}

const getCategoryText = (category) => {
  return getTradingStrategyCategoryText(category)
}

const loadStrategies = () => {
  strategies.value = [...tradingStrategyConfig]
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
    holdingTime: strategy.parameters.holdingTime,
    capitalRequirement: strategy.parameters.capitalRequirement,
    skillLevel: strategy.parameters.skillLevel,
    riskLevel: strategy.parameters.riskLevel,
    pros: strategy.pros.join('、'),
    cons: strategy.cons.join('、')
  })
  showEditDialog.value = true
  isEditMode.value = true
}

const toggleStrategy = (strategy) => {
  updateTradingStrategyConfig(strategy.id, { enabled: !strategy.enabled })
  loadStrategies()
  MessagePlugin.success(`${strategy.enabled ? '禁用' : '启用'}成功`)
}

const deleteStrategy = (strategy) => {
  DialogPlugin.confirm({
    header: '确认删除',
    body: `确定要删除策略"${strategy.name}"吗？`,
    onConfirm: () => {
      deleteTradingStrategyUtil(strategy.id)
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
        holdingTime: editForm.holdingTime,
        capitalRequirement: editForm.capitalRequirement,
        skillLevel: editForm.skillLevel,
        riskLevel: editForm.riskLevel
      },
      pros: editForm.pros.split('、'),
      cons: editForm.cons.split('、')
    }
    updateTradingStrategyConfig(editForm.id, updates)
    MessagePlugin.success('编辑成功')
  } else {
    // 新增模式
    const newStrategy = {
      id: `trading_strategy_${Date.now()}`,
      name: editForm.name,
      description: editForm.description,
      detail: editForm.detail,
      category: editForm.category,
      winRate: editForm.winRate,
      suitableFor: editForm.suitableFor,
      parameters: {
        holdingTime: editForm.holdingTime,
        capitalRequirement: editForm.capitalRequirement,
        skillLevel: editForm.skillLevel,
        riskLevel: editForm.riskLevel
      },
      pros: editForm.pros.split('、'),
      cons: editForm.cons.split('、'),
      enabled: true
    }
    addTradingStrategyUtil(newStrategy)
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
    holdingTime: '',
    capitalRequirement: '',
    skillLevel: '',
    riskLevel: '',
    pros: '',
    cons: ''
  })
}

const resetToDefault = () => {
  DialogPlugin.confirm({
    header: '确认重置',
    body: '确定要重置为默认交易策略配置吗？这将覆盖所有自定义设置。',
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
.trading-strategy-settings {
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
  align-items: flex-start;
  padding: 20px;
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
  margin-right: 16px;
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
  margin-bottom: 12px;
}

.win-rate {
  color: #52c41a;
  font-weight: 500;
}

.risk-level {
  color: #fa8c16;
  font-weight: 500;
}

.holding-time {
  color: #1890ff;
  font-weight: 500;
}

.strategy-details {
  font-size: 12px;
  color: #666;
}

.pros-cons {
  display: flex;
  gap: 20px;
}

.pros {
  color: #52c41a;
}

.cons {
  color: #ff4d4f;
}

.strategy-actions {
  display: flex;
  flex-direction: column;
  gap: 8px;
  min-width: 120px;
}
</style>
