<template>
  <div id="app">
    <t-layout>
      <t-header>
        <t-head-menu theme="dark" :value="currentMenuValue">
          <t-menu-item value="1" @click="goToHome">
            <template #icon>
              <t-icon name="home" />
            </template>
            首页
          </t-menu-item>
          <t-menu-item value="2" @click="goToTradingPlan">
            <template #icon>
              <t-icon name="money" />
            </template>
            交易计划
          </t-menu-item>
          <t-menu-item value="3" @click="goToTradingLog">
            <template #icon>
              <t-icon name="file" />
            </template>
            交易日志
          </t-menu-item>
          <t-menu-item value="4" @click="goToTradingReview">
            <template #icon>
              <t-icon name="chart" />
            </template>
            交易复盘
          </t-menu-item>
          <t-menu-item value="5" @click="goToStockManagement">
            <template #icon>
              <t-icon name="stock" />
            </template>
            股票管理
          </t-menu-item>
        </t-head-menu>
      </t-header>
      <t-content>
        <router-view />
      </t-content>
    </t-layout>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const router = useRouter()
const route = useRoute()

// 根据当前路由计算菜单值
const currentMenuValue = computed(() => {
  const path = route.path
  if (path === '/' || path.startsWith('/home')) {
    return '1'
  } else if (path.startsWith('/trading-plan')) {
    return '2'
  } else if (path.startsWith('/trading-log')) {
    return '3'
  } else if (path.startsWith('/trading-review')) {
    return '4'
  }
  return '1' // 默认选中首页
})

const goToHome = () => {
  router.push('/')
}

const goToTradingPlan = () => {
  router.push('/trading-plan')
}

const goToTradingLog = () => {
  router.push('/trading-log')
}

const goToTradingReview = () => {
  router.push('/trading-review')
}

const goToStockManagement = () => {
  router.push('/stock')
}
</script>

<style>
#app {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, 'Open Sans', 'Helvetica Neue', sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

body {
  margin: 0;
  padding: 0;
}

/* 防止横向滚动条 */
* {
  box-sizing: border-box;
}

html, body {
  overflow-x: hidden;
}

#app {
  max-width: 100vw;
  overflow-x: hidden;
}

/* 固定头部导航 */
.t-layout__header {
  position: fixed !important;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
  width: 100%;
}

/* 为内容区域添加顶部边距，避免被固定头部遮挡 */
.t-layout__content {
  margin-top: var(--td-comp-size-xxl) !important;
  min-height: calc(100vh - var(--td-comp-size-xxl));
  padding: 16px;
}
</style>
