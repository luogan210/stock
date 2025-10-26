<template>
  <FormContainer
    :title="isEditMode ? '编辑交易复盘' : '创建新的交易复盘'"
    :submit-button-text="isEditMode ? '更新复盘' : '创建复盘'"
    :is-submitting="isSubmitting"
    @go-back="goBack"
    @reset="resetForm"
    @save-draft="saveDraft"
    @submit="handleFormSubmit"
  >
    <TradingReviewForm 
      ref="formRef"
      :is-edit-mode="isEditMode"
      @submit="handleFormSubmit"
    />
  </FormContainer>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useTradingReviewStore } from './store'
import { MessagePlugin } from 'tdesign-vue-next'
import TradingReviewForm from './TradingReviewForm.vue'
import FormContainer from '@/components/FormContainer.vue'

const props = defineProps({
  isEditMode: {
    type: Boolean,
    default: false
  }
})

const router = useRouter()
const route = useRoute()
const tradingReviewStore = useTradingReviewStore()
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
    const reviewData = formRef.value.getFormData()
    
    if (props.isEditMode) {
      // 更新现有复盘
      await tradingReviewStore.updateReview(route.params.id, reviewData)
      MessagePlugin.success('交易复盘更新成功！')
    } else {
      // 创建新复盘
      await tradingReviewStore.addReview(reviewData)
      MessagePlugin.success('交易复盘创建成功！')
    }
    
    router.push('/trading-review')
  } catch (error) {
    MessagePlugin.error(`${props.isEditMode ? '更新' : '创建'}失败: ${error.message || '请重试'}`)
  } finally {
    isSubmitting.value = false
  }
}

/**
 * 加载复盘数据 - 编辑模式下从后端加载现有复盘
 */
const loadReviewData = async () => {
  if (!props.isEditMode) return
  
  try {
    const reviewId = route.params.id
    const review = await tradingReviewStore.getReviewById(reviewId)
    if (review) {
      // 通过 Form 的 setFormData 方法设置数据
      formRef.value.setFormData({
        period: review.period || 'daily',
        reviewDate: review.reviewDate || '',
        title: review.title || '',
        buyCount: review.buyCount || 0,
        sellCount: review.sellCount || 0,
        totalProfit: review.totalProfit || 0,
        summary: review.summary || '',
        improvements: review.improvements || ''
      })
    }
  } catch (error) {
    console.error('加载复盘数据失败:', error)
    MessagePlugin.error('加载复盘数据失败')
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
  loadReviewData()
})

// 暴露方法给父组件
defineExpose({
  isSubmitting
})
</script>

