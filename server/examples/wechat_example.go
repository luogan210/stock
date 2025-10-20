package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// WechatUrlLinkRequest 生成微信小程序urlLink请求
type WechatUrlLinkRequest struct {
	Path           string `json:"path"`            // 小程序页面路径
	Query          string `json:"query"`           // 查询参数
	IsExpire       bool   `json:"is_expire"`       // 是否设置过期时间
	ExpireType     int    `json:"expire_type"`     // 过期类型：0-指定时间失效，1-指定天数失效
	ExpireTime     int64  `json:"expire_time"`     // 过期时间戳（当expire_type为0时使用）
	ExpireInterval int    `json:"expire_interval"` // 过期天数（当expire_type为1时使用）
}

// WechatUrlLinkResponse 生成微信小程序urlLink响应
type WechatUrlLinkResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    struct {
		UrlLink    string `json:"url_link"`
		ExpireTime int64  `json:"expire_time"`
	} `json:"data"`
}

// 测试JSON转义问题
func testJSONEscape() {
	fmt.Println("=== JSON转义测试 ===\n")

	testCases := []struct {
		name  string
		query string
	}{
		{
			name:  "简单查询参数",
			query: "id=123",
		},
		{
			name:  "包含&字符",
			query: "id=123&type=product",
		},
		{
			name:  "复杂查询参数",
			query: "id=123&type=product&category=electronics&brand=apple&model=iphone",
		},
		{
			name:  "包含特殊字符",
			query: "id=123&type=product&name=iPhone 14 Pro&price=999.99",
		},
	}

	for _, tc := range testCases {
		fmt.Printf("测试用例: %s\n", tc.name)
		fmt.Printf("查询参数: %s\n", tc.query)

		// 测试标准JSON编码
		standardJSON, _ := json.Marshal(map[string]interface{}{
			"path":  "pages/product/detail",
			"query": tc.query,
		})
		fmt.Printf("标准JSON: %s\n", string(standardJSON))

		// 测试自定义编码（模拟）
		customJSON := marshalJSONWithoutEscape(map[string]interface{}{
			"path":  "pages/product/detail",
			"query": tc.query,
		})
		fmt.Printf("自定义JSON: %s\n", string(customJSON))

		fmt.Println("---")
	}
}

// marshalJSONWithoutEscape 模拟自定义JSON编码
func marshalJSONWithoutEscape(v interface{}) []byte {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false) // 关键：不转义HTML字符
	encoder.SetIndent("", "")    // 不缩进

	if err := encoder.Encode(v); err != nil {
		return nil
	}

	// 移除末尾的换行符
	data := buf.Bytes()
	if len(data) > 0 && data[len(data)-1] == '\n' {
		data = data[:len(data)-1]
	}

	return data
}

// 示例1：生成不设置过期时间的urlLink
func generatePermanentUrlLink() {
	req := WechatUrlLinkRequest{
		Path:     "pages/index/index",
		Query:    "id=123&type=product",
		IsExpire: false, // 不设置过期时间
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Fatal("JSON序列化失败:", err)
	}

	resp, err := http.Post("http://localhost:8080/api/wechat/url-link",
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("请求失败:", err)
	}
	defer resp.Body.Close()

	var response WechatUrlLinkResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal("响应解析失败:", err)
	}

	fmt.Printf("永久urlLink: %s\n", response.Data.UrlLink)
}

// 示例2：生成指定时间过期的urlLink
func generateExpiredUrlLink() {
	// 设置7天后过期
	expireTime := time.Now().Add(7 * 24 * time.Hour).Unix()

	req := WechatUrlLinkRequest{
		Path:       "pages/product/detail",
		Query:      "id=456",
		IsExpire:   true,
		ExpireType: 0, // 指定时间失效
		ExpireTime: expireTime,
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Fatal("JSON序列化失败:", err)
	}

	resp, err := http.Post("http://localhost:8080/api/wechat/url-link",
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("请求失败:", err)
	}
	defer resp.Body.Close()

	var response WechatUrlLinkResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal("响应解析失败:", err)
	}

	fmt.Printf("过期urlLink: %s\n", response.Data.UrlLink)
	fmt.Printf("过期时间: %s\n", time.Unix(response.Data.ExpireTime, 0).Format("2006-01-02 15:04:05"))
}

// 示例3：生成指定天数后过期的urlLink
func generateIntervalUrlLink() {
	req := WechatUrlLinkRequest{
		Path:           "pages/activity/share",
		Query:          "activity_id=789",
		IsExpire:       true,
		ExpireType:     1,  // 指定天数失效
		ExpireInterval: 30, // 30天后过期
	}

	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Fatal("JSON序列化失败:", err)
	}

	resp, err := http.Post("http://localhost:8080/api/wechat/url-link",
		"application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal("请求失败:", err)
	}
	defer resp.Body.Close()

	var response WechatUrlLinkResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatal("响应解析失败:", err)
	}

	fmt.Printf("30天过期urlLink: %s\n", response.Data.UrlLink)
	fmt.Printf("过期时间: %s\n", time.Unix(response.Data.ExpireTime, 0).Format("2006-01-02 15:04:05"))
}

func main() {
	fmt.Println("=== 微信小程序 urlLink 生成示例 ===\n")

	// 首先测试JSON转义问题
	testJSONEscape()

	fmt.Println("\n=== 实际API调用示例 ===\n")

	fmt.Println("1. 生成不设置过期时间的urlLink:")
	generatePermanentUrlLink()
	fmt.Println()

	fmt.Println("2. 生成指定时间过期的urlLink:")
	generateExpiredUrlLink()
	fmt.Println()

	fmt.Println("3. 生成指定天数后过期的urlLink:")
	generateIntervalUrlLink()
	fmt.Println()

	fmt.Println("=== 示例完成 ===")
}
