package menu_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

// ImageSort 记录菜单项目下的图片和对应图片次序
type ImageSort struct {
	ImageID uint `json:"image_id"'`
	Sort    int  `json:"sort"`
}

// CreateMenuView 创建菜单项目
//
//		@Tags			菜单管理
//		@Summary		创建菜单项目
//		@description	创建菜单项目
//		@Router			/api/menu/ [POST]
//	 	@Success       	200	{object}	response.Response
func (MenuApi) CreateMenuView(c *gin.Context) {
	var mnReq MenuRequest
	err := c.ShouldBindJSON(&mnReq)
	// 判断参数是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &mnReq, c)
		return
	}
	// 重复值判断
	var menuModel model.MenuModel
	err = global.Db.Where("title = ? and path = ?", mnReq.Title, mnReq.Path).First(&menuModel).Error
	if err == nil { // 重复
		response.FailWithMessage("该菜单项已经存在", c)
		return
	}

	menuModel = model.MenuModel{
		Title:              mnReq.Title,
		Path:               mnReq.Path,
		Slogan:             mnReq.Slogan,
		Abstract:           mnReq.Abstract,
		AbstractSwitchTime: mnReq.AbstractSwitchTime,
		BannerSwitchTime:   mnReq.BannerSwitchTime,
		Sort:               mnReq.Sort,
	}
	// 创建menu入库
	err = global.Db.Create(&menuModel).Error
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("菜单记录入库失败", c)
		return
	}

	// 如果该菜单项目没有绑定banner图片
	if len(mnReq.ImageSortList) == 0 {
		response.OKWithMessage("菜单记录入库成功", c)
		return
	}
	// 给Banner和Menu的中间表入库
	var MenuBannerList []model.MenuBanner
	for _, imageSort := range mnReq.ImageSortList {
		// 这里先判断以下这张图片是否真的存在，如果存在才能与菜单相关联
		var image model.BannerModel
		err = global.Db.First(&image, imageSort.ImageID).Error
		if err == nil {
			MenuBannerList = append(MenuBannerList, model.MenuBanner{
				MenuID:   menuModel.ID,
				BannerID: imageSort.ImageID,
				Sort:     imageSort.Sort,
			})
		}
	}
	// 对中间表进行批量入库
	err = global.Db.Create(&MenuBannerList).Error
	if err != nil {
		response.FailWithMessage("MenuBanner中间表记录插入失败", c)
		global.Log.Error("MenuBanner中间表记录插入失败")
		return
	}
	response.OKWithMessage("菜单记录插入成功", c)
}
