package model

type BannerModel struct {
	MODEL
	Path string `json:"path"`                // 图片URL
	Hash string `json:"hash"`                // 图片的Hash值，用以判断重复图片
	Name string `gorm:"size:36" json:"name"` // 图片的名称
}
