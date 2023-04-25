package ctype

import "encoding/json"

type ImageStorageMode int

const (
	Local ImageStorageMode = 1 // 将图片存储在本地
	QiNiu ImageStorageMode = 2 // 将图片存储在七牛云服务器
)

func (i ImageStorageMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i ImageStorageMode) String() string {
	switch i {
	case Local:
		return "本地"
	case QiNiu:
		return "七牛云"

	default:
		return "OtherUser"
	}
}
