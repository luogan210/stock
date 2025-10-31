package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`    // 业务状态码
	Message string      `json:"message"` // 响应消息
	Data    interface{} `json:"data"`    // 响应数据
}

// 业务状态码常量
const (
	CodeSuccess  = 0   // 成功
	CodeError    = 500 // 错误
	CodeInvalid  = 400 // 参数无效
	CodeNotFound = 404 // 未找到
)

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

// getRequestBody 获取请求体内容（用于日志记录）
func getRequestBody(c *gin.Context) string {
	// 优先从上下文中获取保存的请求体（由中间件保存）
	if bodyBytes, exists := c.Get("request_body"); exists {
		if body, ok := bodyBytes.([]byte); ok && len(body) > 0 {
			// 尝试格式化 JSON
			var jsonData interface{}
			if err := json.Unmarshal(body, &jsonData); err == nil {
				formatted, _ := json.MarshalIndent(jsonData, "", "  ")
				return string(formatted)
			}

			// 如果不是 JSON，直接返回字符串（限制长度避免日志过大）
			if len(body) > 1000 {
				return string(body[:1000]) + "... (truncated)"
			}
			return string(body)
		}
	}

	// 如果上下文中没有，尝试从请求中读取（可能已被消耗）
	if c.Request.Body == nil {
		return ""
	}

	// 读取请求体
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return ""
	}

	// 恢复请求体，因为 ReadAll 会消耗掉
	c.Request.Body = io.NopCloser(bytes.NewBuffer(body))

	// 尝试格式化 JSON
	var jsonData interface{}
	if err := json.Unmarshal(body, &jsonData); err == nil {
		formatted, _ := json.MarshalIndent(jsonData, "", "  ")
		return string(formatted)
	}

	// 如果不是 JSON，直接返回字符串（限制长度避免日志过大）
	if len(body) > 1000 {
		return string(body[:1000]) + "... (truncated)"
	}
	return string(body)
}

// getRequestInfo 获取请求详细信息
func getRequestInfo(c *gin.Context) map[string]interface{} {
	info := map[string]interface{}{
		"method":    c.Request.Method,
		"path":      c.Request.URL.Path,
		"query":     c.Request.URL.RawQuery,
		"clientIP":  c.ClientIP(),
		"userAgent": c.Request.UserAgent(),
	}

	// 获取路径参数
	if len(c.Params) > 0 {
		params := make(map[string]string)
		for _, param := range c.Params {
			params[param.Key] = param.Value
		}
		info["pathParams"] = params
	}

	// 获取查询参数
	if len(c.Request.URL.Query()) > 0 {
		queryParams := make(map[string][]string)
		for k, v := range c.Request.URL.Query() {
			queryParams[k] = v
		}
		info["queryParams"] = queryParams
	}

	// 对于 POST/PUT/PATCH 请求，获取请求体
	if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "PATCH" {
		body := getRequestBody(c)
		if body != "" {
			info["body"] = body
		}
	}

	return info
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
	// 记录错误请求到日志
	requestInfo := getRequestInfo(c)

	// 构建详细的错误日志
	utils.LogError("========== 请求处理失败 ==========")
	utils.LogError("错误消息: %s", message)
	utils.LogError("业务状态码: %d", code)
	utils.LogError("HTTP方法: %s", requestInfo["method"])
	utils.LogError("请求路径: %s", requestInfo["path"])
	utils.LogError("客户端IP: %s", requestInfo["clientIP"])

	// 记录 User-Agent
	if userAgent, ok := requestInfo["userAgent"].(string); ok && userAgent != "" {
		utils.LogError("User-Agent: %s", userAgent)
	}

	// 记录查询参数
	if queryParams, ok := requestInfo["queryParams"].(map[string][]string); ok && len(queryParams) > 0 {
		queryJSON, _ := json.MarshalIndent(queryParams, "", "  ")
		utils.LogError("查询参数:\n%s", string(queryJSON))
	}

	// 记录路径参数
	if pathParams, ok := requestInfo["pathParams"].(map[string]string); ok && len(pathParams) > 0 {
		pathJSON, _ := json.MarshalIndent(pathParams, "", "  ")
		utils.LogError("路径参数:\n%s", string(pathJSON))
	}

	// 记录请求体（如果是 POST/PUT/PATCH）
	if body, ok := requestInfo["body"].(string); ok && body != "" {
		utils.LogError("请求体:\n%s", body)
	}

	utils.LogError("====================================")

	// 返回错误响应
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    nil,
	})
}

// SuccessWithMessage 带自定义消息的成功响应
func SuccessWithMessage(c *gin.Context, data interface{}, message string) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: message,
		Data:    data,
	})
}
