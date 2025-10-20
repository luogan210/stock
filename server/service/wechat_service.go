package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"go-demo/config"
	"go-demo/model"
	"net/http"
	"time"
)

// WechatService 微信小程序服务
type WechatService struct {
	config *config.WechatConfig
	// 这里可以注入缓存、数据库等依赖
	// cacheRepo repository.CacheRepository
	// wechatRepo repository.WechatRepository
}

// NewWechatService 创建微信服务实例
func NewWechatService() *WechatService {
	return &WechatService{
		config: config.GetWechatConfig(),
	}
}

// GenerateUrlLink 生成微信小程序urlLink
func (s *WechatService) GenerateUrlLink(req *model.WechatUrlLinkRequest) (*model.WechatUrlLinkResponse, error) {
	// 1. 业务逻辑验证
	if err := s.validateUrlLinkRequest(req); err != nil {
		return nil, err
	}

	// 2. 获取访问令牌
	accessToken, err := s.getAccessToken(s.config.AppID, s.config.AppSecret)
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

	// 5. 保存记录（可选）
	// err = s.saveUrlLinkRecord(req, urlLink, expireTime)

	return &model.WechatUrlLinkResponse{
		UrlLink:    urlLink,
		ExpireTime: expireTime,
	}, nil
}

// GetUrlLinkInfo 获取urlLink信息
func (s *WechatService) GetUrlLinkInfo(id uint) (*model.WechatUrlLinkInfo, error) {
	// 这里应该从数据库查询
	// 暂时返回模拟数据
	if id == 0 {
		return nil, errors.New("记录ID不能为空")
	}

	return &model.WechatUrlLinkInfo{
		ID:         id,
		Path:       "pages/index/index",
		Query:      "id=123",
		UrlLink:    "https://wxaurl.cn/xxx",
		ExpireTime: time.Now().Add(24 * time.Hour).Unix(),
		CreatedAt:  time.Now().Unix(),
		Status:     1,
	}, nil
}

// validateUrlLinkRequest 验证请求参数
func (s *WechatService) validateUrlLinkRequest(req *model.WechatUrlLinkRequest) error {
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

// getAccessToken 获取微信访问令牌（appid和secret以参数的形式传入）
func (s *WechatService) getAccessToken(appid, secret string) (string, error) {
	// 这里应该先从缓存中获取，如果过期则重新获取
	// token, err := s.cacheRepo.GetAccessToken(appid)
	// if err == nil && token != "" {
	//     return token, nil
	// }

	// 调用微信API获取access_token
	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s",
		appid, secret)

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

	// 保存到缓存
	// s.cacheRepo.SetAccessToken(appid, result.AccessToken, result.ExpiresIn)

	return result.AccessToken, nil
}

// calculateExpireTime 计算过期时间
func (s *WechatService) calculateExpireTime(req *model.WechatUrlLinkRequest) int64 {
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

// callWechatUrlLinkAPI 调用微信API生成urlLink
func (s *WechatService) callWechatUrlLinkAPI(accessToken string, req *model.WechatUrlLinkRequest, expireTime int64) (string, error) {
	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/generate_urllink?access_token=%s", accessToken)

	// 构建请求参数
	requestBody := map[string]interface{}{
		"path": req.Path,
	}

	if req.Query != "" {
		requestBody["query"] = req.Query
	}

	// 根据过期类型设置参数
	if req.IsExpire {
		requestBody["expire_type"] = req.ExpireType
		if req.ExpireType == 0 { // 指定时间失效
			requestBody["expire_time"] = req.ExpireTime
		} else if req.ExpireType == 1 { // 指定天数失效
			requestBody["expire_interval"] = req.ExpireInterval
		}
	}

	fmt.Printf("requestBody %+v\n", requestBody)

	// 使用自定义的JSON编码器，避免转义特殊字符
	jsonData, err := s.MarshalJSONWithoutEscape(requestBody)
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

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if result.ErrCode != 0 {
		fmt.Printf("调用微信生成urlLink接口失败, errcode=%d, errmsg=%s\n", result.ErrCode, result.ErrMsg)
		return "", fmt.Errorf("生成urlLink失败: %s", result.ErrMsg)
	}

	return result.UrlLink, nil
}

// MarshalJSONWithoutEscape 自定义JSON编码，避免转义特殊字符
func (s *WechatService) MarshalJSONWithoutEscape(v interface{}) ([]byte, error) {
	// 创建buffer
	var buf bytes.Buffer

	// 创建编码器并设置选项
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false) // 关键：不转义HTML字符
	encoder.SetIndent("", "")    // 不缩进

	// 编码数据
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}

	// 移除末尾的换行符
	data := buf.Bytes()
	if len(data) > 0 && data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	return data, nil
}

// saveUrlLinkRecord 保存urlLink记录
func (s *WechatService) saveUrlLinkRecord(req *model.WechatUrlLinkRequest, urlLink string, expireTime int64) error {
	// 这里应该保存到数据库
	// record := &model.WechatUrlLinkInfo{
	//     Path:       req.Path,
	//     Query:      req.Query,
	//     UrlLink:    urlLink,
	//     ExpireTime: expireTime,
	//     CreatedAt:  time.Now().Unix(),
	//     Status:     1,
	// }
	// return s.wechatRepo.SaveUrlLink(record)
	return nil
}
