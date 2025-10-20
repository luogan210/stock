package service

import (
	"bytes"
	"errors"
	"fmt"
	"go-demo/config"
	"go-demo/model"
	"net/http"
	"time"

	jsoniter "github.com/json-iterator/go"
)

// WechatServiceAlternative 微信小程序服务（使用json-iterator）
type WechatServiceAlternative struct {
	config *config.WechatConfig
	json   jsoniter.API
}

// NewWechatServiceAlternative 创建微信服务实例
func NewWechatServiceAlternative() *WechatServiceAlternative {
	return &WechatServiceAlternative{
		config: config.GetWechatConfig(),
		json:   jsoniter.ConfigCompatibleWithStandardLibrary,
	}
}

// GenerateUrlLink 生成微信小程序urlLink
func (s *WechatServiceAlternative) GenerateUrlLink(req *model.WechatUrlLinkRequest) (*model.WechatUrlLinkResponse, error) {
	// 1. 业务逻辑验证
	if err := s.validateUrlLinkRequest(req); err != nil {
		return nil, err
	}

	// 2. 获取访问令牌
	accessToken, err := s.getAccessToken()
	if err != nil {
		return nil, fmt.Errorf("获取访问令牌失败: %w", err)
	}

	// 3. 计算过期时间
	expireTime := s.calculateExpireTime(req)

	// 4. 调用微信API生成urlLink
	urlLink, err := s.callWechatUrlLinkAPI(accessToken, req, expireTime)
	if err != nil {
		return nil, fmt.Errorf("生成urlLink失败: %w", err)
	}

	return &model.WechatUrlLinkResponse{
		UrlLink:    urlLink,
		ExpireTime: expireTime,
	}, nil
}

// callWechatUrlLinkAPI 调用微信API生成urlLink（使用json-iterator）
func (s *WechatServiceAlternative) callWechatUrlLinkAPI(accessToken string, req *model.WechatUrlLinkRequest, expireTime int64) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s", accessToken)

	// 构建请求参数
	requestBody := map[string]interface{}{
		"path": req.Path,
	}

	if req.Query != "" {
		requestBody["query"] = req.Query
	}

	if expireTime > 0 {
		requestBody["expire_type"] = 1
		requestBody["expire_time"] = expireTime
	} else {
		requestBody["expire_type"] = 0
	}

	fmt.Printf("requestBody %+v\n", requestBody)

	// 使用json-iterator编码，对特殊字符处理更友好
	jsonData, err := s.json.Marshal(requestBody)
	if err != nil {
		return "", err
	}

	fmt.Printf("jsonData %s\n", string(jsonData))

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		UrlLink string `json:"url_link"`
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	if err := s.json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		fmt.Printf("调用微信生成urlLink接口失败, errcode=%d, errmsg=%s\n", result.ErrCode, result.ErrMsg)
		return "", fmt.Errorf("生成urlLink失败: %s", result.ErrMsg)
	}

	return result.UrlLink, nil
}

// validateUrlLinkRequest 验证请求参数
func (s *WechatServiceAlternative) validateUrlLinkRequest(req *model.WechatUrlLinkRequest) error {
	if req.Path == "" {
		return errors.New("页面路径不能为空")
	}

	// 验证过期时间设置
	if req.IsExpire {
		switch req.ExpireType {
		case 0: // 指定时间失效
			if req.ExpireTime <= time.Now().Unix() {
				return errors.New("过期时间不能早于当前时间")
			}
		case 1: // 指定天数失效
			if req.ExpireInterval <= 0 {
				return errors.New("过期天数必须大于0")
			}
		default:
			return errors.New("无效的过期类型，只能是0（指定时间）或1（指定天数）")
		}
	}

	return nil
}

// getAccessToken 获取微信访问令牌
func (s *WechatServiceAlternative) getAccessToken() (string, error) {
	// 调用微信API获取access_token
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		s.config.AppID, s.config.AppSecret)

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
		ErrCode     int    `json:"errcode"`
		ErrMsg      string `json:"errmsg"`
	}

	if err := s.json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		return "", fmt.Errorf("获取access_token失败: %s", result.ErrMsg)
	}

	return result.AccessToken, nil
}

// calculateExpireTime 计算过期时间
func (s *WechatServiceAlternative) calculateExpireTime(req *model.WechatUrlLinkRequest) int64 {
	if !req.IsExpire {
		return 0 // 不设置过期时间
	}

	switch req.ExpireType {
	case 0: // 指定时间失效
		return req.ExpireTime
	case 1: // 指定天数失效
		return time.Now().Add(time.Duration(req.ExpireInterval) * 24 * time.Hour).Unix()
	default:
		return 0
	}
}
