<template>
  <FormContainer
    :title="isEditMode ? '编辑交易日志' : '创建新的交易日志'"
    :submit-button-text="isEditMode ? '更新日志' : '创建日志'"
    :is-submitting="isSubmitting"
    @go-back="goBack"
    @reset="resetForm"
    @save-draft="saveDraft"
    @submit="handleFormSubmit"
  >
    <TradingLogForm 
      ref="formRef"
      :is-edit-mode="isEditMode"
      @submit="handleFormSubmit"
    />
  </FormContainer>
</template>

<script setup>
import { ref, onMounted, nextTick } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useTradingLogStore } from './store'
import { MessagePlugin } from 'tdesign-vue-next'
import TradingLogForm from './TradingLogForm.vue'
import FormContainer from '@/components/FormContainer.vue'

const props = defineProps({
  isEditMode: {
    type: Boolean,
    default: false
  }
})

const router = useRouter()
const route = useRoute()
const tradingLogStore = useTradingLogStore()
const formRef = ref()
const isSubmitting = ref(false)



/**
 * 处理表单提交
 */
const handleFormSubmit = async () => {
  if (!formRef.value) {
    MessagePlugin.warning('表单未就绪')
    return
  }

  isSubmitting.value = true
  
  try {
    // 校验表单
    const validateResult = await formRef.value.validate()
    if (validateResult !== true) {
      MessagePlugin.warning('请填写完整的表单信息')
      isSubmitting.value = false
      return
    }
   
    const logData = formRef.value.getFormData()

    if (props.isEditMode) {
      await tradingLogStore.updateLog(route.params.id, logData)
      MessagePlugin.success('交易日志更新成功！')
    } else {
      await tradingLogStore.addLog(logData)
      MessagePlugin.success('交易日志创建成功！')
    }
    // 成功后跳转到日志列表
    router.push('/trading-log')
  } catch (error) {
    MessagePlugin.error(`${props.isEditMode ? '更新' : '创建'}失败: ${error.message || '请重试'}`)
  } finally {
    isSubmitting.value = false
  }
}

/**
 * 处理提交（保留用于子组件调用）
 */
const handleSubmit = async (formRef) => {
  const logData = formRef.getFormData()
  
  if (props.isEditMode) {
    await tradingLogStore.updateLog(route.params.id, logData)
    MessagePlugin.success('交易日志更新成功！')
  } else {
    await tradingLogStore.addLog(logData)
    MessagePlugin.success('交易日志创建成功！')
  }
}

/**
 * 加载日志数据 - 编辑模式下从后端加载现有日志
 */
const loadLogData = async (formRef) => {
  if (!props.isEditMode || !formRef) return
  
  try {
    const logId = route.params.id
    const log = await tradingLogStore.getLogById(logId)
    if (log) {
      formRef.setFormData({
        planName: log.planName || '',
        stockCode: log.stockCode || '',
        stockName: log.stockName || '',
        type: log.type || 'buy',
        tradingTime: log.tradingTime || '',
        price: log.price || null,
        quantity: log.quantity || null,
        strategy: log.strategy || '',
        remark: log.remark || ''
      })
    }
  } catch (error) {
    console.error('加载日志数据失败:', error)
    MessagePlugin.error('加载日志数据失败')
  }
}

/**
 * 重置表单
 */
const resetForm = () => {
  // 重置逻辑
}

/**
 * 保存草稿
 */
const saveDraft = () => {
  // 保存草稿逻辑
}

/**
 * 返回上一页
 */
const goBack = () => {
  router.back()
}

// 组件挂载时加载数据
onMounted(async () => {
  if (props.isEditMode) {
    await nextTick()
    if (formRef.value) {
      loadLogData(formRef.value)
    }
  }
})
</script>

