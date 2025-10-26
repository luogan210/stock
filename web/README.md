# Vue3 TDesign Stock 股票分析系统

一个基于 Vue3 + TDesign + Pinia + Vue Router 的现代化股票分析和管理系统，提供完整的交易计划、日志记录、复盘分析等功能。

## 📋 项目概述

本项目是一个功能完整的股票分析系统，采用现代化的前端技术栈，提供直观的用户界面和丰富的功能模块。系统支持股票管理、交易计划制定、交易日志记录、复盘分析等核心功能。

## 🚀 技术栈

- **前端框架**: Vue 3.4.0
- **UI 组件库**: TDesign Vue Next 1.8.0
- **状态管理**: Pinia 2.1.7
- **路由管理**: Vue Router 4.2.5
- **HTTP 客户端**: Axios 1.12.2
- **构建工具**: Vite 5.0.8
- **代码规范**: ESLint + Prettier

## 📁 项目结构

```
web/
├── src/
│   ├── components/           # 公共组件
│   │   └── FlexRow.vue      # 弹性布局组件
│   ├── router/              # 路由配置
│   │   └── index.js         # 路由定义
│   ├── services/            # API 服务
│   │   ├── api.js           # 统一 API 入口
│   │   ├── apiService.js    # API 服务封装
│   │   ├── http.js          # HTTP 客户端配置
│   │   └── interceptors.js  # 请求拦截器
│   ├── stores/              # 状态管理
│   │   ├── index.js         # Store 统一入口
│   │   ├── user/            # 用户状态
│   │   └── stock/           # 股票状态
│   ├── styles/              # 样式文件
│   │   ├── theme.css        # 主题样式
│   │   └── form.css         # 表单样式
│   ├── utils/               # 工具函数
│   │   ├── constants.js     # 常量定义
│   │   ├── helpers.js       # 辅助函数
│   │   ├── mockData.js      # 模拟数据
│   │   ├── stockConfig.js   # 股票配置
│   │   ├── strategyConfig.js # 策略配置
│   │   └── tradingStrategyConfig.js # 交易策略配置
│   └── views/               # 页面组件
│       ├── home/            # 首页
│       │   ├── Home.vue     # 首页组件
│       │   ├── api/         # API 接口
│       │   └── store/       # 状态管理
│       ├── settings/        # 设置页面
│       ├── stock/           # 股票管理
│       │   └── StockManagement.vue
│       ├── trading-log/     # 交易日志
│       │   ├── TradingLog.vue
│       │   ├── TradingLogForm.vue
│       │   ├── api/         # API 接口
│       │   └── store/       # 状态管理
│       ├── trading-plan/    # 交易计划
│       │   ├── TradingPlan.vue
│       │   ├── TradingPlanForm.vue
│       │   ├── api/         # API 接口
│       │   └── store/       # 状态管理
│       └── trading-review/  # 交易复盘
│           ├── TradingReview.vue
│           ├── TradingReviewForm.vue
│           ├── api/         # API 接口
│           └── store/       # 状态管理
├── index.html               # HTML 入口
├── package.json             # 项目配置
├── vite.config.js          # Vite 配置
├── start.bat               # Windows 启动脚本
├── start.sh                # Linux/Mac 启动脚本
└── README.md               # 项目说明
```

## 🎯 核心功能

### 1. 首页概览
- 系统统计数据展示
- 快速操作入口
- 关键指标监控

### 2. 股票管理
- 股票信息维护
- 多市场支持（A股、港股、美股）
- 股票分类和标签管理
- 风险等级评估

### 3. 交易计划
- 制定交易策略
- 设置止盈止损
- 计划状态跟踪
- 策略配置管理

### 4. 交易日志
- 交易记录管理
- 盈亏统计
- 交易分析
- 历史数据查询

### 5. 交易复盘
- 交易回顾分析
- 经验总结
- 策略优化建议
- 学习记录

## 🛠️ 开发环境

### 环境要求
- Node.js >= 16.0.0
- npm >= 8.0.0

### 安装依赖
```bash
npm install
```

### 启动开发服务器
```bash
# 使用 npm
npm run dev

# 或使用启动脚本
# Windows
start.bat

# Linux/Mac
./start.sh
```

### 构建生产版本
```bash
npm run build
```

### 预览生产版本
```bash
npm run preview
```

### 代码检查
```bash
npm run lint
```

## 📊 数据配置

### 股票配置 (stockConfig.js)
- 支持多市场股票配置
- 包含股票基础信息、财务指标
- 风险等级和流动性评估
- 行业分类和标签管理

### 策略配置 (strategyConfig.js)
- 技术分析策略
- 基本面分析策略
- 策略参数配置
- 胜率和适用场景

### 常量定义 (constants.js)
- 交易类型和状态
- 风险等级定义
- 策略类型映射
- 分页配置

## 🔧 配置说明

### Vite 配置
- 自动导入 TDesign 组件
- 路径别名配置 (`@` 指向 `src`)
- 开发服务器配置（端口 3000）

### 路由配置
- 支持嵌套路由
- 路由守卫和标题设置
- 404 页面重定向

### 状态管理
- 模块化 Store 设计
- 统一的状态管理入口
- 支持异步操作

## 🎨 UI 设计

### 设计系统
- 基于 TDesign 设计语言
- 统一的视觉风格
- 响应式布局设计

### 主题配置
- 支持主题切换
- 自定义样式变量
- 组件样式覆盖

## 📱 功能特性

### 响应式设计
- 适配不同屏幕尺寸
- 移动端友好
- 触摸操作优化

### 数据管理
- 本地数据持久化
- 模拟数据支持
- 真实 API 集成

### 用户体验
- 直观的操作界面
- 快速的数据加载
- 友好的错误提示

## 🔄 开发流程

### 代码规范
- ESLint 代码检查
- Prettier 代码格式化
- 统一的代码风格

### 组件开发
- 组件化开发模式
- 可复用的业务组件
- 清晰的组件接口

### 状态管理
- 集中式状态管理
- 模块化的 Store
- 类型安全的状态操作

## 📈 性能优化

### 构建优化
- Vite 快速构建
- 代码分割和懒加载
- 资源压缩和优化

### 运行时优化
- 组件懒加载
- 虚拟滚动
- 内存管理

## 🚀 部署说明

### 生产构建
```bash
npm run build
```

### 静态部署
构建后的文件位于 `dist` 目录，可直接部署到静态文件服务器。

### 环境变量
支持通过环境变量配置 API 地址等配置项。

## 🤝 贡献指南

### 开发规范
1. 遵循 ESLint 和 Prettier 配置
2. 编写清晰的组件和函数注释
3. 保持代码风格一致

### 提交规范
- 使用清晰的提交信息
- 遵循约定式提交规范
- 包含必要的测试

## 📄 许可证

ISC License

## 📞 联系方式

如有问题或建议，请通过以下方式联系：
- 项目 Issues
- 邮箱联系
- 技术交流群

---

**注意**: 本项目为股票分析系统，仅供学习和研究使用，不构成投资建议。投资有风险，入市需谨慎。