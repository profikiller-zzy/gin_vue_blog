package model

import (
	"gin_vue_blog_AfterEnd/model/ctype"
)

// MenuModel 菜单表
type MenuModel struct {
	MODEL
	MenuTitle        string        `gorm:"size:32" json:"menu_title"`                                               // 菜单名
	MenuTitleEn      string        `gorm:"size:32" json:"menu_title_en"`                                            // 英文菜单名
	Slogan           string        `gorm:"size:64" json:"slogan"`                                                   // 标语
	Abstract         ctype.Array   `json:"abstract"`                                                                // 简介
	AbstractTime     string        `json:"abstract_time"`                                                           // 简介的切换时间
	MenuBanner       []BannerModel `gorm:"many2many:menu_banner;foreignKey:ID;joinForeignKey:ID" json:"menu_image"` // 菜单的图片列表
	BannerSwitchTime int           `json:"banner_switch_time"`                                                      // 菜单背景图片的切换周期，为0表示不切换
	Sort             int           `gorm:"size:10" json:"sort"`                                                     // 菜单项目从左往右是第几位
}
