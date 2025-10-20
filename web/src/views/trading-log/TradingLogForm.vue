<template>
  <div class="create-trading-log">
    <t-card :title="isEditMode ? '编辑交易日志' : '新建交易日志'" class="create-card" :bordered="false">
      <template #header>
        <div class="card-header">
          <span>{{ isEditMode ? '编辑交易日志' : '创建新的交易日志' }}</span>
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
          <t-form-item label="执行交易计划" name="planName">
            <t-select class="form-input-md" v-model="formData.planName" placeholder="请选择关联的交易计划" clearable style="width: 300px;" @change="handlePlanChange">
              <t-option value="" label="无关联计划" />
              <t-option value="平安银行买多计划" label="平安银行买多计划" />
              <t-option value="贵州茅台买空计划" label="贵州茅台买空计划" />
              <t-option value="招商银行买多计划" label="招商银行买多计划" />
            </t-select>
          </t-form-item>
          
          <t-form-item label="日志标题" name="title">
            <t-input
              class="form-input-md"
              v-model="formData.title"
              placeholder="请输入日志标题"
              maxlength="50"
              show-word-limit
            />
          </t-form-item>
          
          <t-form-item label="交易方向" name="type">
            <t-radio-group v-model="formData.type">
              <t-radio value="buy">买多</t-radio>
              <t-radio value="sell">买空</t-radio>
            </t-radio-group>
          </t-form-item>
          
          <t-form-item label="选择股票" name="stockCode">
            <t-select 
              class="form-input-md" 
              v-model="formData.stockCode" 
              placeholder="请选择股票"
              @change="handleStockChange"
              :disabled="!!formData.planName"
            >
              <t-option value="000001" label="平安银行 (000001)" />
              <t-option value="000002" label="万科A (000002)" />
              <t-option value="600036" label="招商银行 (600036)" />
              <t-option value="600519" label="贵州茅台 (600519)" />
              <t-option value="000858" label="五粮液 (000858)" />
              <t-option value="002415" label="海康威视 (002415)" />
              <t-option value="300059" label="东方财富 (300059)" />
              <t-option value="000725" label="京东方A (000725)" />
            </t-select>
            <div v-if="formData.planName" class="plan-stock-tip">
              <t-icon name="info-circle" />
              <span>股票已根据选择的交易计划自动关联</span>
            </div>
          </t-form-item>
          
          <FlexRow :gap="16">
            <t-form-item label="成交价格" name="price">
              <t-input-number
                class="form-input-md"
                v-model="formData.price"
                placeholder="请输入成交价格"
                :min="0"
                :precision="2"
              />
            </t-form-item>
            <t-form-item label="成交数量" name="quantity">
              <t-input-number
                class="form-input-md"
                v-model="formData.quantity"
                placeholder="请输入成交数量"
                :min="1"
              />
            </t-form-item>
          </FlexRow>
          
          <t-form-item label="交易状态" name="status">
            <t-select class="form-input-md" v-model="formData.status" placeholder="请选择交易状态">
              <t-option value="pending" label="进行中" />
              <t-option value="completed" label="已经结束" />
            </t-select>
          </t-form-item>
          
          <t-form-item label="交易时间" name="tradingTime">
            <t-date-picker
              v-model="formData.tradingTime"
              placeholder="请选择交易时间"
              :enable-time-picker="true"
              format="YYYY-MM-DD HH:mm:ss"
              value-format="YYYY-MM-DD HH:mm:ss"
            />
          </t-form-item>
          
          <t-form-item label="时区设置" name="timezone">
            <t-radio-group v-model="formData.timezone">
              <t-radio value="local">中国时间 (UTC+8)</t-radio>
              <t-radio value="us">美股时间 (美东时间)</t-radio>
            </t-radio-group>
            <div class="timezone-tip">
              <t-icon name="info-circle" />
              <span v-if="formData.timezone === 'local'">
                按中国时间记录，便于本地管理
              </span>
              <span v-else-if="formData.timezone === 'us'">
                按美股市场时间记录，便于技术分析
              </span>
            </div>
          </t-form-item>
          
          <t-form-item label="交易费用" name="fee">
            <t-input-number
              class="form-input-md"
              v-model="formData.fee"
              placeholder="请输入交易费用"
              :min="0"
              :precision="2"
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
          
          <t-form-item label="风险等级" name="riskLevel">
            <t-radio-group v-model="formData.riskLevel">
              <t-radio value="low">低风险</t-radio>
              <t-radio value="medium">中风险</t-radio>
              <t-radio value="high">高风险</t-radio>
            </t-radio-group>
          </t-form-item>
          
          <t-form-item label="日志内容" name="content">
            <t-textarea
              v-model="formData.content"
              placeholder="请详细描述交易过程、分析思路、操作细节等"
              :maxlength="1000"
              :autosize="{ minRows: 4, maxRows: 8 }"
              show-word-limit
            />
          </t-form-item>
          
          <t-form-item label="备注" name="remark">
            <t-textarea
              v-model="formData.remark"
              placeholder="请输入备注信息"
              :maxlength="200"
              :autosize="{ minRows: 2, maxRows: 4 }"
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
                创建日志
              </t-button>
            </t-space>
          </t-form-item>
        </t-form>
      </div>
    </t-card>
    
    <!-- 预览对话框 -->
    <t-dialog
      v-model:visible="showPreviewDialog"
      header="日志预览"
      width="600px"
      :confirm-btn="null"
      :cancel-btn="null"
    >
      <div v-if="formData" class="log-preview">
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
import { ref, reactive, computed, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useTradingLogStore } from './store'
import { MessagePlugin } from 'tdesign-vue-next'
import { 
  getEnabledTradingStrategies, 
  getTradingStrategyById 
} from '@/utils/tradingStrategyConfig'

const router = useRouter()
const route = useRoute()
const tradingLogStore = useTradingLogStore()

const formRef = ref()
const isSubmitting = ref(false)
const showPreviewDialog = ref(false)

// 判断是否为编辑模式
const isEditMode = computed(() => !!route.params.id)
const logId = computed(() => route.params.id)

// 获取可用交易策略
const availableTradingStrategies = computed(() => {
  return getEnabledTradingStrategies()
})

const formData = reactive({
  title: '',
  type: 'buy',
  stockCode: '',
  stockName: '',
  price: null,
  quantity: null,
  profit: null,
  status: 'pending',
  tradingTime: null,
  timezone: 'local', // 默认使用中国时间
  fee: null,
  planName: '',
  strategy: '',
  riskLevel: 'medium',
  content: '',
  remark: ''
})

const formRules = {
  planName: [
    { required: true, message: '请选择执行交易计划', type: 'error' }
  ],
  title: [
    { required: true, message: '请输入日志标题', type: 'error' }
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
  price: [
    { required: true, message: '请输入成交价格', type: 'error' },
    { type: 'number', min: 0, message: '成交价格必须大于0', type: 'error' }
  ],
  quantity: [
    { required: true, message: '请输入成交数量', type: 'error' },
    { type: 'number', min: 1, message: '成交数量必须大于0', type: 'error' }
  ],
  status: [
    { required: true, message: '请选择交易状态', type: 'error' }
  ],
  tradingTime: [
    { required: true, message: '请选择交易时间', type: 'error' }
  ],
  strategy: [
    { required: true, message: '请选择交易策略', type: 'error' }
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
  } else {
    formData.stockName = ''
  }
}

// 计划与股票的映射关系
const planStockMapping = {
  '平安银行买多计划': { code: '000001', name: '平安银行' },
  '贵州茅台买空计划': { code: '600519', name: '贵州茅台' },
  '招商银行买多计划': { code: '600036', name: '招商银行' }
}

// 处理计划选择变化
const handlePlanChange = (value) => {
  if (value && planStockMapping[value]) {
    const stockInfo = planStockMapping[value]
    formData.stockCode = stockInfo.code
    formData.stockName = stockInfo.name
    
    // 根据计划名称推断交易方向
    if (value.includes('买多')) {
      formData.type = 'buy'
    } else if (value.includes('买空')) {
      formData.type = 'sell'
    }
    
    // 自动生成日志标题
    formData.title = `执行${value}`
  } else {
    // 清空关联的股票信息
    formData.stockCode = ''
    formData.stockName = ''
    formData.title = ''
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
  { label: '日志标题', content: formData.title },
  { label: '交易方向', content: formData.type === 'buy' ? '买多' : '买空' },
  { label: '股票代码', content: formData.stockCode },
  { label: '股票名称', content: formData.stockName },
  { label: '成交价格', content: `¥${formData.price}` },
  { label: '成交数量', content: formData.quantity },
  { label: '盈亏金额', content: `¥${formData.profit}` },
  { label: '交易状态', content: getStatusText(formData.status) },
  { label: '执行交易计划', content: formData.planName || '无关联计划' },
    { label: '交易时间', content: formData.tradingTime ? new Date(formData.tradingTime).toLocaleString() : '未设置' },
    { label: '时区设置', content: formData.timezone === 'local' ? '中国时间 (UTC+8)' : '美股时间 (美东时间)' },
  { label: '交易费用', content: formData.fee ? `¥${formData.fee}` : '未设置' },
  { label: '交易策略', content: getStrategyText(formData.strategy) },
  { label: '风险等级', content: getRiskLevelText(formData.riskLevel) },
  { label: '日志内容', content: formData.content || '无' },
  { label: '备注', content: formData.remark || '无' }
]

const getStatusText = (status) => {
  const statuses = {
    completed: '已经结束',
    pending: '进行中'
  }
  return statuses[status] || status
}

const getStrategyText = (strategy) => {
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
    if (isEditMode.value) {
      // 编辑模式：更新现有日志
      const logData = {
        ...formData,
        id: logId.value,
        updateTime: new Date().toLocaleString()
      }
      
      tradingLogStore.updateLog(logId.value, logData)
      MessagePlugin.success('交易日志更新成功！')
    } else {
      // 新建模式：创建新日志
      const logData = {
        ...formData,
        id: Date.now().toString(),
        createTime: new Date().toLocaleString(),
        updateTime: new Date().toLocaleString()
      }
      
      tradingLogStore.addLog(logData)
      MessagePlugin.success('交易日志创建成功！')
    }
    
    showPreviewDialog.value = false
    router.push('/trading-log')
  } catch (error) {
    MessagePlugin.error('保存失败，请重试')
  } finally {
    isSubmitting.value = false
  }
}

const resetForm = () => {
  formRef.value.reset()
  Object.assign(formData, {
    title: '',
    type: 'buy',
    stockCode: '',
    stockName: '',
    price: null,
    quantity: null,
    profit: null,
    status: 'pending',
    tradingTime: null,
    timezone: 'local',
    fee: null,
    planName: '',
    strategy: '',
    riskLevel: 'medium',
    content: '',
    remark: ''
  })
}

const saveDraft = () => {
  // 保存草稿逻辑
  MessagePlugin.success('草稿已保存')
}

// 加载编辑数据
const loadEditData = async () => {
  if (isEditMode.value && logId.value) {
    try {
      const logs = tradingLogStore.getLogs
      const log = logs.find(l => l.id === logId.value)
      if (log) {
        Object.assign(formData, log)
      } else {
        MessagePlugin.error('未找到要编辑的交易日志')
        goBack()
      }
    } catch (error) {
      console.error('加载交易日志失败:', error)
      MessagePlugin.error('加载交易日志失败')
    }
  }
}

// 组件挂载时加载数据
onMounted(() => {
  loadEditData()
})

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

.timezone-tip {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 8px;
  padding: 8px 12px;
  background: #f0f9ff;
  border: 1px solid #bae6fd;
  border-radius: 4px;
  font-size: 12px;
  color: #0369a1;
}

.plan-stock-tip {
  display: flex;
  align-items: center;
  gap: 8px;
  margin-top: 8px;
  padding: 8px 12px;
  background: #f0fdf4;
  border: 1px solid #bbf7d0;
  border-radius: 4px;
  font-size: 12px;
  color: #166534;
}
</style>
