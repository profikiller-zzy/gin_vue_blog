package common_service

import "regexp"

var ValidURLRegular = regexp.MustCompile(`(http|ftp|https):\/\/[\w\-_]+(\.[\w\-_]+)+([\w\-\.,@?^=%&:/~\+#]*[\w\-\@?^=%&/~\+#])?`)

// CheckURLValid 使用正则式判断跳转链接是否合法
func CheckURLValid(url string) bool {
	result := ValidURLRegular.FindAllStringSubmatch(url, -1)
	if result == nil { // 没有匹配合法URL正则式
		return false
	}
	return true
}
