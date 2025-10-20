package handler

import (
	"go-demo/service"

	"github.com/gin-gonic/gin"
)

// HelloHandler 处理Hello相关的请求
type HelloHandler struct {
	*BaseHandler
	helloService *service.HelloService
}

// NewHelloHandler 创建HelloHandler实例
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{
		BaseHandler:  NewBaseHandler(),
		helloService: service.NewHelloService(),
	}
}

// Hello 返回Hello World
func (h *HelloHandler) Hello(c *gin.Context) {
	data := h.helloService.GetHelloMessage()

	// 使用BaseHandler的方法获取当前用户信息
	user := h.GetCurrentUser(c)

	if user != nil {
		data["user"] = map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"nickname": user.Nickname,
			"email":    user.Email,
		}
		data["message"] = "Hello, " + user.Nickname + "!"
	}

	// 使用基类的统一响应方法
	h.Success(c, data)
}

// RegisterHelloRoutes 注册hello相关路由
func RegisterHelloRoutes(rg *gin.RouterGroup) {
	helloHandler := NewHelloHandler()
	rg.GET("/hello", helloHandler.Hello)
}
