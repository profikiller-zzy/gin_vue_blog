package config

import (
	"fmt"
	"strconv"
)

type Mysql struct {
	Host     string `yaml:"host"`      // 服务器地址
	Port     int    `yaml:"port"`      // 端口
	Db       string `yaml:"db"`        // 数据库名
	User     string `yaml:"user"`      // 数据库用户名
	Password string `yaml:"password"`  // 密码
	LogLevel string `yaml:"log_level"` // 日志等级是指日志消息的重要性和优先性
	Config   string `yaml:"config"`    // 高级设置，列如charset
}

// Dsn 返回mysql连接名
func (m Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?%s",
		m.User,
		m.Password,
		m.Host,
		strconv.Itoa(m.Port),
		m.Db,
		m.Config)
}
