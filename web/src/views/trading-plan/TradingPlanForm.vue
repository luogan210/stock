<template>
  <div class="trading-plan-form">
    <t-form
      ref="formRef"
      :data="formData"
      :rules="formRules"
      label-width="120px"
      scroll-to-first-error="smooth"
    >
      <t-form-item label="选择股票" name="stockCode">
        <t-select
          class="form-input-md"
          v-model="formData.stockCode"
          placeholder="请选择股票"
          @change="handleStockChange"
        >
          <t-option
            v-for="stock in availableStocks"
            :key="stock.code"
            :value="stock.code"
            :label="`${stock.name} (${stock.code}) - ${getStockCategoryText(stock.category)}`"
          />
        </t-select>
      </t-form-item>

      <t-form-item label="选股策略" name="strategy">
        <div class="strategy-field">
          <t-select class="form-input-md" v-model="formData.strategy" placeholder="请选择选股策略" @change="handleStrategyChange">
            <t-option
              v-for="strategy in availableStrategies"
              :key="strategy.id"
              :value="strategy.id"
              :label="`${strategy.name} - ${strategy.description}`"
            />
          </t-select>
          <t-popup
            :visible="true"
            placement="right"
            trigger="click"
            :overlay-style="{ padding: '8px' }"
            :destroy-on-close="false"
            :z-index="9999"
            :show-arrow="true"
          >
            <t-button
              variant="text"
              shape="circle"
              size="small"
            >
              <t-icon name="help-circle" />
            </t-button>
            <template #content>
              <div class="strategy-popup-content">
                <div v-if="formData.strategy">
                  <h4>{{ getStrategyById(formData.strategy)?.name || '策略详情' }}</h4>
                  <ul>
                    <li><strong>策略描述</strong>：{{ getStrategyById(formData.strategy)?.detail || '暂无详细说明' }}</li>
                    <li><strong>胜率预期</strong>：{{ getStrategyById(formData.strategy)?.winRate || '未知' }}</li>
                    <li><strong>适用场景</strong>：{{ getStrategyById(formData.strategy)?.suitableFor || '未知' }}</li>
                    <li><strong>风险等级</strong>：{{ getStrategyById(formData.strategy)?.parameters?.riskLevel || '未知' }}</li>
                    <li><strong>时间框架</strong>：{{ getStrategyById(formData.strategy)?.parameters?.timeframe || '未知' }}</li>
                    <li><strong>技术指标</strong>：{{ getStrategyById(formData.strategy)?.parameters?.indicators?.join('、') || '未知' }}</li>
                  </ul>
                </div>
                <div v-else>
                  <p>请先选择一个策略</p>
                </div>
              </div>
            </template>
          </t-popup>
        </div>
      </t-form-item>

      <t-form-item label="计划名称" name="name">
        <t-input
          class="form-input-md"
          v-model="formData.name"
          placeholder="请输入计划名称"
        />
      </t-form-item>

      <t-form-item label="交易方向" name="type">
        <t-radio-group v-model="formData.type" @change="generatePlanName">
          <t-radio value="buy">买多</t-radio>
          <t-radio value="sell">买空</t-radio>
        </t-radio-group>
      </t-form-item>

      <t-form-item label="交易策略" name="tradingStrategy">
        <t-select class="form-input-md" v-model="formData.tradingStrategy" placeholder="请选择交易策略" @change="updateRiskLevel">
          <t-option
            v-for="strategy in availableTradingStrategies"
            :key="strategy.id"
            :value="strategy.id"
            :label="`${strategy.name} - ${strategy.description}`"
          />
        </t-select>
      </t-form-item>

      <!-- 风险等级显示 -->
      <t-form-item label="风险等级">
        <div class="risk-level-display">
          <t-tag
            :color="riskInfo.color"
            variant="light-outline"
            size="medium"
          >
            {{ riskInfo.text }}
          </t-tag>
          <span class="risk-score">评分: {{ riskInfo.score }}/3</span>
          <t-popup
            placement="right"
            trigger="hover"
            :overlay-style="{ padding: '12px', maxWidth: '400px' }"
          >
            <t-button variant="text" size="small">
              <t-icon name="help-circle" />
            </t-button>
            <template #content>
              <div class="risk-details">
                <h4>风险分析详情</h4>
                <p><strong>选股策略风险:</strong> {{ getRiskLevelText(riskInfo.details?.strategyRisk) }}</p>
                <p><strong>交易策略风险:</strong> {{ getRiskLevelText(riskInfo.details?.tradingStrategyRisk) }}</p>
                <p><strong>综合评分:</strong> {{ riskInfo.details?.combinedScore }}/3</p>
                <p><strong>风险说明:</strong> {{ riskInfo.details?.explanation }}</p>
                <div class="risk-suggestions" v-if="riskSuggestions">
                  <h5>投资建议</h5>
                  <ul>
                    <li>{{ riskSuggestions.positionSize }}</li>
                    <li>{{ riskSuggestions.stopLoss }}</li>
                    <li>{{ riskSuggestions.takeProfit }}</li>
                    <li>{{ riskSuggestions.timeframe }}</li>
                  </ul>
                  <h5>注意事项</h5>
                  <ul>
                    <li v-for="tip in riskSuggestions.tips" :key="tip">{{ tip }}</li>
                  </ul>
                </div>
              </div>
            </template>
          </t-popup>
        </div>
      </t-form-item>
      <FlexRow>
        <t-form-item label="计划买进价格" name="targetPrice">
        <t-input-number
          class="form-input-md"
          v-model="formData.targetPrice"
          placeholder="请输入目标价格"
          :min="0"
          :precision="2"
          @change="validateRiskParameters"
        />
      </t-form-item>

      <t-form-item label="计划数量" name="quantity">
        <t-input-number
          class="form-input-md"
          v-model="formData.quantity"
          placeholder="请输入数量"
          :min="1"
        />
      </t-form-item>
      </FlexRow>

      <FlexRow>
        <t-form-item label="止损价格" name="stopLoss">
        <t-input-number
          class="form-input-md"
          v-model="formData.stopLoss"
          placeholder="请输入止损价格"
          :min="0"
          :precision="2"
          @change="validateRiskParameters"
        />
      </t-form-item>

      <t-form-item label="止盈价格" name="takeProfit">
        <t-input-number
          class="form-input-md"
          v-model="formData.takeProfit"
          placeholder="请输入止盈价格"
          :min="0"
          :precision="2"
          @change="validateRiskParameters"
        />
      </t-form-item>
      </FlexRow>

      <!-- 风险参数验证提示 -->
      <t-form-item v-if="riskValidation.warnings && riskValidation.warnings.length > 0">
        <t-alert theme="warning" :message="`风险提示: ${riskValidation.warnings.join('; ')}`" />
      </t-form-item>


      <t-form-item label="计划原因描述" name="description">
        <t-textarea
          class="form-input-lg"
          v-model="formData.description"
          placeholder="请描述选择该股票的原因和计划依据"
          :maxlength="500"
          :autosize="{ minRows: 3, maxRows: 6 }"
        />
      </t-form-item>

      <t-form-item label="备注" name="remark">
        <t-textarea
          class="form-input-lg"
          v-model="formData.remark"
          placeholder="其他备注信息"
          :maxlength="200"
          :autosize="{ minRows: 2, maxRows: 4 }"
        />
      </t-form-item>
    </t-form>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, watch } from 'vue'
import {
  getStrategyById
} from '@/utils/strategyConfig'

import {useStockStore,useStrategyStore,useTradingStrategyStore} from "@/stores/index.js";
import { getStockCategoryText } from '@/views/stock/utils'
import {
  calculateRiskLevel,
  getRiskSuggestions,
  validateRiskParameters as validateRisk,
  RISK_LEVEL_TEXT
} from '@/utils/riskCalculator'

const props = defineProps({
  isEditMode: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['submit'])

const formRef = ref()
const showStrategyPopup = ref(false)
const strategyStore = useStrategyStore()
const tradingStrategyStore = useTradingStrategyStore()

// 获取可用策略
const availableStrategies = computed(() => {
  return strategyStore.strategies
})

// 获取可用交易策略
const availableTradingStrategies = computed(() => {
  return tradingStrategyStore.tradingStrategies
})

// 股票store
const stockStore = useStockStore()

// 获取可用股票
const availableStocks = computed(() => {
  const stocks = stockStore.stocks
  return stocks ? stocks.filter(stock => stock.enabled) : []
})

// 表单数据
const formData = reactive({
  name: '',
  type: 'buy',
  stockCode: '',
  stockName: '',
  targetPrice: null,
  quantity: null,
  stopLoss: null,
  takeProfit: null,
  startTime: '',
  endTime: '',
  riskLevel: 'medium',
  description: '',
  remark: '',
  strategy: '',
  tradingStrategy: ''
})

// 风险信息
const riskInfo = ref({
  level: null,
  text: '',
  color: '',
  score: 0,
  details: null
})

// 风险建议
const riskSuggestions = computed(() => {
  return riskInfo.value.level ? getRiskSuggestions(riskInfo.value.level) : null
})

// 风险参数验证
const riskValidation = ref({
  valid: true,
  warnings: []
})

// 表单验证规则
const formRules = {
  name: [
    { required: true, message: '请输入计划名称', trigger: 'blur' }
  ],
  stockCode: [
    { required: true, message: '请选择股票', trigger: 'change' }
  ],
  strategy: [
    { required: true, message: '请选择选股策略', trigger: 'change' }
  ],
  type: [
    { required: true, message: '请选择交易方向', trigger: 'change' }
  ],
  targetPrice: [
    { required: true, message: '请输入目标价格', trigger: 'blur' }
  ],
  quantity: [
    { required: true, message: '请输入数量', trigger: 'blur' }
  ],
  stopLoss: [
    { required: true, message: '请输入止损价格', trigger: 'blur' }
  ],
  tradingStrategy: [
    { required: true, message: '请选择交易策略', trigger: 'change' }
  ]
}

const handleStockChange = (value) => {
  const selectedStock = availableStocks.value.find(v=>v.code===value)
  if (selectedStock) {
    formData.stockName = selectedStock.name
    generatePlanName()
  } else {
    formData.stockName = ''
  }
}

const generatePlanName = () => {
  if (formData.stockName && formData.strategy && formData.type) {
    const typeText = formData.type === 'buy' ? '买多' : '买空'
    const strategyText = availableStrategies.value.find(v=>v.id===formData.strategy)?.name || formData.strategy

    formData.name = `${formData.stockName} ${strategyText} ${typeText}计划`
  }
}

// 处理选股策略变化
const handleStrategyChange = (value) => {
  generatePlanName()
  updateRiskLevel()
}

// 更新风险等级
const updateRiskLevel = () => {
  if (formData.strategy && formData.tradingStrategy) {
    const risk = calculateRiskLevel(formData.strategy, formData.tradingStrategy)
    riskInfo.value = risk
    formData.riskLevel = risk.level

    // 验证当前价格参数
    validateRiskParameters()
  } else {
    riskInfo.value = {
      level: null,
      text: '',
      color: '',
      score: 0,
      details: null
    }
  }
}

// 验证风险参数
const validateRiskParameters = () => {
  if (formData.targetPrice && formData.stopLoss && formData.takeProfit && riskInfo.value.level) {
    const validation = validateRisk(
      formData.targetPrice,
      formData.stopLoss,
      formData.takeProfit,
      riskInfo.value.level
    )
    riskValidation.value = validation
  } else {
    riskValidation.value = {
      valid: true,
      warnings: []
    }
  }
}

// 获取风险等级文本
const getRiskLevelText = (riskLevel) => {
  return RISK_LEVEL_TEXT[riskLevel] || '未知'
}

// 暴露方法给 Service 组件
const getFormData = () => {
  // 确保风险等级被包含在表单数据中
  const data = { ...formData }

  // 如果有计算出的风险等级，使用计算结果
  if (riskInfo.value.level) {
    data.riskLevel = riskInfo.value.level
  }

  console.log('提交的表单数据:', data)
  return data
}

const setFormData = (data) => {
  Object.assign(formData, data)

  // 如果设置了策略数据，重新计算风险等级
  if (data.strategy && data.tradingStrategy) {
    // 使用 nextTick 确保数据已经更新
    setTimeout(() => {
      updateRiskLevel()
    }, 100)
  } else if (data.riskLevel) {
    // 如果直接设置了风险等级，显示它
    const riskLevelTexts = {
      'low': { level: 'low', text: '低风险', color: '#52c41a' },
      'medium': { level: 'medium', text: '中风险', color: '#faad14' },
      'high': { level: 'high', text: '高风险', color: '#ff4d4f' }
    }
    const riskData = riskLevelTexts[data.riskLevel]
    if (riskData) {
      riskInfo.value = {
        ...riskData,
        score: data.riskLevel === 'low' ? 1 : data.riskLevel === 'medium' ? 2 : 3,
        details: {
          explanation: '从已保存的计划中加载的风险等级'
        }
      }
    }
  }
}

const validate = () => {
  return formRef.value.validate()
}

// 监听策略变化，自动更新风险等级
watch([() => formData.strategy, () => formData.tradingStrategy], () => {
  updateRiskLevel()
}, { immediate: false })

// 组件挂载时不自动加载数据，避免无限递归
onMounted(() => {
  // 股票数据由父组件或用户手动触发加载
  console.log('TradingPlanForm mounted - 股票数据需要手动加载')
})

defineExpose({
  getFormData,
  setFormData,
  validate
})
</script>

<style scoped>
.trading-plan-form {
  padding: 1px 0;
  padding-top: 0;
}

.form-input-md {
  width: 300px;
}

.form-input-lg {
  width: 100%;
}

.strategy-field {
  display: flex;
  align-items: center;
  gap: 10px;
}

.strategy-popup-content {
  max-width: 350px;
}

.strategy-popup-content h4 {
  margin: 0 0 8px 0;
  font-size: 14px;
  font-weight: 600;
  color: var(--td-text-color-primary);
  border-bottom: 1px solid var(--td-brand-color);
  padding-bottom: 4px;
}

.strategy-popup-content ul {
  margin: 0;
  padding: 0;
  list-style: none;
}

.strategy-popup-content li {
  margin-bottom: 1px;
  padding: 4px 8px;
  background: var(--td-bg-color-container-hover);
  border-radius: 3px;
  font-size: 12px;
  line-height: 1.3;
}

.strategy-popup-content li strong {
  color: var(--td-brand-color);
  font-weight: 600;
}

/* 风险等级显示样式 */
.risk-level-display {
  display: flex;
  align-items: center;
  gap: 12px;
}

.risk-score {
  font-size: 12px;
  color: var(--td-text-color-secondary);
  font-weight: 500;
}

.risk-details {
  max-width: 380px;
}

.risk-details h4 {
  margin: 0 0 12px 0;
  font-size: 14px;
  font-weight: 600;
  color: var(--td-text-color-primary);
  border-bottom: 1px solid var(--td-brand-color);
  padding-bottom: 6px;
}

.risk-details h5 {
  margin: 12px 0 6px 0;
  font-size: 13px;
  font-weight: 600;
  color: var(--td-brand-color);
}

.risk-details p {
  margin: 6px 0;
  font-size: 12px;
  line-height: 1.4;
  color: var(--td-text-color-primary);
}

.risk-details ul {
  margin: 6px 0;
  padding-left: 16px;
}

.risk-details li {
  margin-bottom: 4px;
  font-size: 12px;
  line-height: 1.3;
  color: var(--td-text-color-secondary);
}

.risk-suggestions {
  margin-top: 12px;
  padding-top: 12px;
  border-top: 1px solid var(--td-border-level-1-color);
}
</style>
