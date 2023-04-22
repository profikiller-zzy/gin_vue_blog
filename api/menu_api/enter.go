package menu_api

import "gin_vue_blog_AfterEnd/model/ctype"

type MenuApi struct {
}

type MenuRequest struct {
	Title              string      `json:"title" binding:"required" msg:"请输入菜单名称" structs:"title"`  // 菜单名
	Path               string      `json:"path" binding:"required" msg:"请输入菜单路径" structs:"path"`    // 菜单跳转路径
	Slogan             string      `json:"slogan" structs:"slogan"`                                 // 标语
	Abstract           ctype.Array `json:"abstract" structs:"abstract"`                             // 简介
	AbstractSwitchTime int         `json:"abstract_switch_time" structs:"abstract_switch_time"`     // 简介的切换周期，单位为秒，为0表示不切换
	BannerSwitchTime   int         `json:"banner_switch_time" structs:"banner_switch_time"`         // 菜单背景图片的切换周期，单位为秒，为0表示不切换
	Sort               int         `json:"sort" binding:"required" msg:"请输入菜单展示的序号" structs:"sort"` // 菜单项目从左往右是第几位
	ImageSortList      []ImageSort `json:"image_sort_list" structs:"-"`
}
