<template>
  <div class="trading-plan-form">
    <t-form
      ref="formRef"
      :data="formData"
      :rules="formRules"
      label-width="120px"
      scroll-to-first-error="smooth"
      @submit="onFormSubmit"
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
            :label="`${stock.name} (${stock.code}) - ${getMarketText(stock.marketType)}`" 
          />
        </t-select>
      </t-form-item>
      
      <t-form-item label="选股策略" name="strategy">
        <div class="strategy-field">
          <t-select class="form-input-md" v-model="formData.strategy" placeholder="请选择选股策略" @change="generatePlanName">
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
        <t-select class="form-input-md" v-model="formData.tradingStrategy" placeholder="请选择交易策略">
          <t-option 
            v-for="strategy in availableTradingStrategies" 
            :key="strategy.id" 
            :value="strategy.id" 
            :label="`${strategy.name} - ${strategy.description}`" 
          />
        </t-select>
      </t-form-item>
      <FlexRow>
        <t-form-item label="计划买进价格" name="targetPrice">
        <t-input-number 
          class="form-input-md" 
          v-model="formData.targetPrice" 
          placeholder="请输入目标价格"
          :min="0"
          :precision="2"
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
        />
      </t-form-item>
      
      <t-form-item label="止盈价格" name="takeProfit">
        <t-input-number 
          class="form-input-md" 
          v-model="formData.takeProfit" 
          placeholder="请输入止盈价格"
          :min="0"
          :precision="2"
        />
      </t-form-item>
      </FlexRow>
      
      
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
import { ref, reactive, computed, onMounted } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'
import { 
  getEnabledStrategies, 
  getStrategyById
} from '@/utils/strategyConfig'
import { 
  getEnabledTradingStrategies
} from '@/utils/tradingStrategyConfig'
import { useStockStore } from '@/stores/stock'
import { getMarketText } from '@/utils/stockConfig'

const props = defineProps({
  isEditMode: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['submit'])

const formRef = ref()
const showStrategyPopup = ref(false)

// 获取可用策略
const availableStrategies = computed(() => {
  return getEnabledStrategies()
})

// 获取可用交易策略
const availableTradingStrategies = computed(() => {
  return getEnabledTradingStrategies()
})

// 股票store
const stockStore = useStockStore()

// 获取可用股票
const availableStocks = computed(() => {
  return stockStore.getStockList.filter(stock => stock.enabled)
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
  const selectedStock = stockStore.getStockByCode(value)
  if (selectedStock) {
    formData.stockName = selectedStock.name
    generatePlanName()
  } else {
    formData.stockName = ''
  }
}

const generatePlanName = () => {
  if (formData.stockName && formData.strategy && formData.type) {
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
    
    const typeText = formData.type === 'buy' ? '买多' : '买空'
    const strategyText = strategyNames[formData.strategy] || formData.strategy
    
    formData.name = `${formData.stockName} ${strategyText} ${typeText}计划`
  }
}

const onFormSubmit = ({ validateResult, firstError, e }) => {
  e?.preventDefault?.()
  if (validateResult === true) {
    // 触发 submit 事件，让 Service 组件处理
    emit('submit')
  } else {
    MessagePlugin.warning(firstError)
  }
}

// 暴露方法给 Service 组件
const getFormData = () => {
  return { ...formData }
}

const setFormData = (data) => {
  Object.assign(formData, data)
}

const validate = () => {
  return formRef.value.validate()
}

// 组件挂载时加载股票数据
onMounted(async () => {
  try {
    await stockStore.loadStocks()
  } catch (error) {
    console.error('加载股票数据失败:', error)
  }
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
</style>
