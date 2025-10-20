package handler

import (
	"go-demo/model"
	"go-demo/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserHandler 用户处理函数
type UserHandler struct {
	*BaseHandler
	userService *service.UserService
}

// NewUserHandler 创建用户处理函数实例
func NewUserHandler() *UserHandler {
	return &UserHandler{
		BaseHandler: NewBaseHandler(),
		userService: service.NewUserService(),
	}
}

// CreateUser 创建用户
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req model.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	userInfo, err := h.userService.CreateUser(&req)
	if err != nil {
		h.ServerError(c, err.Error())
		return
	}

	h.Success(c, userInfo)
}

// GetUser 获取用户信息
func (h *UserHandler) GetUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.ParamError(c, "用户ID格式错误")
		return
	}

	userInfo, err := h.userService.GetUserByID(uint(id))
	if err != nil {
		h.ServerError(c, err.Error())
		return
	}

	h.Success(c, userInfo)
}

// UpdateUser 更新用户信息
func (h *UserHandler) UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		h.ParamError(c, "用户ID格式错误")
		return
	}

	var req model.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	userInfo, err := h.userService.UpdateUser(uint(id), &req)
	if err != nil {
		h.ServerError(c, err.Error())
		return
	}

	h.Success(c, userInfo)
}

// Login 用户登录
func (h *UserHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.ParamError(c, "参数错误: "+err.Error())
		return
	}

	loginResp, err := h.userService.Login(&req)
	if err != nil {
		h.ServerError(c, err.Error())
		return
	}

	h.Success(c, loginResp)
}

// GetCurrentUser 获取当前用户信息（从Context中获取）
func (h *UserHandler) GetCurrentUser(c *gin.Context) {
	// 使用BaseHandler的方法获取当前用户信息
	user := h.GetCurrentUserOrNil(c)
	if user == nil {
		h.UnauthorizedError(c, "未登录")
		return
	}

	h.Success(c, user)
}

// GetUserProfile 获取用户个人资料（需要认证）
func (h *UserHandler) GetUserProfile(c *gin.Context) {
	// 使用BaseHandler的方法获取当前用户信息
	user, ok := h.GetCurrentUserRequired(c)
	if !ok {
		return // 错误已经在GetCurrentUserRequired中处理
	}

	// 这里可以添加额外的业务逻辑
	h.Success(c, user)
}

// RegisterUserRoutes 注册用户相关路由
func RegisterUserRoutes(rg *gin.RouterGroup) {
	userHandler := NewUserHandler()

	// 用户相关路由
	users := rg.Group("/users")
	{
		users.POST("/", userHandler.CreateUser)           // 创建用户
		users.GET("/:id", userHandler.GetUser)            // 获取用户信息
		users.PUT("/:id", userHandler.UpdateUser)         // 更新用户信息
		users.POST("/login", userHandler.Login)           // 用户登录
		users.GET("/current", userHandler.GetCurrentUser) // 获取当前用户信息
		users.GET("/profile", userHandler.GetUserProfile) // 获取个人资料
	}
}
