<template>
  <div class="stock-form">
    <t-form
      ref="formRef"
      :data="formData"
      :rules="formRules"
      label-width="120px"
      scroll-to-first-error="smooth"
      @submit="onFormSubmit"
    >
      <t-form-item label="股票代码" name="code">
        <t-input 
          class="form-input-md" 
          v-model="formData.code" 
          placeholder="请输入股票代码"
          @blur="handleCodeBlur"
        />
      </t-form-item>
      
      <t-form-item label="股票名称" name="name">
        <t-input 
          class="form-input-md" 
          v-model="formData.name" 
          placeholder="请输入股票名称"
        />
      </t-form-item>
      
      <t-form-item label="地区" name="region">
        <t-select class="form-input-md" v-model="formData.region" placeholder="请选择地区">
          <t-option value="china" label="中国" />
          <t-option value="hongkong" label="香港" />
          <t-option value="usa" label="美国" />
        </t-select>
      </t-form-item>
      
      <t-form-item label="股票分类" name="category">
        <t-select class="form-input-md" v-model="formData.category" placeholder="请选择股票分类">
          <t-option value="main_board" label="主板" />
          <t-option value="hk_main" label="港股主板" />
          <t-option value="us_nasdaq" label="纳斯达克" />
          <t-option value="us_nyse" label="纽约证券交易所" />
        </t-select>
      </t-form-item>
      
      <t-form-item label="是否启用" name="enabled">
        <t-radio-group v-model="formData.enabled">
          <t-radio :value="true">启用</t-radio>
          <t-radio :value="false">禁用</t-radio>
        </t-radio-group>
      </t-form-item>
      
      <t-form-item label="备注" name="remark">
        <t-textarea 
          class="form-input-lg" 
          v-model="formData.remark" 
          placeholder="其他备注信息"
          :maxlength="200"
          :autosize="{ minRows: 2, maxRows: 4 }"
        />
      </t-form-item>
    </t-form>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { MessagePlugin } from 'tdesign-vue-next'

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
  code: '',
  name: '',
  region: '',
  category: '',
  enabled: true,
  remark: ''
})

// 表单验证规则
const formRules = {
  code: [
    { required: true, message: '请输入股票代码', trigger: 'blur' },
    { min: 1, max: 20, message: '股票代码长度应在1-20个字符', trigger: 'blur' }
  ],
  name: [
    { required: true, message: '请输入股票名称', trigger: 'blur' },
    { min: 1, max: 50, message: '股票名称长度应在1-50个字符', trigger: 'blur' }
  ],
  region: [
    { required: true, message: '请选择地区', trigger: 'change' }
  ],
  category: [
    { required: true, message: '请选择股票分类', trigger: 'change' }
  ],
}

// 股票代码失焦处理
const handleCodeBlur = () => {
  // 可以根据股票代码自动填充一些信息
  if (formData.code && !formData.name) {
    // 这里可以添加自动获取股票名称的逻辑
    console.log('股票代码:', formData.code)
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
.stock-form {
  padding: 20px 0;
  padding-top: 0;
}

.form-input-md {
  width: 300px;
}

.form-input-lg {
  width: 100%;
}
</style>
