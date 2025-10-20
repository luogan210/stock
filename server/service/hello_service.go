package service

// HelloService 处理Hello相关的业务逻辑
type HelloService struct{}

// NewHelloService 创建HelloService实例
func NewHelloService() *HelloService {
	return &HelloService{}
}

// GetHelloMessage 获取Hello消息
func (s *HelloService) GetHelloMessage() map[string]interface{} {
	return map[string]interface{}{
		"message": "Hello, World!",
		"service": "hello",
	}
}
