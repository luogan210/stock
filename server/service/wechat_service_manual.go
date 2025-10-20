package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-demo/config"
	"go-demo/model"
	"net/http"
	"strconv"
	"time"
)

// WechatServiceManual 微信小程序服务（手动构建JSON）
type WechatServiceManual struct {
	config *config.WechatConfig
}

// NewWechatServiceManual 创建微信服务实例
func NewWechatServiceManual() *WechatServiceManual {
	return &WechatServiceManual{
		config: config.GetWechatConfig(),
	}
}

// GenerateUrlLink 生成微信小程序urlLink
func (s *WechatServiceManual) GenerateUrlLink(req *model.WechatUrlLinkRequest) (*model.WechatUrlLinkResponse, error) {
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

// callWechatUrlLinkAPI 调用微信API生成urlLink（手动构建JSON）
func (s *WechatServiceManual) callWechatUrlLinkAPI(accessToken string, req *model.WechatUrlLinkRequest, expireTime int64) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s", accessToken)

	// 手动构建JSON字符串，避免转义问题
	jsonStr := s.buildJSONString(req, expireTime)

	fmt.Printf("jsonData %s\n", jsonStr)

	resp, err := http.Post(url, "application/json", bytes.NewBufferString(jsonStr))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		UrlLink string `json:"url_link"`
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}

	// 使用标准库解码响应
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		fmt.Printf("调用微信生成urlLink接口失败, errcode=%d, errmsg=%s\n", result.ErrCode, result.ErrMsg)
		return "", fmt.Errorf("生成urlLink失败: %s", result.ErrMsg)
	}

	return result.UrlLink, nil
}

// buildJSONString 手动构建JSON字符串
func (s *WechatServiceManual) buildJSONString(req *model.WechatUrlLinkRequest, expireTime int64) string {
	var jsonStr bytes.Buffer

	jsonStr.WriteString("{")
	jsonStr.WriteString(`"path":"`)
	jsonStr.WriteString(req.Path)
	jsonStr.WriteString(`"`)

	if req.Query != "" {
		jsonStr.WriteString(`,"query":"`)
		jsonStr.WriteString(req.Query) // 直接写入，不转义
		jsonStr.WriteString(`"`)
	}

	if expireTime > 0 {
		jsonStr.WriteString(`,"expire_type":1`)
		jsonStr.WriteString(`,"expire_time":`)
		jsonStr.WriteString(strconv.FormatInt(expireTime, 10))
	} else {
		jsonStr.WriteString(`,"expire_type":0`)
	}

	jsonStr.WriteString("}")

	return jsonStr.String()
}

// validateUrlLinkRequest 验证请求参数
func (s *WechatServiceManual) validateUrlLinkRequest(req *model.WechatUrlLinkRequest) error {
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
func (s *WechatServiceManual) getAccessToken() (string, error) {
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

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		return "", fmt.Errorf("获取access_token失败: %s", result.ErrMsg)
	}

	return result.AccessToken, nil
}

// calculateExpireTime 计算过期时间
func (s *WechatServiceManual) calculateExpireTime(req *model.WechatUrlLinkRequest) int64 {
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
