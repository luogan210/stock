import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/home/Home.vue'
import TradingPlan from '../views/trading-plan/TradingPlan.vue'
import TradingPlanForm  from '../views/trading-plan/TradingPlanForm.vue'
import TradingLog from '../views/trading-log/TradingLog.vue'
import TradingLogForm from '../views/trading-log/TradingLogForm.vue'
import TradingReview from '../views/trading-review/TradingReview.vue'
import TradingReviewForm from '../views/trading-review/TradingReviewForm.vue'
import StockManagement from '../views/stock/StockManagement.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/stock',
    redirect: '/trading-plan'
  },
  {
    path: '/trading-plan',
    name: 'TradingPlan',
    component: TradingPlan,
    meta: {
      title: '交易计划'
    }
  },
  {
    path: '/trading-plan/create',
    name: 'CreateTradingPlan',
    component: TradingPlanForm,
    meta: {
      title: '新建交易计划'
    }
  },
  {
    path: '/trading-plan/edit/:id',
    name: 'EditTradingPlan',
    component: TradingPlanForm,
    meta: {
      title: '编辑交易计划'
    }
  },
  {
    path: '/trading-log',
    name: 'TradingLog',
    component: TradingLog,
    meta: {
      title: '交易日志'
    }
  },
  {
    path: '/trading-log/create',
    name: 'TradingLogForm',
    component: TradingLogForm,
    meta: {
      title: '新建交易日志'
    }
  },
  {
    path: '/trading-log/edit/:id',
    name: 'EditTradingLog',
    component: TradingLogForm,
    meta: {
      title: '编辑交易日志'
    }
  },
  {
    path: '/trading-review',
    name: 'TradingReview',
    component: TradingReview,
    meta: {
      title: '交易复盘'
    }
  },
  {
    path: '/trading-review/create',
    name: 'CreateTradingReview',
    component: TradingReviewForm,
    meta: {
      title: '新建交易复盘'
    }
  },
  {
    path: '/trading-review/edit/:id',
    name: 'EditTradingReview',
    component: TradingReviewForm,
    meta: {
      title: '编辑交易复盘'
    }
  },
  {
    path: '/stock',
    name: 'StockManagement',
    component: StockManagement,
    meta: {
      title: '股票管理'
    }
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/'
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  // 设置页面标题
  if (to.meta.title) {
    document.title = `${to.meta.title} - Vue3 TDesign Stock`
  } else {
    document.title = 'Vue3 TDesign Stock'
  }
  next()
})

export default router
