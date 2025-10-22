import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/home/Home.vue'
import TradingPlan from '../views/trading-plan/TradingPlan.vue'
import PlanFormService from '../views/trading-plan/PlanFormService.vue'
import TradingLog from '../views/trading-log/TradingLog.vue'
import LogFormService from '../views/trading-log/LogFormService.vue'
import TradingReview from '../views/trading-review/TradingReview.vue'
import ReviewFormService from '../views/trading-review/ReviewFormService.vue'
import StockManagement from '../views/stock/StockManagement.vue'
import StockFormService from '../views/stock/StockFormService.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
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
    component: PlanFormService,
    meta: {
      title: '新建交易计划'
    }
  },
  {
    path: '/trading-plan/edit/:id',
    name: 'EditTradingPlan',
    component: PlanFormService,
    props: {
      isEditMode: true
    },
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
    name: 'CreateTradingLog',
    component: LogFormService,
    meta: {
      title: '新建交易日志'
    }
  },
  {
    path: '/trading-log/edit/:id',
    name: 'EditTradingLog',
    component: LogFormService,
    props: {
      isEditMode: true
    },
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
    component: ReviewFormService,
    meta: {
      title: '新建交易复盘'
    }
  },
  {
    path: '/trading-review/edit/:id',
    name: 'EditTradingReview',
    component: ReviewFormService,
    props: {
      isEditMode: true
    },
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
    path: '/stock/create',
    name: 'CreateStock',
    component: StockFormService,
    meta: {
      title: '新增股票'
    }
  },
  {
    path: '/stock/edit/:id',
    name: 'EditStock',
    component: StockFormService,
    props: {
      isEditMode: true
    },
    meta: {
      title: '编辑股票'
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
