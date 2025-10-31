# 选股策略 Store - 简化版

这是一个简化的选股策略store，只负责选股策略数据存储，不包含复杂的API调用和业务逻辑。

## 功能

- 存储选股策略数据
- 提供基本的增删改查操作
- 支持按分类筛选策略

## 使用方法

```javascript
import { useStrategyStore } from '@/stores/strategy'

const strategyStore = useStrategyStore()

// 获取数据
const strategies = strategyStore.strategies

// 添加选股策略
strategyStore.addStrategy({
  name: '新策略',
  description: '策略描述',
  category: 'technical'
})

// 更新选股策略
strategyStore.updateStrategy(strategyId, {
  name: '更新后的策略名称'
})

// 删除选股策略
strategyStore.removeStrategy(strategyId)

// 获取启用的策略
const enabledStrategies = strategyStore.getEnabledStrategies()

// 根据分类获取策略
const technicalStrategies = strategyStore.getStrategiesByCategory('technical')
```

## API

### 数据
- `strategies` - 选股策略列表

### 方法
- `addStrategy(strategy)` - 添加选股策略
- `updateStrategy(id, updates)` - 更新选股策略
- `removeStrategy(id)` - 删除选股策略
- `getStrategyById(id)` - 根据ID获取选股策略
- `getEnabledStrategies()` - 获取启用的策略
- `getStrategiesByCategory(category)` - 根据分类获取策略
