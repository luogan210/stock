# 微信小程序 urlLink API 文档

## 概述

本API提供微信小程序urlLink生成功能，支持生成带查询参数的小程序链接，并支持设置过期时间。

## API 接口

### 1. 生成 urlLink

**接口地址：** `POST /api/wechat/url-link`

**请求参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| path | string | 是 | 小程序页面路径，如：pages/index/index |
| query | string | 否 | 查询参数，如：id=123&type=product |
| is_expire | boolean | 否 | 是否设置过期时间，默认false |
| expire_type | int | 否 | 过期类型：0-指定时间失效，1-指定天数失效 |
| expire_time | int64 | 否 | 过期时间戳（当expire_type为0时使用） |
| expire_interval | int | 否 | 过期天数（当expire_type为1时使用） |

**过期类型说明：**
- `expire_type = 0`：指定时间失效，需要提供 `expire_time`（Unix时间戳）
- `expire_type = 1`：指定天数失效，需要提供 `expire_interval`（天数）

**请求示例：**

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

**响应格式：**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "url_link": "https://wxaurl.cn/xxxxx",
    "expire_time": 1704067200
  }
}
```

### 2. 获取 urlLink 信息

**接口地址：** `GET /api/wechat/url-link/{id}`

**请求参数：**

| 参数名 | 类型 | 必填 | 说明 |
|--------|------|------|------|
| id | uint | 是 | urlLink记录ID |

**请求示例：**

```bash
curl -X GET http://localhost:8080/api/wechat/url-link/1
```

**响应格式：**

```json
{
  "code": 200,
  "message": "success",
  "data": {
    "id": 1,
    "path": "pages/index/index",
    "query": "id=123&type=product",
    "url_link": "https://wxaurl.cn/xxxxx",
    "expire_time": 1704067200,
    "created_at": 1704067200,
    "status": 1
  }
}
```

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 成功 |
| 400 | 请求参数错误 |
| 401 | 未授权 |
| 500 | 服务器内部错误 |

## 注意事项

1. **查询参数处理**：系统会自动处理查询参数中的特殊字符，确保`&`等字符不会被转义
2. **过期时间**：如果不设置过期时间，urlLink将永久有效
3. **时间格式**：`expire_time`使用Unix时间戳格式
4. **配置要求**：需要在环境变量中设置正确的微信小程序AppID和AppSecret

## 环境变量配置

```bash
# 微信小程序配置
WECHAT_APP_ID=your_app_id
WECHAT_APP_SECRET=your_app_secret
```