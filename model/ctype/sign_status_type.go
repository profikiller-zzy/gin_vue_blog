package ctype

import "encoding/json"

type SignStatus int

const (
	SignQQ    = 1 // QQ注册
	SignEmail = 2 // 邮箱注册
	SignGitee = 3 // Gitee注册
)

func (s SignStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s SignStatus) String() string {
	switch s {
	case SignQQ:
		return "QQ"
	case SignEmail:
		return "邮箱"
	case SignGitee:
		return "Gitee"
	default:
		return "其他途径"
	}
}
