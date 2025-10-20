package handler

import (
	"go-demo/middleware"
	"go-demo/model"

	"github.com/gin-gonic/gin"
)

// BaseHandler 基础处理器，提供统一的响应方法
type BaseHandler struct{}

// NewBaseHandler 创建基础处理器实例
func NewBaseHandler() *BaseHandler {
	return &BaseHandler{}
}

// GetCurrentUser 获取当前用户信息
func (h *BaseHandler) GetCurrentUser(c *gin.Context) *model.UserInfo {
	return middleware.GetCurrentUser(c)
}

// GetCurrentUserOrNil 获取当前用户信息，如果未登录返回nil
func (h *BaseHandler) GetCurrentUserOrNil(c *gin.Context) *model.UserInfo {
	return middleware.GetCurrentUser(c)
}

// GetCurrentUserRequired 获取当前用户信息，如果未登录返回错误
func (h *BaseHandler) GetCurrentUserRequired(c *gin.Context) (*model.UserInfo, bool) {
	user := middleware.GetCurrentUser(c)
	if user == nil {
		h.UnauthorizedError(c, "请先登录")
		return nil, false
	}
	return user, true
}

// Success 统一成功响应
func (h *BaseHandler) Success(c *gin.Context, data interface{}) {
	Success(c, data)
}

// SuccessWithMessage 带自定义消息的成功响应
func (h *BaseHandler) SuccessWithMessage(c *gin.Context, data interface{}, message string) {
	SuccessWithMessage(c, data, message)
}

// Error 统一错误响应
func (h *BaseHandler) Error(c *gin.Context, code int, message string) {
	Error(c, code, message)
}

// ErrorWithData 带数据的错误响应
func (h *BaseHandler) ErrorWithData(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(200, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// ParamError 参数错误响应
func (h *BaseHandler) ParamError(c *gin.Context, message string) {
	h.Error(c, CodeInvalid, message)
}

// ServerError 服务器错误响应
func (h *BaseHandler) ServerError(c *gin.Context, message string) {
	h.Error(c, CodeError, message)
}

// UnauthorizedError 未授权错误响应
func (h *BaseHandler) UnauthorizedError(c *gin.Context, message string) {
	h.Error(c, 401, message)
}

// NotFoundError 资源未找到错误响应
func (h *BaseHandler) NotFoundError(c *gin.Context, message string) {
	h.Error(c, 404, message)
}
