# Vue3 TDesign Stock 项目

这是一个基于 Vue3 + TDesign + Pinia + Vue Router 构建的股票分析前端项目。

## 技术栈

- **Vue 3.4.0** - 渐进式 JavaScript 框架
- **TDesign Vue Next 1.8.0** - 企业级设计语言和 Vue 组件库
- **Pinia 2.1.7** - Vue 的状态管理库
- **Vue Router 4.2.5** - Vue.js 官方路由管理器
- **Vite 5.0.8** - 下一代前端构建工具

## 功能特性

- 🏠 **首页概览** - 系统状态和快速操作
- 📊 **股票分析** - 股票搜索、查看和关注功能
- 💰 **交易计划** - 创建、管理和执行交易计划
- 📝 **交易日志** - 记录、管理和分析交易历史
- 👤 **用户中心** - 个人信息管理和数据统计
- ⚙️ **系统设置** - 界面配置和数据管理

## 项目结构

```
src/
├── main.js              # 应用入口
├── App.vue              # 根组件
├── router/              # 路由配置
│   └── index.js
├── stores/              # Pinia 状态管理
│   └── index.js
└── views/               # 页面组件
    ├── Home.vue         # 首页
    ├── StockAnalysis.vue # 股票分析
    ├── TradingPlan.vue  # 交易计划列表
    ├── CreateTradingPlan.vue # 新建交易计划
    ├── EditTradingPlan.vue  # 编辑交易计划
    ├── TradingLog.vue   # 交易日志列表
    ├── CreateTradingLog.vue # 新建交易日志
    ├── UserCenter.vue   # 用户中心
    └── Settings.vue     # 系统设置
```

## 快速开始

### 安装依赖

```bash
npm install
```

### 启动开发服务器

```bash
npm run dev
```

### 构建生产版本

```bash
npm run build
```

### 预览生产构建

```bash
npm run preview
```

## 开发说明

### 状态管理

项目使用 Pinia 进行状态管理，包含以下 store：

- `useUserStore` - 用户信息管理
- `useStockStore` - 股票数据管理
- `useTradingPlanStore` - 交易计划管理
- `useTradingLogStore` - 交易日志管理

### 路由配置

- `/` - 首页
- `/stock` - 股票分析
- `/trading-plan` - 交易计划列表
- `/trading-plan/create` - 新建交易计划
- `/trading-plan/edit/:id` - 编辑交易计划
- `/trading-log` - 交易日志列表
- `/trading-log/create` - 新建交易日志
- `/user` - 用户中心
- `/settings` - 系统设置

### 组件库

项目集成了 TDesign Vue Next 组件库，支持按需引入和自动导入。

## 浏览器支持

- Chrome >= 87
- Firefox >= 78
- Safari >= 14
- Edge >= 88

## 许可证

MIT License
