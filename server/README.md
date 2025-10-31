# Go Gin Stock API Server

一个基于 Go + Gin + SQLite 的现代化股票分析系统后端服务，提供完整的股票管理、交易计划、日志记录、复盘分析等API接口。

## 📋 项目概述

本项目是一个功能完整的股票分析系统后端服务，采用现代化的Go技术栈，提供RESTful API接口。系统支持股票管理、交易计划制定、交易日志记录、复盘分析等核心功能，并集成了文件上传、数据持久化等企业级特性。

## 🚀 技术栈

- **Web框架**: Gin 1.9.1
- **数据库**: SQLite (modernc.org/sqlite)
- **配置管理**: godotenv
- **UUID生成**: google/uuid
- **HTTP客户端**: 内置net/http
- **JSON处理**: 内置encoding/json
- **文件上传**: 分片上传支持
- **数据库迁移**: 自动建表
- **中间件**: 认证、响应统一处理

## 📁 项目结构

```
server/
├── main.go                    # 应用入口
├── go.mod / go.sum           # Go模块依赖
├── config/                   # 配置管理
│   ├── app.env              # 环境变量配置
│   ├── config.go            # 配置加载器
│   └── wechat.go            # 微信配置
├── router/                   # 路由管理
│   ├── router.go            # 主路由配置
│   └── frontend.go          # 前端路由
├── handler/                  # 处理器层
│   ├── base.go              # 基础处理器
│   ├── response.go          # 响应处理
│   ├── upload.go            # 文件上传处理
│   └── frontend.go          # 前端页面处理
├── middleware/              # 中间件
│   ├── auth.go              # 认证中间件
│   └── response.go          # 响应中间件
├── modules/                 # 业务模块
│   ├── stock/               # 股票管理模块
│   │   ├── handler.go       # 股票处理器
│   │   ├── model.go         # 股票模型
│   │   ├── service.go       # 股票服务
│   │   └── stock_repository.go # 股票仓储
│   ├── plan/                # 交易计划模块
│   │   ├── handler.go       # 计划处理器
│   │   ├── model.go         # 计划模型
│   │   ├── service.go       # 计划服务
│   │   └── plan_repository.go # 计划仓储
│   ├── log/                 # 交易日志模块
│   │   ├── handler.go       # 日志处理器
│   │   ├── model.go         # 日志模型
│   │   └── service.go       # 日志服务
│   ├── review/              # 交易复盘模块
│   │   ├── handler.go       # 复盘处理器
│   │   ├── model.go         # 复盘模型
│   │   └── service.go       # 复盘服务
│   └── modules.go           # 模块注册
├── storage/                 # 数据存储
│   └── sqlite.go           # SQLite数据库
├── db/                      # 数据库全局
│   └── global.go           # 全局数据库实例
├── utils/                   # 工具函数
│   ├── id.go               # ID生成工具
│   └── response.go        # 响应工具
├── frontend/               # 前端静态文件
│   ├── static/            # 静态资源
│   ├── templates/         # 模板文件
│   └── test_upload.html   # 上传测试页面
├── examples/              # 示例代码
│   └── wechat_example.go # 微信集成示例
├── data/                  # 数据目录
│   └── app.db            # SQLite数据库文件
└── uploads/              # 上传文件目录
```

## 🎯 核心功能

### 1. 股票管理
- 股票信息CRUD操作
- 多条件搜索和过滤
- 分页查询支持
- 股票状态管理

### 2. 交易计划
- 交易计划制定和管理
- 策略配置和参数设置
- 计划状态跟踪
- 风险等级评估

### 3. 交易日志
- 交易记录管理
- 盈亏统计和分析
- 历史数据查询
- 实时数据同步

### 4. 交易复盘
- 交易回顾分析
- 经验总结记录
- 策略优化建议
- 绩效评估报告

### 5. 文件上传
- 分片上传支持
- 大文件处理
- 上传进度跟踪
- 文件完整性验证

### 6. 数据管理
- SQLite数据库
- 自动表结构迁移
- 数据持久化
- 事务支持

## 🛠️ 开发环境

### 环境要求
- Go >= 1.20
- SQLite3 支持

### 安装依赖
```bash
go mod tidy
```

### 启动开发服务器
```bash
# 直接运行
go run main.go

# 或编译后运行
go build -o server main.go
./server
```

### 环境配置
```bash
# 复制环境配置文件
cp config/app.env config/app.env.local

# 编辑配置文件
vim config/app.env.local
```

## 📊 数据库设计

### 股票表 (stocks)
```sql
CREATE TABLE stocks (
    id TEXT PRIMARY KEY,
    code TEXT NOT NULL,
    name TEXT NOT NULL,
    region TEXT,
    currency TEXT,
    category TEXT,
    enabled INTEGER DEFAULT 1,
    remark TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### 交易计划表 (plans)
```sql
CREATE TABLE plans (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT,
    stock_code TEXT,
    stock_name TEXT,
    strategy TEXT,
    trading_strategy TEXT,
    target_price REAL,
    quantity INTEGER,
    stop_loss REAL,
    take_profit REAL,
    start_time TEXT,
    end_time TEXT,
    risk_level TEXT,
    description TEXT,
    remark TEXT,
    status TEXT DEFAULT 'active',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### 交易日志表 (logs)
```sql
CREATE TABLE logs (
    id TEXT PRIMARY KEY,
    plan_name TEXT,
    stock_code TEXT NOT NULL,
    stock_name TEXT,
    type TEXT NOT NULL,
    trading_time TEXT NOT NULL,
    price REAL NOT NULL,
    quantity INTEGER NOT NULL,
    strategy TEXT,
    remark TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

### 交易复盘表 (reviews)
```sql
CREATE TABLE reviews (
    id TEXT PRIMARY KEY,
    period TEXT NOT NULL,
    review_date TEXT NOT NULL,
    title TEXT NOT NULL,
    buy_count INTEGER DEFAULT 0,
    sell_count INTEGER DEFAULT 0,
    total_profit REAL DEFAULT 0,
    summary TEXT NOT NULL,
    improvements TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

## 🔧 配置说明

### 环境变量配置
```env
# 应用环境
APP_ENV=development
HTTP_PORT=8080
HTTP_READ_TIMEOUT=15s
HTTP_WRITE_TIMEOUT=15s
HTTP_IDLE_TIMEOUT=60s
TRUST_PROXY=false

# 数据库配置
SQLITE_PATH=data/app.db

# 认证配置
JWT_SECRET=please_change_me
JWT_EXPIRE_MINUTES=1440

# 文件上传配置
UPLOAD_DIR=uploads
MAX_UPLOAD_SIZE_BYTES=5368709120
CHUNK_SIZE_BYTES=2097152
UPLOAD_SESSION_TTL=24h

# 微信配置
WECHAT_APP_ID=
WECHAT_APP_SECRET=
```

### 路由配置
- API路由前缀: `/api`
- 静态文件路由: `/static`
- 前端页面路由: `/`
- 文件上传路由: `/api/upload`

### 中间件配置
- 认证中间件 (可选)
- 响应统一处理
- 错误处理
- 跨域支持

## 📡 API接口

### 基础URL
```
http://localhost:8080/api
```

### 股票管理接口

#### 获取股票列表
```http
GET /api/stocks/getList
```
查询参数:
- `keyword`: 搜索关键词
- `region`: 地区筛选
- `category`: 分类筛选
- `page`: 页码
- `pageSize`: 每页数量

#### 获取股票详情
```http
GET /api/stocks/getDetail/:id
```

#### 创建股票
```http
POST /api/stocks/create
```
请求体:
```json
{
  "code": "000001",
  "name": "平安银行",
  "region": "china",
  "currency": "CNY",
  "category": "main_board",
  "enabled": true,
  "remark": "备注信息"
}
```

#### 更新股票
```http
PUT /api/stocks/update/:id
```

#### 删除股票
```http
DELETE /api/stocks/delete/:id
```

### 文件上传接口

#### 初始化上传
```http
POST /api/upload/init
```
请求体:
```json
{
  "fileId": "unique_file_id",
  "fileName": "example.pdf",
  "fileSize": 1024000,
  "chunkSize": 2097152
}
```

#### 上传分片
```http
POST /api/upload/chunk
```
表单数据:
- `fileId`: 文件ID
- `chunkIndex`: 分片索引
- `file`: 分片文件

#### 完成上传
```http
POST /api/upload/complete
```
请求体:
```json
{
  "fileId": "unique_file_id"
}
```

#### 获取上传进度
```http
GET /api/upload/progress/:fileId
```

### 统一响应格式

#### 成功响应
```json
{
  "code": 200,
  "message": "success",
  "data": { ... }
}
```

#### 错误响应
```json
{
  "code": 400,
  "message": "错误信息",
  "data": null
}
```

## 🎨 架构设计

### 分层架构
- **Handler层**: 处理HTTP请求和响应
- **Service层**: 业务逻辑处理
- **Repository层**: 数据访问层
- **Model层**: 数据模型定义

### 模块化设计
- 按业务功能划分模块
- 每个模块独立的路由、处理器、服务
- 统一的错误处理和响应格式

### 数据库设计
- SQLite轻量级数据库
- 自动表结构迁移
- 支持事务操作
- 数据完整性约束

## 📱 功能特性

### 高性能
- Gin高性能Web框架
- SQLite快速数据访问
- 分片上传大文件支持
- 内存优化和连接池

### 可扩展性
- 模块化架构设计
- 插件式中间件
- 配置化管理
- 微服务友好

### 安全性
- JWT认证支持
- 文件上传安全验证
- SQL注入防护
- 跨域请求控制

### 开发体验
- 热重载开发
- 统一错误处理
- 详细日志记录
- API文档自动生成

## 🔄 开发流程

### 代码规范
- Go标准代码格式
- 统一命名规范
- 错误处理规范
- 注释文档要求

### 模块开发
- 按功能划分模块
- 统一的接口设计
- 完整的测试覆盖
- 文档同步更新

### 数据库操作
- 使用Repository模式
- 事务处理规范
- 数据验证和约束
- 迁移脚本管理

## 📈 性能优化

### 数据库优化
- 索引优化
- 查询语句优化
- 连接池配置
- 事务管理

### 内存优化
- 对象池复用
- 内存泄漏检测
- 垃圾回收优化
- 缓存策略

### 网络优化
- HTTP/2支持
- 压缩传输
- 分片上传
- 连接复用

## 🚀 部署说明

### 生产构建
```bash
# 编译生产版本
go build -ldflags "-s -w" -o server main.go

# 设置环境变量
export APP_ENV=production
export HTTP_PORT=8080
export SQLITE_PATH=/data/app.db

# 启动服务
./server
```

### Docker部署
```dockerfile
FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o server main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/server .
COPY --from=builder /app/config ./config
EXPOSE 8080
CMD ["./server"]
```

### 环境变量
生产环境需要配置以下环境变量:
- `APP_ENV=production`
- `HTTP_PORT=8080`
- `SQLITE_PATH=/data/app.db`
- `JWT_SECRET=your_secret_key`

## 🤝 贡献指南

### 开发规范
1. 遵循Go代码规范
2. 编写完整的单元测试
3. 更新相关文档
4. 提交前运行测试

### 提交规范
- 使用清晰的提交信息
- 遵循约定式提交规范
- 包含必要的测试
- 更新CHANGELOG

## 📄 许可证

ISC License

## 📞 联系方式

如有问题或建议，请通过以下方式联系：
- 项目Issues
- 邮箱联系
- 技术交流群

---

**注意**: 本项目为股票分析系统后端服务，仅供学习和研究使用，不构成投资建议。投资有风险，入市需谨慎。