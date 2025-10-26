<template>
  <FormContainer
    :title="isEditMode ? '编辑交易计划' : '创建新的交易计划'"
    :submit-button-text="isEditMode ? '更新计划' : '创建计划'"
    :is-submitting="isSubmitting"
    @go-back="goBack"
    @reset="resetForm"
    @save-draft="saveDraft"
    @submit="handleFormSubmit"
  >
    <TradingPlanForm 
      ref="formRef"
      :is-edit-mode="isEditMode"
      @submit="handleFormSubmit"
    />
  </FormContainer>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useTradingPlanStore } from './store'
import { MessagePlugin } from 'tdesign-vue-next'
import TradingPlanForm from './TradingPlanForm.vue'
import FormContainer from '@/components/FormContainer.vue'

const props = defineProps({
  isEditMode: {
    type: Boolean,
    default: false
  }
})

const router = useRouter()
const route = useRoute()
const tradingPlanStore = useTradingPlanStore()
const formRef = ref()

const isSubmitting = ref(false)

/**
 * 处理表单提交
 */
const handleFormSubmit = async () => {
  isSubmitting.value = true
  
  try {
    // 先验证表单
    const validateResult = await formRef.value.validate()
    if (validateResult !== true) {
      MessagePlugin.warning('请填写完整的表单信息')
      return
    }
    
    // 获取表单数据
    const planData = formRef.value.getFormData()
    
    if (props.isEditMode) {
      // 更新现有计划
      await tradingPlanStore.updatePlan(route.params.id, planData)
      MessagePlugin.success('交易计划更新成功！')
    } else {
      // 创建新计划
      await tradingPlanStore.addPlan(planData)
      MessagePlugin.success('交易计划创建成功！')
    }
    
    router.push('/trading-plan')
  } catch (error) {
    MessagePlugin.error(`${props.isEditMode ? '更新' : '创建'}失败: ${error.message || '请重试'}`)
  } finally {
    isSubmitting.value = false
  }
}

/**
 * 加载计划数据 - 编辑模式下从后端加载现有计划
 */
const loadPlanData = async () => {
  if (!props.isEditMode) return
  
  try {
    const planId = route.params.id
    const plan = await tradingPlanStore.getPlanById(planId)
    if (plan) {
      // 通过 Form 的 setFormData 方法设置数据
      formRef.value.setFormData({
        name: plan.name || '',
        type: plan.type || 'buy',
        stockCode: plan.stockCode || '',
        stockName: plan.stockName || '',
        targetPrice: plan.targetPrice || null,
        quantity: plan.quantity || null,
        stopLoss: plan.stopLoss || null,
        takeProfit: plan.takeProfit || null,
        startTime: plan.startTime || '',
        endTime: plan.endTime || '',
        riskLevel: plan.riskLevel || 'medium',
        description: plan.description || '',
        remark: plan.remark || '',
        strategy: plan.strategy || '',
        tradingStrategy: plan.tradingStrategy || ''
      })
    }
  } catch (error) {
    console.error('加载计划数据失败:', error)
    MessagePlugin.error('加载计划数据失败')
  }
}

/**
 * 重置表单
 */
const resetForm = () => {
  MessagePlugin.success('已重置')
}

/**
 * 保存草稿
 */
const saveDraft = () => {
  MessagePlugin.success('草稿已保存')
}

/**
 * 返回上一页
 */
const goBack = () => {
  router.back()
}

/**
 * 确认提交 - 触发表单验证和提交
 */
const confirmCreate = () => {
  handleFormSubmit()
}

// 组件挂载时加载数据
onMounted(() => {
  loadPlanData()
})

// 暴露方法给父组件
defineExpose({
  isSubmitting
})
</script>

