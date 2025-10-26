<template>
  <!-- 通用表单容器组件 - 只封装UI -->
  <div class="form-container">
    <t-card class="form-card">
      <template #header>
        <div class="card-header">
          <span>{{ title }}</span>
          <t-button theme="default" @click="$emit('go-back')" class="back-btn">
            <template #icon>
              <t-icon name="arrow-left" />
            </template>
            返回
          </t-button>
        </div>
      </template>
      
      <!-- 插槽：表单组件 -->
      <slot />
      <div class="form-buttons">
      <t-space size="small">
        <t-button @click="$emit('reset')" theme="default">
          <template #icon>
            <t-icon name="refresh" />
          </template>
          重置
        </t-button>
        <t-button @click="$emit('save-draft')" theme="default">
          <template #icon>
            <t-icon name="save" />
          </template>
          保存草稿
        </t-button>
        <t-button @click="$emit('submit')" theme="primary" :loading="isSubmitting">
          <template #icon>
            <t-icon name="check" />
          </template>
          {{ submitButtonText }}
        </t-button>
      </t-space>
    </div>
    </t-card>
    
    <!-- 按钮组 -->
  
  </div>
</template>

<script setup>
const props = defineProps({
  // 页面标题
  title: {
    type: String,
    required: true
  },
  // 提交按钮文本
  submitButtonText: {
    type: String,
    default: '提交'
  },
  // 是否正在提交
  isSubmitting: {
    type: Boolean,
    default: false
  }
})

// 只定义事件，不处理逻辑
const emit = defineEmits(['go-back', 'reset', 'save-draft', 'submit'])
</script>

<style scoped>
.form-container {
  display: flex;
  flex-direction: column;
  gap: 20px;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.back-btn {
  margin-left: auto;
}

.form-buttons {
  display: flex;
  justify-content: flex-start;
  padding: 8px;
  border-radius: 4px;
}
</style>
