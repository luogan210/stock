# 统一响应规范指南

## 概述

为了确保所有 API 响应的一致性，我们提供了多种统一响应方案，让开发者能够轻松遵循相同的响应规范。

## 方案1：使用 BaseHandler（推荐）

### 优势
- 简单易用，代码清晰
- 提供语义化的方法名
- 支持继承，便于扩展

### 使用方法

#### 1. 创建 Handler 时继承 BaseHandler

```go
type YourHandler struct {
    *BaseHandler
    yourService *service.YourService
}

func NewYourHandler() *YourHandler {
    return &YourHandler{
        BaseHandler: NewBaseHandler(),
        yourService: service.NewYourService(),
    }
}
```

#### 2. 使用统一的响应方法

```go
func (h *YourHandler) YourMethod(c *gin.Context) {
    // 成功响应
    h.Success(c, data)
    
    // 带自定义消息的成功响应
    h.SuccessWithMessage(c, data, "操作成功")
    
    // 参数错误
    h.ParamError(c, "参数错误信息")
    
    // 服务器错误
    h.ServerError(c, "服务器错误信息")
    
    // 未授权错误
    h.UnauthorizedError(c, "请先登录")
    
    // 资源未找到
    h.NotFoundError(c, "资源不存在")
}
```

### 可用的响应方法

| 方法名 | 用途 | 状态码 |
|--------|------|--------|
| `Success(c, data)` | 成功响应 | 200 |
| `SuccessWithMessage(c, data, message)` | 带自定义消息的成功响应 | 200 |
| `Error(c, code, message)` | 自定义错误响应 | 自定义 |
| `ErrorWithData(c, code, message, data)` | 带数据的错误响应 | 自定义 |
| `ParamError(c, message)` | 参数错误 | 400 |
| `ServerError(c, message)` | 服务器错误 | 500 |
| `UnauthorizedError(c, message)` | 未授权错误 | 401 |
| `NotFoundError(c, message)` | 资源未找到 | 404 |

## 方案2：使用中间件自动包装（高级）

### 优势
- 完全自动化，无需手动调用
- 减少重复代码
- 统一的错误处理

### 使用方法

#### 1. 在路由中注册中间件

```go
// 在 router.go 中添加
r.Use(middleware.ResponseWrapper())
```

#### 2. 在 Handler 中设置响应数据

```go
func (h *YourHandler) YourMethod(c *gin.Context) {
    data := map[string]interface{}{
        "message": "你的数据",
    }
    
    // 设置响应数据，中间件会自动包装
    middleware.SetResponseData(c, data)
}
```

## 响应格式

所有响应都遵循以下统一格式：

```json
{
    "code": 200,           // 业务状态码
    "message": "success",  // 响应消息
    "data": {              // 响应数据
        // 具体数据内容
    }
}
```

## 状态码规范

| 状态码 | 含义 | 使用场景 |
|--------|------|----------|
| 200 | 成功 | 正常业务操作成功 |
| 400 | 参数错误 | 请求参数验证失败 |
| 401 | 未授权 | 用户未登录或token无效 |
| 404 | 资源未找到 | 请求的资源不存在 |
| 500 | 服务器错误 | 服务器内部错误 |

## 最佳实践

### 1. 错误处理
```go
// 推荐：使用语义化的错误方法
h.ParamError(c, "用户名不能为空")

// 不推荐：直接调用 Error
h.Error(c, 400, "用户名不能为空")
```

### 2. 成功响应
```go
// 推荐：使用 Success 方法
h.Success(c, data)

// 推荐：需要自定义消息时
h.SuccessWithMessage(c, data, "用户创建成功")
```

### 3. 数据格式
```go
// 推荐：使用 map 或结构体
data := map[string]interface{}{
    "id": 1,
    "name": "张三",
    "email": "zhangsan@example.com",
}
h.Success(c, data)
```

## 示例

查看 `handler/example.go` 文件中的完整示例，包括：
- 各种成功响应示例
- 各种错误响应示例
- 中间件自动包装示例

## 迁移指南

### 从旧版本迁移

1. 将 Handler 结构体添加 `*BaseHandler` 字段
2. 在构造函数中初始化 `BaseHandler`
3. 将 `Success(c, data)` 替换为 `h.Success(c, data)`
4. 将 `Error(c, code, message)` 替换为对应的语义化方法

### 示例迁移

```go
// 旧版本
func (h *UserHandler) CreateUser(c *gin.Context) {
    // ... 业务逻辑
    Success(c, userInfo)
}

// 新版本
type UserHandler struct {
    *BaseHandler
    userService *service.UserService
}

func (h *UserHandler) CreateUser(c *gin.Context) {
    // ... 业务逻辑
    h.Success(c, userInfo)
}
```

## 注意事项

1. **继承顺序**：确保 `*BaseHandler` 是结构体的第一个字段
2. **初始化**：在构造函数中必须调用 `NewBaseHandler()`
3. **方法调用**：使用 `h.` 前缀调用基类方法
4. **错误处理**：优先使用语义化的错误方法，而不是通用的 `Error` 方法