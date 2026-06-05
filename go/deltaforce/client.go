// 三角洲数据帝 - 开放平台 API SDK (Go)
// Base URL: https://orzice.com/workApi
package deltaforce

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

// BaseURL 默认基础URL，可按需修改
var BaseURL = "https://orzice.com/workApi"

// Token 全局API密钥，按需修改
var Token = ""

// Client API客户端
type Client struct {
	BaseURL string
	Token   string
	client  *http.Client
}

// NewClient 创建客户端，自动使用全局 BaseURL 和 Token
func NewClient(timeout time.Duration) *Client {
	if timeout == 0 {
		timeout = 30 * time.Second
	}
	return &Client{
		BaseURL: BaseURL,
		Token:   Token,
		client:  &http.Client{Timeout: timeout},
	}
}

// Get 发送GET请求，自动附加 token
func (c *Client) Get(path string, params map[string]string) (map[string]interface{}, error) {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, fmt.Errorf("解析URL失败: %w", err)
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	q.Set("token", c.Token)
	u.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("User-Agent", "DeltaForceApi/1.0")
	req.Header.Set("Accept", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("无权限访问，请检查 token 是否正确")
	}
	if resp.StatusCode != 200 {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("请求失败，状态码: %d, 响应: %s", resp.StatusCode, string(body))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	return result, nil
}

// Post 发送POST请求
func (c *Client) Post(path string, params map[string]string, data map[string]string) (map[string]interface{}, error) {
	u, err := url.Parse(c.BaseURL + path)
	if err != nil {
		return nil, fmt.Errorf("解析URL失败: %w", err)
	}

	q := u.Query()
	for k, v := range params {
		q.Set(k, v)
	}
	u.RawQuery = q.Encode()

	form := url.Values{}
	for k, v := range data {
		form.Set(k, v)
	}

	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}
	req.Header.Set("User-Agent", "DeltaForceApi/1.0")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == 404 {
		return nil, fmt.Errorf("无权限访问，请检查 token 是否正确")
	}
	if resp.StatusCode != 200 {
		respBody, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("请求失败，状态码: %d, 响应: %s", resp.StatusCode, string(respBody))
	}

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("解析响应失败: %w", err)
	}
	return result, nil
}
