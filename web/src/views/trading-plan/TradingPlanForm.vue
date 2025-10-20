<template>
  <div class="create-trading-plan">
    <t-card title="新建交易计划" class="create-card">
      <template #header>
        <div class="card-header">
          <span>创建新的交易计划</span>
          <t-button theme="default" @click="goBack" class="back-btn">
            <template #icon>
              <t-icon name="arrow-left" />
            </template>
            返回
          </t-button>
        </div>
      </template>
      
      <div class="create-content">
        <t-form
          ref="formRef"
          :model="formData"
          :rules="formRules"
          label-width="120px"
          @submit="handleSubmit"
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
                  @click="showStrategyPopup = !showStrategyPopup"
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
              maxlength="50"
              show-word-limit
            />
          </t-form-item>
          
          <t-form-item label="交易方向" name="type">
            <t-radio-group v-model="formData.type" @change="generatePlanName">
              <t-radio value="buy">买多</t-radio>
              <t-radio value="sell">买空</t-radio>
            </t-radio-group>
          </t-form-item>
          
          <FlexRow :gap="16">
            <t-form-item label="计划买进价格" name="targetPrice">
              <t-input-number
                class="form-input-md"
                v-model="formData.targetPrice"
                placeholder="请输入计划买进价格"
                :min="0"
                :precision="2"
              />
            </t-form-item>
            <t-form-item label="计划数量" name="quantity">
              <t-input-number
                class="form-input-md"
                v-model="formData.quantity"
                placeholder="请输入计划数量"
                :min="1"
              />
            </t-form-item>
          </FlexRow>
          
          <FlexRow :gap="16">
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
          
          
          <t-form-item label="风险等级" name="riskLevel">
            <t-radio-group v-model="formData.riskLevel">
              <t-radio value="low">低风险</t-radio>
              <t-radio value="medium">中风险</t-radio>
              <t-radio value="high">高风险</t-radio>
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
          
          <t-form-item label="计划原因描述" name="description">
            <t-textarea
              v-model="formData.description"
              placeholder="请输入计划原因描述"
              :maxlength="500"
              show-word-limit
            />
          </t-form-item>
          
          <t-form-item label="备注" name="remark">
            <t-textarea
              v-model="formData.remark"
              placeholder="请输入备注信息"
              :maxlength="200"
              show-word-limit
            />
          </t-form-item>
          
          <t-form-item label="">
            <t-space size="small">
              <t-button theme="default" @click="resetForm">
                <template #icon>
                  <t-icon name="refresh" />
                </template>
                重置
              </t-button>
              <t-button theme="default" @click="saveDraft">
                <template #icon>
                  <t-icon name="save" />
                </template>
                保存草稿
              </t-button>
              <t-button theme="primary" type="submit" :loading="isSubmitting">
                <template #icon>
                  <t-icon name="check" />
                </template>
                创建计划
              </t-button>
            </t-space>
          </t-form-item>
        </t-form>
      </div>
    </t-card>
    
    <!-- 预览对话框 -->
    <t-dialog
      v-model:visible="showPreviewDialog"
      header="计划预览"
      width="600px"
      :confirm-btn="null"
      :cancel-btn="null"
    >
      <div v-if="formData" class="plan-preview">
        <t-descriptions :data="getPreviewData()" :column="2" />
      </div>
      
      <template #footer>
        <t-space size="small">
          <t-button @click="showPreviewDialog = false">关闭</t-button>
          <t-button theme="primary" @click="confirmCreate">
            确认创建
          </t-button>
        </t-space>
      </template>
    </t-dialog>
    
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { useTradingPlanStore } from './store'
import { MessagePlugin } from 'tdesign-vue-next'
import { 
  getEnabledStrategies, 
  getStrategyById, 
  getStrategyDetailContent 
} from '@/utils/strategyConfig'
import { 
  getEnabledTradingStrategies, 
  getTradingStrategyById, 
  getTradingStrategyDetailContent 
} from '@/utils/tradingStrategyConfig'
import { 
  getEnabledStocks, 
  getStockByCode, 
  getStockInfo,
  getMarketText,
  getCurrencyText
} from '@/utils/stockConfig'

const router = useRouter()
const tradingPlanStore = useTradingPlanStore()

const formRef = ref()
const isSubmitting = ref(false)
const showPreviewDialog = ref(false)
const showStrategyPopup = ref(false)

// 获取可用策略
const availableStrategies = computed(() => {
  return getEnabledStrategies()
})

// 获取可用交易策略
const availableTradingStrategies = computed(() => {
  return getEnabledTradingStrategies()
})

// 获取可用股票
const availableStocks = computed(() => {
  return getEnabledStocks()
})

const formData = reactive({
  name: '',
  type: 'buy',
  stockCode: '',
  stockName: '',
  targetPrice: null,
  quantity: null,
  stopLoss: null,
  takeProfit: null,
  startTime: null,
  endTime: null,
  strategy: '',
  tradingStrategy: '',
  riskLevel: 'medium',
  description: '',
  remark: ''
})

const formRules = {
  name: [
    { required: true, message: '请输入计划名称', type: 'error' }
  ],
  type: [
    { required: true, message: '请选择交易方向', type: 'error' }
  ],
  stockCode: [
    { required: true, message: '请输入股票代码', type: 'error' },
    { pattern: /^\d{6}$/, message: '股票代码必须为6位数字', type: 'error' }
  ],
  stockName: [
    { required: true, message: '请输入股票名称', type: 'error' }
  ],
  targetPrice: [
    { required: true, message: '请输入计划买进价格', type: 'error' },
    { type: 'number', min: 0, message: '计划买进价格必须大于0', type: 'error' }
  ],
  quantity: [
    { required: true, message: '请输入计划数量', type: 'error' },
    { type: 'number', min: 1, message: '计划数量必须大于0', type: 'error' }
  ],
  strategy: [
    { required: true, message: '请选择选股策略', type: 'error' }
  ],
  tradingStrategy: [
    { required: true, message: '请选择交易策略', type: 'error' }
  ],
  stopLoss: [
    { required: true, message: '请输入止损价格', type: 'error' },
    { type: 'number', min: 0, message: '止损价格必须大于等于0', type: 'error' }
  ]
}

// 模拟股票数据
const stockDatabase = {
  '000001': '平安银行',
  '000002': '万科A',
  '600036': '招商银行',
  '600519': '贵州茅台',
  '000858': '五粮液',
  '002415': '海康威视',
  '300059': '东方财富',
  '000725': '京东方A'
}

const handleStockChange = (value) => {
  if (value && stockDatabase[value]) {
    formData.stockName = stockDatabase[value]
    generatePlanName()
  } else {
    formData.stockName = ''
  }
}

// 生成计划名称
const generatePlanName = () => {
  if (formData.stockName && formData.strategy) {
    const strategy = getStrategyById(formData.strategy)
    const strategyName = strategy ? strategy.name : '选股策略'
    const direction = formData.type === 'buy' ? '买多' : '买空'
    
    formData.name = `${formData.stockName}${strategyName}${direction}计划`
  }
}

const handleStockCodeBlur = () => {
  if (formData.stockCode && stockDatabase[formData.stockCode]) {
    formData.stockName = stockDatabase[formData.stockCode]
  } else if (formData.stockCode) {
    formData.stockName = ''
    MessagePlugin.warning('未找到该股票代码对应的股票名称')
  }
}

const getPreviewData = () => [
  { label: '计划名称', content: formData.name },
  { label: '交易方向', content: formData.type === 'buy' ? '买多' : '买空' },
  { label: '股票代码', content: formData.stockCode },
  { label: '股票名称', content: formData.stockName },
  { label: '计划买进价格', content: `¥${formData.targetPrice}` },
  { label: '计划数量', content: formData.quantity },
  { label: '止损价格', content: formData.stopLoss ? `¥${formData.stopLoss}` : '未设置' },
  { label: '止盈价格', content: formData.takeProfit ? `¥${formData.takeProfit}` : '未设置' },
  { label: '选股策略', content: getStrategyText(formData.strategy) },
  { label: '交易策略', content: getTradingStrategyText(formData.tradingStrategy) },
  { label: '风险等级', content: getRiskLevelText(formData.riskLevel) },
  { label: '计划原因描述', content: formData.description || '无' },
  { label: '备注', content: formData.remark || '无' }
]

const getStrategyText = (strategy) => {
  const strategies = {
    // 技术分析策略
    trend_following: '趋势跟踪',
    breakout: '突破策略',
    reversal: '反转策略',
    bollinger_reversal: '布林带反转',
    momentum: '动量策略',
    
    // 其他推荐策略（注释中保留）
    // 基本面分析策略：
    // value_investing: '价值投资',
    // growth_investing: '成长投资',
    // dividend_strategy: '分红策略',
    // sector_rotation: '行业轮动',
    
    // 量化选股策略：
    // multi_factor: '多因子模型',
    // technical_indicators: '技术指标',
    // capital_flow: '资金流向',
    // sentiment_indicators: '情绪指标'
  }
  return strategies[strategy] || strategy
}

const getTradingStrategyText = (strategy) => {
  const strategyDetail = getTradingStrategyById(strategy)
  return strategyDetail ? strategyDetail.name : strategy
}


const getRiskLevelText = (level) => {
  const levels = {
    low: '低风险',
    medium: '中风险',
    high: '高风险'
  }
  return levels[level] || level
}

const handleSubmit = async () => {
  const result = await formRef.value.validate()
  if (result === true) {
    showPreviewDialog.value = true
  }
}


const confirmCreate = async () => {
  isSubmitting.value = true
  
  try {
    const planData = {
      ...formData,
      id: Date.now().toString(),
      status: 'active',
      createTime: new Date().toLocaleString(),
      updateTime: new Date().toLocaleString()
    }
    
    tradingPlanStore.addPlan(planData)
    
    MessagePlugin.success('交易计划创建成功！')
    showPreviewDialog.value = false
    router.push('/trading-plan')
  } catch (error) {
    MessagePlugin.error('创建失败，请重试')
  } finally {
    isSubmitting.value = false
  }
}

const resetForm = () => {
  formRef.value.reset()
  Object.assign(formData, {
    name: '',
    type: 'buy',
    stockCode: '',
    stockName: '',
    targetPrice: null,
    quantity: null,
    stopLoss: null,
    takeProfit: null,
    startTime: null,
    endTime: null,
    strategy: '',
    tradingStrategy: '',
    riskLevel: 'medium',
    description: '',
    remark: ''
  })
}

const saveDraft = () => {
  // 保存草稿逻辑
  MessagePlugin.success('草稿已保存')
}

const goBack = () => {
  router.back()
}
</script>

<style scoped>
.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
}

.strategy-field {
  display: flex;
  align-items: center;
  gap: 8px;
}

.strategy-field .form-input-md {
  flex: 1;
}

.strategy-popup-content {
  max-width: 280px;
  max-height: 200px;
  /* padding: 8px 12px; */
  background: white;
  border-radius: 6px;
  /* box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1); */
  overflow-y: auto;
}

.strategy-popup-content h4 {
  margin: 0 0 8px 0;
  color: #0369a1;
  font-size: 13px;
  font-weight: 600;
  border-bottom: 1px solid #bae6fd;
  padding-bottom: 4px;
}

.strategy-popup-content ul {
  margin: 0;
  padding-left: 14px;
  color: #374151;
  font-size: 11px;
  line-height: 1.4;
}

.strategy-popup-content li {
  margin-bottom: 3px;
}

.strategy-popup-content strong {
  color: #0369a1;
  font-weight: 600;
}




/* 表单项样式由全局form.css提供 */
</style>
