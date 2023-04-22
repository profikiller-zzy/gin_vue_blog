package model

import (
	"gin_vue_blog_AfterEnd/model/ctype"
)

// MenuModel 菜单表
// Path可以是`/path/to/yourFile`，也可以是一个URL(以`http`或`https`开头的)，跳转到其他地方
type MenuModel struct {
	MODEL
	Title              string        `gorm:"size:32" json:"title"`                                                                   // 菜单名
	Path               string        `gorm:"size:32" json:"path"`                                                                    // 路径
	Slogan             string        `gorm:"size:64" json:"slogan"`                                                                  // 标语
	Abstract           ctype.Array   `json:"abstract"`                                                                               // 简介
	AbstractSwitchTime int           `json:"abstract_switch_time"`                                                                   // 简介的切换时间，单位为秒，为0表示不切换
	Banners            []BannerModel `gorm:"many2many:menu_banners;joinForeignKey:MenuID;joinReferences:BannerID" json:"menu_image"` // 菜单的图片列表
	BannerSwitchTime   int           `json:"banner_switch_time"`                                                                     // 菜单背景图片的切换周期，单位为秒，为0表示不切换
	Sort               int           `gorm:"size:10" json:"sort"`                                                                    // 菜单项目从左往右是第几位
}
