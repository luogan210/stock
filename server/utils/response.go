package utils

import (
	"net/http"

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
	CodeSuccess = 0   // 成功
	CodeError   = 500 // 错误
	CodeInvalid = 400 // 参数无效
)

// Success 成功响应
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

// Error 错误响应
func Error(c *gin.Context, code int, message string) {
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

// ParamError 参数错误
func ParamError(c *gin.Context, message string) {
	Error(c, CodeInvalid, message)
}

// ServerError 服务器错误
func ServerError(c *gin.Context, message string) {
	Error(c, CodeError, message)
}

// UnauthorizedError 未授权错误
func UnauthorizedError(c *gin.Context, message string) {
	Error(c, 401, message)
}

// NotFoundError 未找到错误
func NotFoundError(c *gin.Context, message string) {
	Error(c, 404, message)
}
