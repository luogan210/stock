<template>
  <div class="trading-log-form">
    <t-form
      ref="formRef"
      :data="formData"
      :rules="formRules"
      label-width="120px" 
      scroll-to-first-error="smooth"
    >
      <t-form-item label="日志名称" name="logName">
        <t-input
          class="form-input-md"
          v-model="formData.logName"
          placeholder="请输入日志名称"
        />
      </t-form-item>

      <t-form-item label="执行交易计划" name="planName">
        <t-select
          class="form-input-md"
          v-model="formData.planName"
          placeholder="请选择要执行的交易计划"
          @change="handlePlanChange"
        >
          <t-option
            v-for="plan in availablePlans"
            :key="plan.id"
            :value="plan.name"
            :label="plan.name"
          />
        </t-select>
      </t-form-item>

      <t-form-item label="选择股票" name="stockCode">
        <t-select
          class="form-input-md"
          v-model="formData.stockCode"
          placeholder="请选择股票"
          :disabled="!!formData.planName"
          @change="handleStockChange"
        >
          <t-option
            v-for="stock in availableStocks"
            :key="stock.code"
            :value="stock.code"
            :label="`${stock.name} (${stock.code})`"
          />
        </t-select>
      </t-form-item>

      <t-form-item label="交易方向" name="type">
        <t-radio-group v-model="formData.type" :disabled="!!formData.planName">
          <t-radio value="buy">买多</t-radio>
          <t-radio value="sell">买空</t-radio>
        </t-radio-group>
      </t-form-item>

      <t-form-item label="交易时间" name="tradingTime">
        <t-date-picker
          class="form-input-md"
          v-model="formData.tradingTime"
          placeholder="选择交易时间"
          enable-time-picker
          format="YYYY-MM-DD HH:mm:ss"
          value-format="YYYY-MM-DD HH:mm:ss"
        />
      </t-form-item>

      <t-form-item label="交易价格" name="price">
        <t-input-number
          class="form-input-md"
          v-model="formData.price"
          placeholder="请输入交易价格"
          :min="0"
          :precision="2"
        />
      </t-form-item>

      <t-form-item label="交易数量" name="quantity">
        <t-input-number
          class="form-input-md"
          v-model="formData.quantity"
          placeholder="请输入交易数量"
          :min="1"
        />
      </t-form-item>

      <t-form-item label="交易策略" name="strategy">
        <t-select class="form-input-md" v-model="formData.strategy" placeholder="请选择交易策略">
          <t-option
            v-for="strategy in availableTradingStrategies"
            :key="strategy.id"
            :value="strategy.id"
            :label="`${strategy.name} - ${strategy.description}`"
          />
        </t-select>
      </t-form-item>

      <t-form-item label="交易状态" name="tradingStatus">
        <t-select class="form-input-md" v-model="formData.tradingStatus" placeholder="请选择交易状态">
          <t-option value="pending" label="待执行" />
          <t-option value="executing" label="执行中" />
          <t-option value="completed" label="已完成" />
          <t-option value="cancelled" label="已取消" />
          <t-option value="failed" label="执行失败" />
        </t-select>
      </t-form-item>

      <t-form-item label="备注" name="remark">
        <t-textarea
          class="form-input-lg"
          v-model="formData.remark"
          placeholder="交易备注信息"
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
  getEnabledTradingStrategies
} from '@/utils/tradingStrategyConfig'
import { useStockStore } from "@/stores/index.js";

const props = defineProps({
  isEditMode: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['submit'])

const formRef = ref()

// 获取可用交易策略
const availableTradingStrategies = computed(() => {
  return getEnabledTradingStrategies()
})

// 股票store
const stockStore = useStockStore()

// 获取可用股票
const availableStocks = computed(() => {
  const stocks = stockStore.stocks
  return stocks ? stocks.filter(stock => stock.enabled) : []
})

// 模拟可用计划
const availablePlans = computed(() => {
  return [
    { id: '1', name: '平安银行技术分析买多计划' },
    { id: '2', name: '贵州茅台价值投资买多计划' },
    { id: '3', name: '招商银行动量策略买空计划' }
  ]
})

// 表单数据
const formData = reactive({
  logName: '',
  planName: '',
  stockCode: '',
  stockName: '',
  type: 'buy',
  tradingTime: '',
  price: null,
  quantity: null,
  strategy: '',
  tradingStatus: 'pending',
  remark: ''
})

// 表单验证规则
const formRules = {
  logName: [
    { required: true, message: '请输入日志名称', trigger: 'blur' }
  ],
  stockCode: [
    { required: true, message: '请选择股票', trigger: 'change' }
  ],
  type: [
    { required: true, message: '请选择交易方向', trigger: 'change' }
  ],
  tradingTime: [
    { required: true, message: '请选择交易时间', trigger: 'change' }
  ],
  price: [
    { required: true, message: '请输入交易价格', trigger: 'blur' }
  ],
  quantity: [
    { required: true, message: '请输入交易数量', trigger: 'blur' }
  ],
  strategy: [
    { required: true, message: '请选择交易策略', trigger: 'change' }
  ],
  tradingStatus: [
    { required: true, message: '请选择交易状态', trigger: 'change' }
  ]
}

// 计划与股票的映射
const planStockMapping = {
  '平安银行技术分析买多计划': { code: '000001', name: '平安银行' },
  '贵州茅台价值投资买多计划': { code: '600519', name: '贵州茅台' },
  '招商银行动量策略买空计划': { code: '600036', name: '招商银行' }
}

const handlePlanChange = (value) => {
  const mapping = planStockMapping[value]
  if (mapping) {
    formData.stockCode = mapping.code
    formData.stockName = mapping.name
    formData.type = 'buy' // 默认买多
  }
}

const handleStockChange = (value) => {
  const selectedStock = stockStore.getStockByCode(value)
  if (selectedStock) {
    formData.stockName = selectedStock.name
  } else {
    formData.stockName = ''
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

// 组件挂载时不自动加载数据，避免undefined错误
onMounted(() => {
  // 股票数据由父组件或用户手动触发加载
  console.log('TradingLogForm mounted - 股票数据需要手动加载')
})

defineExpose({
  getFormData,
  setFormData,
  validate
})
</script>

<style scoped>
.trading-log-form {
  padding: 1px 0;
  padding-top: 0;
}

.form-input-md {
  width: 300px;
}

.form-input-lg {
  width: 100%;
}
</style>
