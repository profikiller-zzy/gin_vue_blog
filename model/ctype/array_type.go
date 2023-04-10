package ctype

import (
	"database/sql/driver"
	"strings"
)

type Array []string

// Scan 用于将一个 []byte 类型的值转换为一个 []string 类型的值，分隔符为"\n"。
// 用于将数据库中值转化为Go语言中的值
func (a *Array) Scan(value interface{}) error {
	v, _ := value.([]byte) // 先将value强制转换为`[]byte`类型
	if string(v) == "" {
		*a = []string{}
		return nil
	}
	*a = strings.Split(string(v), "\n")
	return nil
}

// Value 将 Go 语言中的值转换为一个空接口类型的值，并返回。这个空接口类型的值可以被存储到数据库中
func (a Array) Value() (driver.Value, error) {
	// Join 将用户传入的分隔符去连接string数组中的值，返回连接结果string
	return strings.Join(a, "\n"), nil
}
