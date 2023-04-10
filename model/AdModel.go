package model

type AdModel struct {
	MODEL
	Title     string `gorm:"size:32" json:"title"` // 广告的标题
	Href      string `json:"href"`                 // 广告的跳转连接
	ImagePath string `json:"image_path"`           // 图片的URL
	IsShow    bool   `json:"is_show"`              // 是否展示
}
