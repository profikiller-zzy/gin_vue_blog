package utils

// IsInStringList 判断输入的字符串是否在输入的字符串切片中
func IsInStringList(str string, strList []string) bool {
	for _, suffix := range strList {
		if str == suffix {
			return true
		}
	}
	return false
}
