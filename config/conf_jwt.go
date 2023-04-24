package config

type Jwt struct {
	SecretKey  string `json:"secret_key" yaml:"secret_key"`   // 密钥
	ExpireTime int64  `json:"expire_time" yaml:"expire_time"` // 过期时间(单位为time.hour)
	Issuer     string `json:"issuer" yaml:"issuer"`           // 颁发者
}
