package config

type QQ struct {
	AppID    string `json:"app_id" yaml:"app_id"` //
	Key      string `json:"key" yaml:"key"`
	Redirect string `json:"redirect" yaml:"redirect"` // 登录之后的回调地址
}

// GetPath 暂定
func (q QQ) GetPath() string {
	return q.Redirect
}
