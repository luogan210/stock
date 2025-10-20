package model

// WechatUrlLinkRequest 生成微信小程序urlLink请求
type WechatUrlLinkRequest struct {
	Path           string `json:"path" binding:"required"` // 小程序页面路径
	Query          string `json:"query"`                   // 查询参数
	IsExpire       bool   `json:"is_expire"`               // 是否设置过期时间
	ExpireType     int    `json:"expire_type"`             // 过期类型：0-指定时间失效，1-指定天数失效
	ExpireTime     int64  `json:"expire_time"`             // 过期时间戳（当expire_type为0时使用）
	ExpireInterval int    `json:"expire_interval"`         // 过期天数（当expire_type为1时使用）
}

// WechatUrlLinkResponse 生成微信小程序urlLink响应
type WechatUrlLinkResponse struct {
	UrlLink    string `json:"url_link"`    // 生成的urlLink
	ExpireTime int64  `json:"expire_time"` // 过期时间戳
}

// WechatConfig 微信小程序配置
type WechatConfig struct {
	AppID     string `json:"app_id"`     // 小程序AppID
	AppSecret string `json:"app_secret"` // 小程序AppSecret
}

// WechatAccessToken 微信访问令牌
type WechatAccessToken struct {
	AccessToken string `json:"access_token"` // 访问令牌
	ExpiresIn   int    `json:"expires_in"`   // 过期时间（秒）
	ExpireTime  int64  `json:"expire_time"`  // 过期时间戳
}

// WechatUrlLinkInfo urlLink信息
type WechatUrlLinkInfo struct {
	ID         uint   `json:"id"`          // 记录ID
	Path       string `json:"path"`        // 页面路径
	Query      string `json:"query"`       // 查询参数
	UrlLink    string `json:"url_link"`    // 生成的urlLink
	ExpireTime int64  `json:"expire_time"` // 过期时间
	CreatedAt  int64  `json:"created_at"`  // 创建时间
	Status     int    `json:"status"`      // 状态：1-有效，0-已过期
}
