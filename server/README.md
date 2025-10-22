# Trade Server (Gin + SQLite)

最小后端：仅保留「股票管理」与「文件上传」相关接口，使用 SQLite 持久化。

## 目录结构（精简后）

```
server/
├── main.go
├── go.mod / go.sum
├── config/
│   ├── app.env         # 环境变量（含 SQLITE_PATH）
│   └── config.go       # 加载配置
├── router/
│   ├── router.go       # 路由与中间件
│   └── frontend.go     # 静态/示例页面（可选）
├── handler/
│   ├── upload.go       # 文件上传接口
│   ├── stock.go        # 股票接口（GET/POST）
│   ├── routes.go       # 注册路由（仅上传与股票）
│   ├── response.go
│   └── base.go
├── middleware/
│   └── response.go     # 统一响应封装（按需）
└── storage/
    └── sqlite.go       # SQLite 初始化与迁移（stocks, plans）
```

## 运行

```bash
# 1) 配置
cp server/config/app.env server/config/app.env.local  # 可选

# 确认端口/DB 文件路径
# HTTP_PORT=8080
# SQLITE_PATH=data/app.db

# 2) 启动
cd server
go run main.go
```

前端开发环境已配置 Vite 代理：`/api -> http://localhost:8080`。

## 环境变量（关键）

```
APP_ENV=development
HTTP_PORT=8080
SQLITE_PATH=data/app.db
```

更多可选项见 `server/config/app.env`。

## SQLite

- 启动时自动创建 `SQLITE_PATH` 指定的数据库文件与目录
- 自动迁移最小表结构：
  - `stocks`（代码、名称、地区、货币、分类、状态、备注等）
  - `plans`（保留结构，后续可用；当前未对外暴露接口）

## API（最小集）

Base URL: `http://localhost:8080/api`

### 股票管理接口

- **获取股票列表**
  - `GET /stocks/getList`
  - 查询参数：`keyword, region, category`
  - 返回：`{"list": [...]}`

- **获取股票详情**
  - `GET /stocks/getDetail/:id`
  - 返回：股票详细信息

- **创建股票**
  - `POST /stocks/create`
  - 请求体：股票数据
  - 返回：`{"id": "new_id"}`

- **更新股票**
  - `PUT /stocks/update/:id`
  - 请求体：更新的股票数据
  - 返回：`{"id": "updated_id"}`

- **删除股票**
  - `DELETE /stocks/delete/:id`
  - 返回：`{"id": "deleted_id"}`

### 股票数据模型

```json
{
  "id": "string",
  "code": "string",
  "name": "string", 
  "region": "string",
  "currency": "string",
  "category": "string",
  "enabled": "boolean",
  "remark": "string"
}
```

### 文件上传接口
- 参考 `handler/upload.go` 暴露的路由（前端可直接调用）

### 统一响应格式

成功响应：
```json
{
  "code": 200,
  "message": "success",
  "data": { ... }
}
```

错误响应：
```json
{
  "code": 400,
  "message": "错误信息",
  "data": null
}
```

## 前端联调

- Vite 已在 `web/vite.config.js` 配置：
  - `server.proxy['/api'] -> http://localhost:8080`
- 前端可直接请求 `/api/...`，无需关心后端端口

## 功能特性

### 已实现功能
- ✅ 股票管理完整CRUD操作
- ✅ 股票数据搜索和过滤
- ✅ 统一响应格式
- ✅ 数据库自动迁移
- ✅ 前端代理配置

### 后续计划（可选）
- 分页与排序功能
- 批量操作接口
- 数据导入导出
- 为上传接口增加上传记录持久化表
- 按需恢复鉴权中间件（当前开发阶段可跳过）