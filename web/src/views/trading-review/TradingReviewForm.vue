<template>
  <div class="trading-review-form">
    <t-card :title="isEditMode ? '编辑交易复盘' : '新建交易复盘'" class="form-card" :bordered="false">
      <template #header>
        <div class="card-header">
          <span>{{ isEditMode ? '编辑交易复盘' : '创建新的交易复盘' }}</span>
          <t-button theme="default" @click="goBack" class="back-btn">
            <template #icon>
              <t-icon name="arrow-left" />
            </template>
            返回
          </t-button>
        </div>
      </template>

      <div class="form-content">
        <t-form
          ref="formRef"
          :model="formData"
          :rules="formRules"
          label-width="120px"
          @submit="handleSubmit"
        >
          <!-- 基本信息 -->
          <t-card title="基本信息" class="section-card">
            <t-form-item label="复盘标题" name="title">
              <t-input
                v-model="formData.title"
                placeholder="请输入复盘标题"
                clearable
              />
            </t-form-item>
            
            <t-form-item label="复盘周期" name="period">
              <t-radio-group v-model="formData.period">
                <t-radio value="daily">日复盘</t-radio>
                <t-radio value="weekly">周复盘</t-radio>
                <t-radio value="monthly">月复盘</t-radio>
              </t-radio-group>
            </t-form-item>
            
            <t-form-item label="复盘日期" name="reviewDate">
              <!-- 日复盘使用普通日期选择器 -->
              <t-date-picker
                v-if="formData.period === 'daily'"
                v-model="formData.reviewDate"
                :placeholder="getDatePickerPlaceholder()"
                mode="date"
                format="YYYY-MM-DD"
                :enable-time-picker="false"
                clearable
              />
              <!-- 周复盘使用周选择器 -->
              <t-date-picker
                v-else-if="formData.period === 'weekly'"
                v-model="formData.reviewDate"
                :placeholder="getDatePickerPlaceholder()"
                mode="week"
                format="YYYY 第 WW 周"
                value-format="YYYY-MM-DD"
                clearable
                allow-input
                :first-day-of-week="1"
                @change="handleWeekChange"
              />
              <!-- 月复盘使用月份选择器 -->
              <t-date-picker
                v-else-if="formData.period === 'monthly'"
                v-model="formData.reviewDate"
                :placeholder="getDatePickerPlaceholder()"
                mode="month"
                format="YYYY-MM"
                :enable-time-picker="false"
                clearable
              />
            </t-form-item>
          </t-card>

          <!-- 交易统计 -->
          <t-card title="交易统计" class="section-card">
            <FlexRow wrap="wrap" gap="16">
              <t-form-item label="总交易次数" name="totalTrades" class="form-item-flex">
                <t-input-number
                  v-model="formData.totalTrades"
                  placeholder="请输入总交易次数"
                  :min="0"
                />
              </t-form-item>
              <t-form-item label="盈利交易" name="winTrades" class="form-item-flex">
                <t-input-number
                  v-model="formData.winTrades"
                  placeholder="请输入盈利交易次数"
                  :min="0"
                />
              </t-form-item>
              <t-form-item label="亏损交易" name="lossTrades" class="form-item-flex">
                <t-input-number
                  v-model="formData.lossTrades"
                  placeholder="请输入亏损交易次数"
                  :min="0"
                />
              </t-form-item>
              <t-form-item label="胜率(%)" name="winRate" class="form-item-flex">
                <t-input-number
                  v-model="formData.winRate"
                  placeholder="请输入胜率"
                  :min="0"
                  :max="100"
                />
              </t-form-item>
            </FlexRow>
          </t-card>

          <!-- 盈亏分析 -->
          <t-card title="盈亏分析" class="section-card">
            <FlexRow wrap="wrap" gap="16">
              <t-form-item label="总盈亏" name="totalProfit" class="form-item-flex">
                <t-input-number
                  v-model="formData.totalProfit"
                  placeholder="请输入总盈亏"
                  :precision="2"
                />
              </t-form-item>
              <t-form-item label="最大回撤" name="maxDrawdown" class="form-item-flex">
                <t-input-number
                  v-model="formData.maxDrawdown"
                  placeholder="请输入最大回撤"
                  :precision="2"
                />
              </t-form-item>
              <t-form-item label="平均盈利" name="avgProfit" class="form-item-flex">
                <t-input-number
                  v-model="formData.avgProfit"
                  placeholder="请输入平均盈利"
                  :precision="2"
                />
              </t-form-item>
              <t-form-item label="平均亏损" name="avgLoss" class="form-item-flex">
                <t-input-number
                  v-model="formData.avgLoss"
                  placeholder="请输入平均亏损"
                  :precision="2"
                />
              </t-form-item>
              <t-form-item label="盈利因子" name="profitFactor" class="form-item-flex">
                <t-input-number
                  v-model="formData.profitFactor"
                  placeholder="请输入盈利因子"
                  :precision="2"
                  :min="0"
                />
              </t-form-item>
              <t-form-item label="夏普比率" name="sharpeRatio" class="form-item-flex">
                <t-input-number
                  v-model="formData.sharpeRatio"
                  placeholder="请输入夏普比率"
                  :precision="2"
                />
              </t-form-item>
            </FlexRow>
          </t-card>

          <!-- 复盘内容 -->
          <t-card title="复盘内容" class="section-card">
            <t-form-item label="复盘总结" name="content">
              <t-textarea
                v-model="formData.content"
                placeholder="请详细描述本次复盘的总结内容"
                :autosize="{ minRows: 4, maxRows: 8 }"
              />
            </t-form-item>
            
            <t-form-item label="经验教训" name="lessons">
              <t-textarea
                v-model="lessonsText"
                placeholder="请输入经验教训，每行一条"
                :autosize="{ minRows: 3, maxRows: 6 }"
                @change="updateLessons"
              />
            </t-form-item>
            
            <t-form-item label="改进计划" name="improvements">
              <t-textarea
                v-model="improvementsText"
                placeholder="请输入改进计划，每行一条"
                :autosize="{ minRows: 3, maxRows: 6 }"
                @change="updateImprovements"
              />
            </t-form-item>
            
            <t-form-item :label="periodPlanLabel" name="nextPlan">
              <t-textarea
                v-model="formData.nextPlan"
                :placeholder="`请输入${periodPlanLabel}`"
                :autosize="{ minRows: 2, maxRows: 4 }"
              />
            </t-form-item>
          </t-card>

          <!-- 操作按钮 -->
          <div class="form-actions">
            <t-button theme="default" @click="goBack">
              取消
            </t-button>
            <t-button theme="default" @click="saveDraft">
              保存草稿
            </t-button>
            <t-button theme="primary" type="submit" :loading="isSubmitting">
              {{ isEditMode ? '更新复盘' : '创建复盘' }}
            </t-button>
          </div>
        </t-form>
      </div>
    </t-card>

    <!-- 预览弹窗 -->
    <t-dialog
      v-model:visible="showPreviewDialog"
      :header="`预览${isEditMode ? '编辑' : '新建'}的复盘`"
      width="800px"
      :footer="false"
    >
      <div class="preview-content">
        <t-descriptions :data="getPreviewData()" />
        
        <div class="preview-section">
          <h4>复盘总结</h4>
          <p>{{ formData.content || '暂无内容' }}</p>
        </div>
        
        <div class="preview-section">
          <h4>经验教训</h4>
          <ul>
            <li v-for="(lesson, index) in formData.lessons" :key="index">
              {{ lesson }}
            </li>
          </ul>
        </div>
        
        <div class="preview-section">
          <h4>改进计划</h4>
          <ul>
            <li v-for="(improvement, index) in formData.improvements" :key="index">
              {{ improvement }}
            </li>
          </ul>
        </div>
        
        <div class="preview-section">
          <h4>{{ periodPlanLabel }}</h4>
          <p>{{ formData.nextPlan || '暂无内容' }}</p>
        </div>
      </div>
      
      <div class="preview-actions">
        <t-button theme="default" @click="showPreviewDialog = false">
          返回编辑
        </t-button>
        <t-button theme="primary" @click="confirmCreate" :loading="isSubmitting">
          确认{{ isEditMode ? '更新' : '创建' }}
        </t-button>
      </div>
    </t-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, computed, onMounted, watch } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useTradingReviewStore } from './store'
import { MessagePlugin } from 'tdesign-vue-next'
import FlexRow from '@/components/FlexRow.vue'

const router = useRouter()
const route = useRoute()
const tradingReviewStore = useTradingReviewStore()

// 响应式数据
const formRef = ref()
const isSubmitting = ref(false)
const showPreviewDialog = ref(false)

// 判断是否为编辑模式
const isEditMode = computed(() => !!route.params.id)
const reviewId = computed(() => route.params.id)

// 表单数据
const formData = reactive({
  title: '',
  period: 'daily',
  reviewDate: null,
  totalTrades: null,
  winTrades: null,
  lossTrades: null,
  winRate: null,
  totalProfit: null,
  maxLoss: null,
  avgProfit: null,
  avgLoss: null,
  content: '',
  lessons: [],
  improvements: [],
  nextPlan: ''
})

// 文本输入（用于处理数组）
const lessonsText = ref('')
const improvementsText = ref('')

// 计算属性
const periodPlanLabel = computed(() => {
  const labels = {
    daily: '明日计划',
    weekly: '下周计划',
    monthly: '下月计划'
  }
  return labels[formData.period] || '下期计划'
})

// 获取日期选择器占位符
const getDatePickerPlaceholder = () => {
  const placeholders = {
    daily: '请选择复盘日期',
    weekly: '请选择周复盘',
    monthly: '请选择月复盘日期'
  }
  return placeholders[formData.period] || '请选择日期'
}


// 监听周期变化，清空日期选择
watch(() => formData.period, () => {
  formData.reviewDate = null
})

// 处理周选择器变化
const handleWeekChange = (value) => {
  if (Array.isArray(value) && value.length > 0) {
    // 如果是数组，取第一个值（周的开始日期）
    formData.reviewDate = value[0]
  } else {
    formData.reviewDate = value
  }
}

// 表单验证规则
const formRules = {
  title: [
    { required: true, message: '请输入复盘标题', trigger: 'blur' }
  ],
  period: [
    { required: true, message: '请选择复盘周期', trigger: 'change' }
  ],
  reviewDate: [
    { 
      required: true, 
      message: computed(() => {
        const messages = {
          daily: '请选择复盘日期',
          weekly: '请选择周复盘',
          monthly: '请选择月复盘日期'
        }
        return messages[formData.period] || '请选择日期'
      }), 
      trigger: 'change' 
    }
  ],
  totalTrades: [
    { required: true, message: '请输入总交易次数', trigger: 'blur' }
  ],
  winTrades: [
    { required: true, message: '请输入盈利交易次数', trigger: 'blur' }
  ],
  lossTrades: [
    { required: true, message: '请输入亏损交易次数', trigger: 'blur' }
  ],
  winRate: [
    { required: true, message: '请输入胜率', trigger: 'blur' }
  ],
  totalProfit: [
    { required: true, message: '请输入总盈亏', trigger: 'blur' }
  ]
}

// 更新经验教训
const updateLessons = () => {
  formData.lessons = lessonsText.value
    .split('\n')
    .map(item => item.trim())
    .filter(item => item)
}

// 更新改进计划
const updateImprovements = () => {
  formData.improvements = improvementsText.value
    .split('\n')
    .map(item => item.trim())
    .filter(item => item)
}

// 提交表单
const handleSubmit = async () => {
  const result = await formRef.value.validate()
  if (result === true) {
    showPreviewDialog.value = true
  }
}

// 确认创建/更新
const confirmCreate = async () => {
  isSubmitting.value = true
  
  try {
    if (isEditMode.value) {
      // 编辑模式：更新现有复盘
      const reviewData = {
        ...formData,
        id: reviewId.value,
        updateTime: new Date().toLocaleString()
      }
      
      await tradingReviewStore.updateReview(reviewId.value, reviewData)
      MessagePlugin.success('复盘更新成功！')
    } else {
      // 新建模式：创建新复盘
      const reviewData = {
        ...formData,
        id: Date.now().toString(),
        createTime: new Date().toLocaleString(),
        updateTime: new Date().toLocaleString(),
        status: 'completed'
      }
      
      await tradingReviewStore.addReview(reviewData)
      MessagePlugin.success('复盘创建成功！')
    }
    
    showPreviewDialog.value = false
    router.push('/trading-review')
  } catch (error) {
    MessagePlugin.error('保存失败，请重试')
  } finally {
    isSubmitting.value = false
  }
}

// 保存草稿
const saveDraft = () => {
  MessagePlugin.success('草稿已保存')
}

// 重置表单
const resetForm = () => {
  formRef.value.reset()
  Object.assign(formData, {
    title: '',
    period: 'daily',
    reviewDate: null,
    totalTrades: null,
    winTrades: null,
    lossTrades: null,
    winRate: null,
    totalProfit: null,
    maxLoss: null,
    avgProfit: null,
    avgLoss: null,
    content: '',
    lessons: [],
    improvements: [],
    nextPlan: ''
  })
  lessonsText.value = ''
  improvementsText.value = ''
}

// 加载编辑数据
const loadEditData = async () => {
  if (isEditMode.value && reviewId.value) {
    try {
      const review = await tradingReviewStore.loadReview(reviewId.value)
      if (review) {
        Object.assign(formData, review)
        lessonsText.value = review.lessons.join('\n')
        improvementsText.value = review.improvements.join('\n')
      } else {
        MessagePlugin.error('未找到要编辑的复盘')
        goBack()
      }
    } catch (error) {
      console.error('加载复盘失败:', error)
      MessagePlugin.error('加载复盘失败')
    }
  }
}

// 获取预览数据
const getPreviewData = () => {
  return [
    { label: '复盘标题', content: formData.title },
    { label: '复盘周期', content: getPeriodText(formData.period) },
    { label: '复盘日期', content: formData.reviewDate },
    { label: '总交易次数', content: formData.totalTrades },
    { label: '盈利交易', content: formData.winTrades },
    { label: '亏损交易', content: formData.lossTrades },
    { label: '胜率', content: `${formData.winRate}%` },
    { label: '总盈亏', content: formatProfit(formData.totalProfit) },
    { label: '平均盈利', content: formatProfit(formData.avgProfit) },
    { label: '平均亏损', content: formatProfit(formData.avgLoss) },
    { label: '最大亏损', content: formatProfit(formData.maxLoss) }
  ]
}

// 获取周期文本
const getPeriodText = (period) => {
  const texts = {
    daily: '日复盘',
    weekly: '周复盘',
    monthly: '月复盘'
  }
  return texts[period] || period
}

// 格式化盈亏
const formatProfit = (profit) => {
  if (profit === null || profit === undefined) return '--'
  return profit >= 0 ? `+${profit.toFixed(2)}` : profit.toFixed(2)
}

// 返回
const goBack = () => {
  router.back()
}

// 组件挂载时加载数据
onMounted(() => {
  loadEditData()
})
</script>

<style scoped>
.card-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-weight: 600;
}


</style>
