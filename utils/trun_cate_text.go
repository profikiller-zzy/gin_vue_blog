package utils

import "unicode/utf8"

func TruncateText(text string, length int) string {
	// 如果文章长度小于等于截取长度，则直接返回原文
	if utf8.RuneCountInString(text) <= length {
		return text
	}

	// 获取文章的前length个字符
	runes := []rune(text)
	resRunes := runes[:length]

	// 添加省略号
	return string(resRunes) + "..."
}
