package config

type Email struct {
	Host             string `yaml:"host" json:"host"` // SMTP 服务器地址
	Port             int    `yaml:"port" json:"port"`
	User             string `yaml:"user" json:""`
	Password         string `yaml:"password" json:"user"`
	DefaultFormEmail string `yaml:"default_form_email" json:"default_form_email"` // 默认发件人名称
	UseSSL           bool   `yaml:"use_ssl" json:"use_ssl"`                       // 是否使用SSL协议进行加密
	UserTsl          bool   `yaml:"user_tsl" json:"user_tsl"`                     // 是否使用Tsl协议进行加密 (User Transport Security Layer)
}
