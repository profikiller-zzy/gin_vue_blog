package config

import (
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiu struct {
	Enable    bool    `yaml:"enable" json:"enable"` // 是否启用七牛云存储，不启用则图片保存在本地
	AccessKey string  `json:"access_key" yaml:"access_key"`
	SecretKey string  `json:"secret_key" yaml:"secret_key"`
	Bucket    string  `json:"bucket" yaml:"bucket"` // 存储桶
	CDN       string  `json:"cdn" yaml:"cdn"`       // Content Delivery Network，通过绑定自定义静态域名来作为静态资源文件的前缀
	Zone      string  `json:"zone" yaml:"zone"`     // 存储的地区
	Size      float64 `json:"size" yaml:"size"`     // 存储图片的大小限制(单位是5MB)
}

// GetUpToken 获取上传七牛云的简单上传凭证
func (q QiNiu) GetUpToken() string {
	if q.AccessKey == "" || q.SecretKey == "" {
		return ""
	}
	bucket := q.Bucket
	putPolicy := storage.PutPolicy{
		Scope:      bucket,
		ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	}
	// 设置2小时为有效期
	putPolicy.Expires = 7200
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

// GetCfg 获取七牛云服务器文件上传的配置
func (q QiNiu) GetCfg() storage.Config {
	Cfg := storage.Config{}
	// 空间对应的机房：华东-浙江2
	region, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	Cfg.Region = &region
	// 是否使用https域名
	Cfg.UseHTTPS = false
	// 上传是否使用CDN加速
	Cfg.UseCdnDomains = false

	return Cfg
}
