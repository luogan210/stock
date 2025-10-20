package handler

import (
	"go-demo/service"

	"github.com/gin-gonic/gin"
)

// ExampleHandler 示例处理器，展示如何使用统一响应方法
type ExampleHandler struct {
	*BaseHandler
	exampleService *service.HelloService // 这里用 HelloService 作为示例
}

// NewExampleHandler 创建示例处理器实例
func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{
		BaseHandler:    NewBaseHandler(),
		exampleService: service.NewHelloService(),
	}
}

// ExampleSuccess 示例：成功响应
func (h *ExampleHandler) ExampleSuccess(c *gin.Context) {
	data := h.exampleService.GetHelloMessage()

	// 使用基类的统一成功响应方法
	h.Success(c, data)
}

// ExampleSuccessWithMessage 示例：带自定义消息的成功响应
func (h *ExampleHandler) ExampleSuccessWithMessage(c *gin.Context) {
	data := map[string]interface{}{
		"message": "这是一个自定义消息的示例",
		"data":    "示例数据",
	}

	// 使用基类的带消息成功响应方法
	h.SuccessWithMessage(c, data, "操作成功完成")
}

// ExampleParamError 示例：参数错误响应
func (h *ExampleHandler) ExampleParamError(c *gin.Context) {
	// 模拟参数验证失败
	h.ParamError(c, "用户名不能为空")
}

// ExampleServerError 示例：服务器错误响应
func (h *ExampleHandler) ExampleServerError(c *gin.Context) {
	// 模拟服务器内部错误
	h.ServerError(c, "数据库连接失败")
}

// ExampleUnauthorizedError 示例：未授权错误响应
func (h *ExampleHandler) ExampleUnauthorizedError(c *gin.Context) {
	// 模拟未授权访问
	h.UnauthorizedError(c, "请先登录")
}

// ExampleNotFoundError 示例：资源未找到错误响应
func (h *ExampleHandler) ExampleNotFoundError(c *gin.Context) {
	// 模拟资源未找到
	h.NotFoundError(c, "用户不存在")
}

// ExampleWithMiddleware 示例：使用中间件自动包装响应
func (h *ExampleHandler) ExampleWithMiddleware(c *gin.Context) {
	// 使用中间件自动包装响应
	data := map[string]interface{}{
		"message":   "使用中间件自动包装的响应",
		"timestamp": "2024-01-01 12:00:00",
	}

	// 直接使用Success方法，更简洁
	h.Success(c, data)
}

// RegisterExampleRoutes 注册示例路由
func RegisterExampleRoutes(rg *gin.RouterGroup) {
	exampleHandler := NewExampleHandler()

	// 示例路由
	examples := rg.Group("/examples")
	{
		examples.GET("/success", exampleHandler.ExampleSuccess)                    // 成功响应示例
		examples.GET("/success-message", exampleHandler.ExampleSuccessWithMessage) // 带消息的成功响应示例
		examples.GET("/param-error", exampleHandler.ExampleParamError)             // 参数错误示例
		examples.GET("/server-error", exampleHandler.ExampleServerError)           // 服务器错误示例
		examples.GET("/unauthorized", exampleHandler.ExampleUnauthorizedError)     // 未授权错误示例
		examples.GET("/not-found", exampleHandler.ExampleNotFoundError)            // 资源未找到示例
		examples.GET("/middleware", exampleHandler.ExampleWithMiddleware)          // 中间件自动包装示例
	}
}
