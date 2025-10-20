# go-demo

一个基于 Gin 框架的 Go 服务端示例项目，支持热启动开发。

## 项目结构

```
go-demo/
├── main.go              # 程序入口
├── go.mod               # 依赖管理
├── go.sum               # 依赖校验
├── README.md            # 项目说明
├── .air.toml            # 热启动配置
├── .gitignore           # Git忽略文件
├── package.json         # 开发脚本
├── scripts/             # 开发脚本目录
│   ├── dev.sh          # Linux/Mac开发脚本
│   └── dev.bat         # Windows开发脚本
├── router/              # 路由层
│   └── router.go        # 路由配置
├── handler/             # 控制器层
│   ├── hello.go         # Hello相关处理函数
│   ├── user.go          # 用户相关处理函数
│   ├── wechat.go        # 微信小程序相关处理函数
│   ├── response.go      # 统一响应处理
│   └── routes.go        # 路由注册
├── service/             # 业务逻辑层
│   ├── hello_service.go # Hello相关业务逻辑
│   ├── user_service.go  # 用户相关业务逻辑
│   └── wechat_service.go # 微信小程序业务逻辑
├── model/               # 数据模型层
│   ├── user.go          # 用户数据模型
│   └── wechat.go        # 微信小程序数据模型
├── config/              # 配置层
│   └── wechat.go        # 微信小程序配置
├── middleware/          # 中间件层
│   ├── auth.go          # 认证中间件
│   └── response.go      # 响应中间件
├── frontend/            # 前端文件夹
│   ├── static/         # 静态资源
│   │   ├── css/       # CSS样式
│   │   └── js/        # JavaScript
│   ├── templates/      # HTML模板
│   └── README.md       # 前端说明
├── examples/            # 示例代码
│   └── wechat_example.go # 微信小程序使用示例
├── WECHAT_API.md        # 微信小程序API文档
└── ARCHITECTURE_GUIDE.md # 架构指南
```

## 功能特性

### 1. 基础功能
- ✅ 统一响应格式
- ✅ 用户管理（注册、登录、查询）
- ✅ 认证中间件
- ✅ 热启动开发

### 2. 微信小程序功能
- ✅ 生成urlLink（支持过期时间设置）
- ✅ 获取urlLink信息
- ✅ 支持永久、指定时间、指定天数过期
- ✅ 完整的API文档和示例

### 3. 前端功能
- ✅ 响应式Web界面
- ✅ 用户登录注册页面
- ✅ 微信功能操作界面
- ✅ 现代化UI设计
- ✅ 完整的JavaScript交互

## 快速开始

### 1. 环境准备

```bash
# 检查Go版本 (需要1.20+)
go version

# 安装air热启动工具
go install github.com/air-verse/air@latest

# 设置PATH (如果air命令找不到)
export PATH="$PATH:$(go env GOPATH)/bin"
```

### 2. 微信小程序配置

设置环境变量或修改配置文件：

```bash
# 方式1：环境变量
export WECHAT_APP_ID="your_app_id"
export WECHAT_APP_SECRET="your_app_secret"

# 方式2：修改配置文件
# 编辑 config/wechat.go 文件
```

### 3. 开发模式（推荐）

#### 方式一：使用npm脚本（跨平台）
```bash
# 安装依赖并启动开发模式
npm run dev

# 或者使用平台特定脚本
npm run dev:linux    # Linux/Mac
npm run dev:windows  # Windows
```

#### 方式二：使用开发脚本
```bash
# Linux/Mac
./scripts/dev.sh

# Windows
scripts\dev.bat
```

#### 方式三：直接使用air
```bash
air
```

### 4. 生产模式

```bash
# 构建项目
npm run build

# 运行服务
npm start
```

## API 接口

### 基础接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | `/api/hello` | Hello World |
| POST | `/api/users/` | 创建用户 |
| GET | `/api/users/:id` | 获取用户信息 |
| PUT | `/api/users/:id` | 更新用户信息 |
| POST | `/api/users/login` | 用户登录 |
| GET | `/api/users/current` | 获取当前用户 |

### 微信小程序接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | `/api/wechat/url-link` | 生成urlLink |
| GET | `/api/wechat/url-link/:id` | 获取urlLink信息 |

详细API文档请参考：[WECHAT_API.md](./WECHAT_API.md)

## 使用示例

### 微信小程序urlLink生成

```bash
# 1. 生成不设置过期时间的urlLink
curl -X POST http://localhost:8080/api/wechat/url-link \
  -H "Content-Type: application/json" \
  -d '{
    "path": "pages/index/index",
    "query": "id=123&type=product"
  }'

# 2. 生成指定时间过期的urlLink
curl -X POST http://localhost:8080/api/wechat/url-link \
  -H "Content-Type: application/json" \
  -d '{
    "path": "pages/product/detail",
    "query": "id=456",
    "is_expire": true,
    "expire_type": 0,
    "expire_time": 1704067200
  }'

# 3. 生成指定天数后过期的urlLink
curl -X POST http://localhost:8080/api/wechat/url-link \
  -H "Content-Type: application/json" \
  -d '{
    "path": "pages/activity/share",
    "query": "activity_id=789",
    "is_expire": true,
    "expire_type": 1,
    "expire_interval": 30
  }'
```

### 运行示例代码

```bash
# 运行微信小程序示例
go run examples/wechat_example.go
```

## 热启动最佳实践

### 1. 配置文件优化

`.air.toml` 配置文件已优化：
- **智能监听**：只监听相关文件类型（`.go`, `.yaml`, `.json`等）
- **排除测试**：自动排除 `_test.go` 文件，避免测试文件修改触发重启
- **延迟设置**：1秒延迟避免频繁重启
- **日志管理**：构建日志保存到 `tmp/air.log`
- **清理机制**：退出时自动清理临时文件

### 2. 开发脚本

提供了完整的开发脚本：
- **环境检查**：自动检查Go和air安装
- **依赖管理**：自动安装和更新依赖
- **测试集成**：启动前自动运行测试
- **清理功能**：自动清理临时文件
- **彩色输出**：友好的彩色日志输出

### 3. 开发工作流

```bash
# 1. 启动开发模式
npm run dev

# 2. 修改代码（自动重启）

# 3. 运行测试
npm test

# 4. 代码格式化
npm run fmt

# 5. 代码检查
npm run vet

# 6. 构建生产版本
npm run build:prod
```

### 4. 性能优化

- **构建优化**：生产构建使用 `-ldflags='-s -w'` 减小二进制文件大小
- **测试优化**：支持并发测试、覆盖率测试、性能测试
- **内存优化**：支持内存和CPU性能分析

### 5. 调试技巧

```bash
# 详细测试输出
npm run test:verbose

# 测试覆盖率
npm run test:coverage

# 竞态检测
npm run test:race

# 性能基准测试
npm run bench

# 性能分析
npm run profile
```

## 访问地址

### 前端界面
- **首页**: http://localhost:8080/
- **登录页面**: http://localhost:8080/login
- **注册页面**: http://localhost:8080/register

### API接口

所有后端接口都以 `/api` 开头：

### Hello接口
- `GET /api/` - Hello World接口（会根据用户登录状态返回不同信息）

### 用户接口
- `POST /api/users/` - 创建用户
- `GET /api/users/:id` - 获取用户信息
- `PUT /api/users/:id` - 更新用户信息
- `POST /api/users/login` - 用户登录
- `GET /api/users/current` - 获取当前用户信息
- `GET /api/users/profile` - 获取个人资料

### 响应格式

所有接口都返回统一的JSON格式：

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "message": "Hello, World!",
    "service": "hello",
    "user": {
      "id": 1,
      "username": "admin",
      "nickname": "管理员",
      "email": "admin@example.com"
    }
  }
}
```

### 状态码说明

- `200`: 成功
- `400`: 参数无效
- `401`: 未授权访问
- `500`: 服务器错误

### 认证说明

项目支持基于Token的认证机制：

1. **登录获取Token**：
```bash
POST /api/users/login
Content-Type: application/json

{
  "username": "admin",
  "password": "123456"
}
```

2. **使用Token访问需要认证的接口**：
```bash
GET /api/users/profile
Authorization: Bearer mock_token_123456
```

3. **获取当前用户信息**：
```bash
GET /api/users/current
Authorization: Bearer mock_token_123456
```

### 用户接口示例

#### 创建用户
```bash
POST /api/users/
Content-Type: application/json

{
  "username": "test_user",
  "email": "test@example.com",
  "password": "123456",
  "nickname": "测试用户"
}
```

#### 用户登录
```bash
POST /api/users/login
Content-Type: application/json

{
  "username": "admin",
  "password": "123456"
}
```

#### 获取用户信息
```bash
GET /api/users/1
```

#### 更新用户信息
```bash
PUT /api/users/1
Content-Type: application/json

{
  "nickname": "新昵称",
  "avatar": "https://example.com/avatar.jpg"
}
```

## 开发命令

### 基础命令
```bash
npm run dev          # 启动开发模式
npm run build        # 构建项目
npm run start        # 运行生产版本
npm run test         # 运行测试
npm run clean        # 清理临时文件
```

### 开发工具
```bash
npm run fmt          # 代码格式化
npm run vet          # 代码检查
npm run lint         # 代码静态分析
npm run deps         # 安装依赖
npm run deps:update  # 更新依赖
```

### 测试相关
```bash
npm run test:verbose   # 详细测试输出
npm run test:coverage  # 测试覆盖率
npm run test:race      # 竞态检测
npm run bench          # 性能基准测试
npm run profile        # 性能分析
```

### 文档
```bash
npm run docs          # 启动Go文档服务器 (http://localhost:6060)
```

## 架构说明

- **main.go**: 程序入口，启动服务器
- **router/**: 路由配置，负责URL映射
- **handler/**: 控制器层，处理HTTP请求和响应
- **service/**: 业务逻辑层，处理具体的业务逻辑
- **model/**: 数据模型层，定义数据结构
- **middleware/**: 中间件层，处理认证等通用逻辑

## 开发工具

- **air**: 热启动开发工具，监听文件变化自动重启服务
- **配置**: `.air.toml` 文件配置了监听规则和构建选项
- **脚本**: 提供了跨平台的开发脚本
- **测试**: 集成了完整的测试工具链
- **性能**: 支持性能分析和基准测试

## 最佳实践

1. **开发流程**：使用 `npm run dev` 启动开发模式
2. **代码质量**：定期运行 `npm run fmt` 和 `npm run vet`
3. **测试驱动**：编写测试用例，使用 `npm test` 验证
4. **性能监控**：使用 `npm run profile` 分析性能瓶颈
5. **依赖管理**：使用 `npm run deps` 管理依赖
6. **版本控制**：`.gitignore` 已配置，避免提交临时文件 

## 环境变量配置（最佳实践）

支持通过系统环境变量和 `.env` 文件配置参数（开发环境建议使用 `.env`）。

示例 `.env`（请复制为 `.env` 并按需修改）：

```env
# App
APP_ENV=development            # development|production|test
HTTP_PORT=8080
HTTP_READ_TIMEOUT=15s
HTTP_WRITE_TIMEOUT=15s
HTTP_IDLE_TIMEOUT=60s
TRUST_PROXY=false

# Auth
JWT_SECRET=please_change_me
JWT_EXPIRE_MINUTES=1440        # 24h

# Uploads
UPLOAD_DIR=uploads
MAX_UPLOAD_SIZE_BYTES=5368709120  # 5GB
CHUNK_SIZE_BYTES=2097152          # 2MB
UPLOAD_SESSION_TTL=24h

# Wechat
WECHAT_APP_ID=your_app_id
WECHAT_APP_SECRET=your_app_secret
```

加载顺序：系统环境变量优先于 `.env`。

## 启动

```bash
# 开发模式（支持 .env）
go run main.go

# 或生产环境
APP_ENV=production HTTP_PORT=80 go run main.go
``` 