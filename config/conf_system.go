package config

import (
	"fmt"
)

type System struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Env  string `yaml:"env"` // 开发模式debug，打印所有信息
}

// Addr 返回服务器地址和端口组成的字符串
func (s System) Addr() string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}
