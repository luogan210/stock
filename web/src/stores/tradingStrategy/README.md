# 交易策略 Store - 简化版

这是一个简化的交易策略store，只负责交易策略数据存储，不包含复杂的API调用和业务逻辑。

## 功能

- 存储交易策略数据
- 提供基本的增删改查操作
- 支持按分类筛选策略

## 使用方法

```javascript
import { useTradingStrategyStore } from '@/stores/tradingStrategy'

const tradingStrategyStore = useTradingStrategyStore()

// 获取数据
const strategies = tradingStrategyStore.strategies

// 添加交易策略
tradingStrategyStore.addStrategy({
  name: '新交易策略',
  description: '交易策略描述',
  category: 'day_trading'
})

// 更新交易策略
tradingStrategyStore.updateStrategy(strategyId, {
  name: '更新后的交易策略名称'
})

// 删除交易策略
tradingStrategyStore.removeStrategy(strategyId)

// 获取启用的策略
const enabledStrategies = tradingStrategyStore.getEnabledStrategies()

// 根据分类获取策略
const dayTradingStrategies = tradingStrategyStore.getStrategiesByCategory('day_trading')
```

## API

### 数据
- `strategies` - 交易策略列表

### 方法
- `addStrategy(strategy)` - 添加交易策略
- `updateStrategy(id, updates)` - 更新交易策略
- `removeStrategy(id)` - 删除交易策略
- `getStrategyById(id)` - 根据ID获取交易策略
- `getEnabledStrategies()` - 获取启用的策略
- `getStrategiesByCategory(category)` - 根据分类获取策略
