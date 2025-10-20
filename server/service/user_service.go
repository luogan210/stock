package service

import (
	"errors"
	"go-demo/model"
)

// UserService 用户服务
type UserService struct {
	// 这里可以注入数据库连接等依赖
}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{}
}

// CreateUser 创建用户
func (s *UserService) CreateUser(req *model.CreateUserRequest) (*model.UserInfo, error) {
	// 这里应该添加数据库操作
	// 暂时返回模拟数据
	user := &model.User{
		ID:       1,
		Username: req.Username,
		Email:    req.Email,
		Nickname: req.Nickname,
		Status:   1,
	}

	return &model.UserInfo{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Nickname: user.Nickname,
		Status:   user.Status,
	}, nil
}

// GetUserByID 根据ID获取用户信息
func (s *UserService) GetUserByID(id uint) (*model.UserInfo, error) {
	// 这里应该从数据库查询
	// 暂时返回模拟数据
	if id == 0 {
		return nil, errors.New("用户不存在")
	}

	return &model.UserInfo{
		ID:       id,
		Username: "demo_user",
		Email:    "demo@example.com",
		Nickname: "演示用户",
		Status:   1,
	}, nil
}

// UpdateUser 更新用户信息
func (s *UserService) UpdateUser(id uint, req *model.UpdateUserRequest) (*model.UserInfo, error) {
	// 这里应该更新数据库
	// 暂时返回模拟数据
	userInfo, err := s.GetUserByID(id)
	if err != nil {
		return nil, err
	}

	if req.Nickname != "" {
		userInfo.Nickname = req.Nickname
	}
	if req.Avatar != "" {
		userInfo.Avatar = req.Avatar
	}

	return userInfo, nil
}

// Login 用户登录
func (s *UserService) Login(req *model.LoginRequest) (*model.LoginResponse, error) {
	// 这里应该验证用户名密码
	// 暂时返回模拟数据
	if req.Username == "admin" && req.Password == "123456" {
		return &model.LoginResponse{
			Token: "mock_token_123456",
			UserInfo: &model.UserInfo{
				ID:       1,
				Username: "admin",
				Email:    "admin@example.com",
				Nickname: "管理员",
				Status:   1,
			},
		}, nil
	}

	return nil, errors.New("用户名或密码错误")
}

// GetCurrentUserInfo 获取当前用户信息（通过token）
func (s *UserService) GetCurrentUserInfo(token string) (*model.UserInfo, error) {
	// 这里应该从数据库或缓存中验证token
	// 暂时使用简单的验证逻辑
	if token == "mock_token_123456" {
		return &model.UserInfo{
			ID:       1,
			Username: "admin",
			Email:    "admin@example.com",
			Nickname: "管理员",
			Status:   1,
		}, nil
	}

	return nil, errors.New("无效的token")
}
