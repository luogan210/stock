package middleware

import (
	"server/config"
	"strings"

	"github.com/gin-gonic/gin"
)

// UserContextKey 用户上下文键
const UserContextKey = "user"

// AuthMiddleware 认证中间件
func AuthMiddleware() gin.HandlerFunc {
	cfg := config.Load()
	_ = cfg // placeholder to show config usage; token validation below remains mock
	return func(c *gin.Context) {
		// 从请求头获取token
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			// 如果没有token，设置空用户信息
			c.Set(UserContextKey, nil)
			c.Next()
			return
		}

		// 检查token格式
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.Set(UserContextKey, nil)
			c.Next()
			return
		}

		token := tokenParts[1]

		// 验证token并获取用户信息
		userInfo := validateToken(token)
		c.Set(UserContextKey, userInfo)

		c.Next()
	}
}

// validateToken 验证token并返回用户信息
func validateToken(token string) any {
	// 这里应该从数据库或缓存中验证token
	// 暂时使用简单的验证逻辑
	if token == "mock_token_123456" {
		return map[string]any{
			"id":       1,
			"username": "admin",
			"email":    "admin@example.com",
			"nickname": "管理员",
			"status":   1,
		}
	}

	return nil
}

// GetCurrentUser 从Context中获取当前用户信息
func GetCurrentUser(c *gin.Context) any {
	user, exists := c.Get(UserContextKey)
	if !exists {
		return nil
	}
	return user
}
