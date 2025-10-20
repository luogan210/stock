package config

// WechatConfig 微信小程序配置
type WechatConfig struct {
	AppID     string `json:"app_id"`     // 小程序AppID
	AppSecret string `json:"app_secret"` // 小程序AppSecret
}

// GetWechatConfig 获取微信配置
func GetWechatConfig() *WechatConfig {
	return &WechatConfig{
		AppID:     getEnv("WECHAT_APP_ID", ""),
		AppSecret: getEnv("WECHAT_APP_SECRET", ""),
	}
}
