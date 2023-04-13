package config

type QiNiu struct {
	Enable    bool    `yaml:"enable" json:"enable"` // 是否启用七牛存储，不启用则图片保存在本地
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` // 存储桶
	CDN       string  `json:"cdn" yaml:"cdn"`       // 访问图片地址的前缀
	Zone      string  `json:"zone" yaml:"zone"`     // 存储的地区
	Size      float64 `json:"size" yaml:"size"`     // 存储图片的大小限制(单位是5MB)
}
