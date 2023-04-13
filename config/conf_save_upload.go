package config

type SaveUpload struct {
	Size int64  `json:"size" yaml:"size"` // 上传图片的大小限制，单位为MB
	Path string `json:"path" yaml:"path"` // 图片保存到本地的路径
}
