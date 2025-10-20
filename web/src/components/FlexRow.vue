<template>
  <div class="flex-row" :style="rowStyle">
    <slot />
  </div>
</template>

<script setup>
import { computed } from 'vue'

const props = defineProps({
  // 间距
  gap: {
    type: [Number, String],
    default: 16
  },
  // 对齐方式
  align: {
    type: String,
    default: 'flex-start',
    validator: (value) => ['flex-start', 'flex-end', 'center', 'baseline', 'stretch'].includes(value)
  },
  // 换行方式
  wrap: {
    type: String,
    default: 'nowrap',
    validator: (value) => ['nowrap', 'wrap', 'wrap-reverse'].includes(value)
  },
  // 主轴对齐方式
  justify: {
    type: String,
    default: 'flex-start',
    validator: (value) => ['flex-start', 'flex-end', 'center', 'space-between', 'space-around', 'space-evenly'].includes(value)
  },
  // 自定义样式
  style: {
    type: Object,
    default: () => ({})
  }
})

const rowStyle = computed(() => {
  const gapValue = typeof props.gap === 'number' ? `${props.gap}px` : props.gap
  
  return {
    display: 'flex',
    flexDirection: 'row',
    alignItems: props.align,
    flexWrap: props.wrap,
    justifyContent: props.justify,
    gap: gapValue,
    ...props.style
  }
})
</script>

<style scoped>
.flex-row {
  width: 100%;
}
</style>
