# Services 目录说明

## 目录结构

```
src/services/
├── api.js            # 统一 API 服务入口 - 集中管理所有业务 API
├── apiService.js     # 通用 API 工具类 - 只提供 HTTP 请求工具方法
├── http.js           # HTTP 客户端配置
├── interceptors.js   # 请求拦截器
└── README.md         # 本说明文件
```

## 文件说明

### api.js
- **作用**: 统一 API 服务入口，集中管理所有业务 API
- **内容**: 
  - 导入各个页面的 API 服务
  - 根据环境自动选择真实 API 或模拟 API
  - 避免循环依赖问题

### apiService.js
- **作用**: 通用 API 工具类，只提供 HTTP 请求工具方法
- **内容**: 
  - 通用请求方法 (`request`)
  - 统一处理 API 基础路径
  - 不涉及具体业务逻辑

### http.js
- **作用**: HTTP 客户端配置
- **内容**: 基于 axios 的 HTTP 客户端实例

### interceptors.js
- **作用**: 请求和响应拦截器
- **内容**: 
  - 请求拦截器（添加 token、通用参数等）
  - 响应拦截器（统一错误处理、业务错误处理等）

## 使用方式

### 业务逻辑接口规范

**重要：所有业务逻辑的接口都应该通过 `@/services/api` 来统一管理，避免循环依赖。**

```javascript
// ✅ 正确：通过统一 API 入口使用
import { currentApi } from '@/services/api'

// 使用当前环境的 API（开发环境自动使用模拟 API）
const plans = await currentApi.tradingPlan.getPlans()
const logs = await currentApi.tradingLog.getLogs()
const stats = await currentApi.home.getHomeStats()

// 或者直接使用特定 API
import { tradingPlanApi, tradingLogApi, homeApi } from '@/services/api'

// 使用通用请求工具（仅用于页面 API 内部）
import { request } from '@/services/apiService'
```

## 页面 API 位置

各个页面的 API 服务位于对应的页面目录下：

- 首页 API: `src/views/home/api/`
- 交易计划 API: `src/views/trading-plan/api/`
- 交易日志 API: `src/views/trading-log/api/`

每个页面 API 目录包含：
- `index.js` - 真实 API 服务
- `mock.js` - 模拟 API 服务（用于开发阶段）
