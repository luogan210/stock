<template>
  <FormContainer
    :title="isEditMode ? '编辑股票信息' : '新增股票信息'"
    :submit-button-text="isEditMode ? '更新股票' : '新增股票'"
    :is-submitting="isSubmitting"
    @go-back="goBack"
    @reset="resetForm"
    @save-draft="saveDraft"
    @submit="handleFormSubmit"
  >
    <StockForm 
      ref="formRef"
      :is-edit-mode="isEditMode"
      @submit="handleFormSubmit"
    />
  </FormContainer>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { MessagePlugin } from 'tdesign-vue-next'
import StockForm from './StockForm.vue'
import FormContainer from '@/components/FormContainer.vue'
import { useStockStore } from './store'

const props = defineProps({
  isEditMode: {
    type: Boolean,
    default: false
  }
})

const router = useRouter()
const route = useRoute()
const stockStore = useStockStore()
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
    const stockData = formRef.value.getFormData()
    
    if (props.isEditMode) {
      // 更新现有股票
      await stockStore.updateStockById(route.params.id, stockData)
      MessagePlugin.success('股票信息更新成功！')
    } else {
      // 创建新股票
      await stockStore.addStock(stockData)
      MessagePlugin.success('股票信息创建成功！')
    }
    
    router.push('/stock')
  } catch (error) {
    MessagePlugin.error(`${props.isEditMode ? '更新' : '创建'}失败: ${error.message || '请重试'}`)
  } finally {
    isSubmitting.value = false
  }
}

/**
 * 加载股票数据 - 编辑模式下从后端加载现有股票
 */
const loadStockData = async () => {
  if (!props.isEditMode || !formRef.value) return
  
  try {
    const stockId = route.params.id
    const stock = await stockStore.getStockById(stockId)
    
    if (stock) {
      // 通过 Form 的 setFormData 方法设置数据
      formRef.value.setFormData({
        code: stock.code || '',
        name: stock.name || '',
        marketType: stock.marketType || '',
        category: stock.category || '',
        industry: stock.industry || '',
        marketCap: stock.marketCap || null,
        pe: stock.pe || null,
        pb: stock.pb || null,
        riskLevel: stock.riskLevel || 'medium',
        enabled: stock.enabled !== undefined ? stock.enabled : true,
        remark: stock.description || ''
      })
    }
  } catch (error) {
    console.error('加载股票数据失败:', error)
    MessagePlugin.error('加载股票数据失败')
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

// 组件挂载时加载数据
onMounted(() => {
  loadStockData()
})

// 暴露方法给父组件
defineExpose({
  isSubmitting
})
</script>
