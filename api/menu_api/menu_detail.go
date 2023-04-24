package menu_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuDetailView(c *gin.Context) {
	id := c.Param("id")
	var menuModel model.MenuModel
	err := global.Db.First(&menuModel, id).Error
	if err != nil {
		response.FailWithMessage("菜单不存在!", c)
		return
	}

	// 查询menu和banner的中间表
	var menuBannerList = make([]model.MenuBanner, 0)
	global.Db.Preload("BannerModel").Order("sort desc").Find(&menuBannerList, id)
	var menuRep = MenuResponse{}
	var banners = make([]Banner, len(menuBannerList))
	for index, menuBanner := range menuBannerList {
		// 对查询出来的菜单项目，使用该菜单项目的ID去联合中间表查出它关联的图片，并且将其写入到menuRep中
		// banners即menu对应的图片列表，并且以sort递减排序
		banners[index] = Banner{
			ID:   menuBanner.BannerID,
			Path: menuBanner.BannerModel.Path,
		}
	}
	menuRep = MenuResponse{
		MenuModel: menuModel,
		Banners:   banners,
	}

	response.OKWithData(menuRep, c)
}
