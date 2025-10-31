package middleware

import (
	"bytes"
	"io"
	"server/utils"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger 请求日志中间件
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery
		method := c.Request.Method
		clientIP := c.ClientIP()

		// 保存请求体（用于错误日志记录）
		var bodyBytes []byte
		if c.Request.Body != nil && (method == "POST" || method == "PUT" || method == "PATCH") {
			bodyBytes, _ = io.ReadAll(c.Request.Body)
			// 恢复请求体，以便后续处理可以使用
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			// 将请求体保存到上下文中，供错误处理使用
			c.Set("request_body", bodyBytes)
		}

		// 处理请求
		c.Next()

		// 结束时间
		latency := time.Since(start)
		statusCode := c.Writer.Status()

		// 构建日志信息
		if raw != "" {
			path = path + "?" + raw
		}

		// 根据状态码选择日志级别
		if statusCode >= 500 {
			utils.LogError("[%s] %s | %d | %v | %s",
				method, path, statusCode, latency, clientIP)
		} else if statusCode >= 400 {
			utils.LogWarning("[%s] %s | %d | %v | %s",
				method, path, statusCode, latency, clientIP)
		} else {
			utils.LogInfo("[%s] %s | %d | %v | %s",
				method, path, statusCode, latency, clientIP)
		}
	}
}
