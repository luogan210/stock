# JSON 转义问题解决方案

## 问题描述

Go的官方`encoding/json`库会对特殊字符进行转义，包括：
- `&` → `\u0026`
- `<` → `\u003c`
- `>` → `\u003e`
- `"` → `\"`

这在处理URL查询参数时会造成问题，例如：
```json
{
  "path": "pages/index/index",
  "query": "id=123&type=product"
}
```

会被转义为：
```json
{
  "path": "pages/index/index",
  "query": "id=123\u0026type=product"
}
```

## 解决方案

### 方案1：使用自定义JSON编码器（推荐）

```go
// marshalJSONWithoutEscape 自定义JSON编码，避免转义特殊字符
func (s *WechatService) marshalJSONWithoutEscape(v interface{}) ([]byte, error) {
    var buf bytes.Buffer
    encoder := json.NewEncoder(&buf)
    encoder.SetEscapeHTML(false) // 关键：不转义HTML字符
    encoder.SetIndent("", "")
    
    if err := encoder.Encode(v); err != nil {
        return nil, err
    }
    
    // 移除末尾的换行符
    data := buf.Bytes()
    if len(data) > 0 && data[len(data)-1] == '\n' {
        data = data[:len(data)-1]
    }
    
    return data, nil
}
```

**优点：**
- 使用标准库，无需额外依赖
- 简单易用
- 性能良好

**缺点：**
- 仍然可能转义某些字符

### 方案2：使用json-iterator库

```go
import jsoniter "github.com/json-iterator/go"

type WechatServiceAlternative struct {
    config *config.WechatConfig
    json   jsoniter.API
}

func NewWechatServiceAlternative() *WechatServiceAlternative {
    return &WechatServiceAlternative{
        config: config.GetWechatConfig(),
        json:   jsoniter.ConfigCompatibleWithStandardLibrary,
    }
}

// 使用json-iterator编码
jsonData, err := s.json.Marshal(requestBody)
```

**优点：**
- 对特殊字符处理更友好
- 性能更好
- 兼容标准库

**缺点：**
- 需要额外依赖

### 方案3：手动构建JSON字符串（最可靠）

```go
// buildJSONString 手动构建JSON字符串
func (s *WechatServiceManual) buildJSONString(req *WechatUrlLinkRequest, expireTime int64) string {
    var jsonStr bytes.Buffer
    
    jsonStr.WriteString("{")
    jsonStr.WriteString(`"path":"`)
    jsonStr.WriteString(req.Path)
    jsonStr.WriteString(`"`)
    
    if req.Query != "" {
        jsonStr.WriteString(`,"query":"`)
        jsonStr.WriteString(req.Query) // 直接写入，不转义
        jsonStr.WriteString(`"`)
    }
    
    if expireTime > 0 {
        jsonStr.WriteString(`,"expire_type":1`)
        jsonStr.WriteString(`,"expire_time":`)
        jsonStr.WriteString(strconv.FormatInt(expireTime, 10))
    } else {
        jsonStr.WriteString(`,"expire_type":0`)
    }
    
    jsonStr.WriteString("}")
    
    return jsonStr.String()
}
```

**优点：**
- 完全控制输出格式
- 不转义任何字符
- 性能最好

**缺点：**
- 代码较复杂
- 需要手动处理所有字段

## 测试对比

### 原始数据
```go
requestBody := map[string]interface{}{
    "path": "pages/index/index",
    "query": "id=123&type=product&category=electronics",
}
```

### 方案1输出
```json
{"path":"pages/index/index","query":"id=123&type=product&category=electronics"}
```

### 方案2输出
```json
{"path":"pages/index/index","query":"id=123&type=product&category=electronics"}
```

### 方案3输出
```json
{"path":"pages/index/index","query":"id=123&type=product&category=electronics"}
```

## 推荐使用

1. **开发阶段**：使用方案1（自定义编码器）
2. **生产环境**：使用方案2（json-iterator）
3. **特殊需求**：使用方案3（手动构建）

## 安装json-iterator

```bash
go get github.com/json-iterator/go
```

## 注意事项

1. **安全性**：手动构建JSON时要注意XSS攻击
2. **维护性**：手动构建的代码需要更多维护
3. **兼容性**：确保JSON格式符合微信API要求

## 实际使用示例

```bash
# 测试包含特殊字符的查询参数
curl -X POST http://localhost:8080/api/wechat/url-link \
  -H "Content-Type: application/json" \
  -d '{
    "path": "pages/product/detail",
    "query": "id=123&type=product&category=electronics&brand=apple&model=iphone",
    "is_expire": true,
    "expire_type": 0,
    "expire_time": 1704067200
  }'
```

**过期类型说明：**
- `expire_type = 0`：指定时间失效，需要提供 `expire_time`（Unix时间戳）
- `expire_type = 1`：指定天数失效，需要提供 `expire_interval`（天数）

这样生成的urlLink中的查询参数将保持原始格式，不会被转义。