package middleware

import (
	"github.com/gin-gonic/gin"
)

// ResponseWrapper 响应包装器中间件
func ResponseWrapper() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 继续处理请求
		c.Next()

		// 检查是否有错误
		if len(c.Errors) > 0 {
			// 有错误，返回错误响应
			c.JSON(200, gin.H{
				"code":    500,
				"message": c.Errors.String(),
				"data":    nil,
			})
		} else {
			// 没有错误，检查是否有数据需要返回
			if data, exists := c.Get("response_data"); exists {
				c.JSON(200, gin.H{
					"code":    200,
					"message": "success",
					"data":    data,
				})
			}
		}
	}
}

// SetResponseData 设置响应数据的辅助函数
func SetResponseData(c *gin.Context, data interface{}) {
	c.Set("response_data", data)
}
