package menu_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID   uint   `json:"id"`
	Path string `json:"path"`
}

type MenuResponse struct {
	model.MenuModel
	Banners []Banner `json:"banners"`
}

// MenuListView 菜单列表，用于返回给后台管理系统的完整菜单列表
//
//		@Tags			菜单管理
//		@Summary		查询菜单列表
//		@description	查询菜单列表
//		@Router			/api/menu/ [GET]
//	 	@Success       	200	{object}	response.Response{Data=[]MenuResponse}
func (MenuApi) MenuListView(c *gin.Context) {
	// 第一步 查询菜单
	var menuList []model.MenuModel
	var menuIDList []uint
	global.Db.Order("sort desc").Find(&menuList).Select("id").Scan(&menuIDList)
	// 查询menu和banner的中间表
	var menuBannerList = make([]model.MenuBanner, 0)
	global.Db.Preload("BannerModel").Order("sort desc").Find(&menuBannerList, menuIDList)
	var menuListRep = make([]MenuResponse, 0)
	for _, menu := range menuList {
		// 这里需要对查询出来的每个菜单项目，使用该菜单项目的ID去联合中间表查出它关联的图片，并且将其写入到MenuListRep中
		// banners即menu对应的图片列表，并且以sort递减排序
		var banners []Banner
		for _, menuBanner := range menuBannerList {
			if menuBanner.MenuID != menu.ID {
				continue
			}
			banners = append(banners, Banner{
				ID:   menuBanner.BannerID,
				Path: menuBanner.BannerModel.Path,
			})
		}
		menuListRep = append(menuListRep, MenuResponse{
			MenuModel: menu,
			Banners:   banners,
		})
	}
	response.OKWithData(menuListRep, c)
}
