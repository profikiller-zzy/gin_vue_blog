package config

import "fmt"

type Redis struct {
	IP       string `json:"ip"`
	Port     int    `json:"port"` // 端口
	Password string `json:"password"`
	PoolSize int    `json:"poolSize"` // 连接池大小
}

func (r Redis) ReturnAddr() string {
	return fmt.Sprintf("%s:%d", r.IP, r.Port)
}
