package config

import "fmt"

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"` //
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` // 登录之后的回调地址
}

// GetPath 暂定
func (q QQ) GetPath() string {
	if q.AppID == "" || q.Redirect == "" || q.Redirect == "" {
		return ""
	}
	// qq登录第三方网站固定跳转地址
	return fmt.Sprintf("https://graph.qq.com/oauth2.0/show?which=Login&display=pc&response_type=code&client_id=%s&redirect_uri=%s", q.AppID, q.Redirect)
}
