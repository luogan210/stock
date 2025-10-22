<template>
  <div class="trading-review-form">
    <t-form
      ref="formRef"
      :data="formData"
      :rules="formRules"
      label-width="120px"
      scroll-to-first-error="smooth"
      @submit="onFormSubmit"
    >
      <t-form-item label="复盘周期" name="period">
        <t-radio-group v-model="formData.period" @change="handlePeriodChange">
          <t-radio value="daily">日复盘</t-radio>
          <t-radio value="weekly">周复盘</t-radio>
          <t-radio value="monthly">月复盘</t-radio>
        </t-radio-group>
      </t-form-item>
      
      <t-form-item label="复盘时间" name="reviewDate">
        <t-date-picker 
          v-if="formData.period === 'daily'"
          v-model="formData.reviewDate" 
          placeholder="选择复盘日期"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
          @change="generateTitle"
        />
        <t-date-picker 
          v-else-if="formData.period === 'weekly'"
          v-model="formData.reviewDate" 
          placeholder="选择复盘周"
          mode="week"
          :first-day-of-week="weekStart"
          format="YYYY-MM-DD"
          value-format="YYYY-MM-DD"
          clearable
          allow-input
          @change="generateTitle"
        />
        <t-date-picker 
          v-else-if="formData.period === 'monthly'"
          v-model="formData.reviewDate" 
          placeholder="选择复盘月份"
          mode="month"
          format="YYYY年MM月"
          value-format="YYYY-MM"
          @change="generateTitle"
        />
      </t-form-item>
      
      <t-form-item label="复盘标题" name="title">
        <t-input 
          class="form-input-md" 
          v-model="formData.title" 
          placeholder="请输入复盘标题"
        />
      </t-form-item>
      
      <t-form-item label="交易统计" name="trades">
        <div class="trades-stats">
          <div class="trade-stat-item buy-stat">
            <label class="stat-label">买入次数</label>
            <t-input-number 
              v-model="formData.buyCount" 
              placeholder="请输入买入次数"
              :min="0"
              class="stat-input"
            />
          </div>
          <div class="trade-stat-item sell-stat">
            <label class="stat-label">卖出次数</label>
            <t-input-number 
              v-model="formData.sellCount" 
              placeholder="请输入卖出次数"
              :min="0"
              class="stat-input"
            />
          </div>
        </div>
      </t-form-item>
      
      <t-form-item label="盈亏情况" name="profitLoss">
        <t-input-number 
          class="form-input-md" 
          v-model="formData.totalProfit" 
          placeholder="请输入盈亏金额"
          :precision="2"
        />
      </t-form-item>
      
      <t-form-item label="复盘总结" name="summary">
        <t-textarea 
          class="form-input-lg" 
          v-model="formData.summary" 
          placeholder="请总结本次复盘的要点和收获"
          :maxlength="1000"
          :autosize="{ minRows: 4, maxRows: 8 }"
        />
      </t-form-item>
      
      <t-form-item label="改进建议" name="improvements">
        <t-textarea 
          class="form-input-lg" 
          v-model="formData.improvements" 
          placeholder="请提出具体的改进建议"
          :maxlength="500"
          :autosize="{ minRows: 3, maxRows: 6 }"
        />
      </t-form-item>
    </t-form>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'
import dayjs from 'dayjs';
import updateLocale from 'dayjs/plugin/updateLocale';
import isoWeek from 'dayjs/plugin/isoWeek';
import weekOfYear from 'dayjs/plugin/weekOfYear';

dayjs.extend(updateLocale);
dayjs.extend(isoWeek);
dayjs.extend(weekOfYear);
dayjs.updateLocale('zh-cn', {
  weekStart: 1,
});
const weekStart = ref(1);
const props = defineProps({
  isEditMode: {
    type: Boolean,
    default: false
  }
})

const emit = defineEmits(['submit'])

const formRef = ref()

// 表单数据
const formData = reactive({
  period: 'daily',
  reviewDate: '',
  title: '',
  buyCount: 0,
  sellCount: 0,
  totalProfit: 0,
  summary: '',
  improvements: ''
})

// 表单验证规则
const formRules = {
  period: [
    { required: true, message: '请选择复盘周期', trigger: 'change' }
  ],
  reviewDate: [
    { required: true, message: '请选择复盘时间', trigger: 'change' }
  ],
  title: [
    { required: true, message: '请输入复盘标题', trigger: 'blur' }
  ],
  summary: [
    { required: true, message: '请输入复盘总结', trigger: 'blur' }
  ]
}

const handlePeriodChange = () => {
  // 清空日期选择
  formData.reviewDate = ''
  formData.title = ''
}

// 生成复盘标题
const generateTitle = () => {
  if (!formData.reviewDate) {
    formData.title = ''
    return
  }
  
  const date = dayjs(formData.reviewDate)
  
  switch (formData.period) {
    case 'daily':
      formData.title = `${date.format('YYYY年MM月DD日')}复盘`
      break
    case 'weekly':
      // 获取周的开始和结束日期
      const weekStart = date.startOf('week').add(1, 'day') // 周一
      const weekEnd = date.endOf('week').add(1, 'day') // 周日
      formData.title = `${weekStart.format('MM月DD日')}-${weekEnd.format('MM月DD日')}周复盘`
      break
    case 'monthly':
      formData.title = `${date.format('YYYY年MM月')}复盘`
      break
    default:
      formData.title = ''
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

defineExpose({
  getFormData,
  setFormData,
  validate
})
</script>

<style scoped>
.trading-review-form {
  padding: 20px 0;
  padding-top: 0;
}

.form-input-md {
  width: 300px;
}

.form-input-lg {
  width: 100%;
}

.trades-stats {
  display: flex;
  gap: 20px;
  align-items: flex-start;
}

.trade-stat-item {
  display: flex;
  flex-direction: column;
  gap: 8px;
  flex: 1;
}

.stat-label {
  font-size: 14px;
  font-weight: 500;
  color: var(--td-text-color-primary);
  margin-bottom: 4px;
}

.stat-input {
  width: 100%;
}

.buy-stat .stat-label {
  color: var(--td-success-color);
}

.sell-stat .stat-label {
  color: var(--td-error-color);
}
</style>