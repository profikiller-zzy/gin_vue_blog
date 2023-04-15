package utils

import (
	"crypto/md5"
	"fmt"
)

// MD5 将输入源文件进行MD5 hash，返回以16进制表示的128位hash值
func MD5(src []byte) string {
	return fmt.Sprintf("%x", md5.Sum(src))
}
