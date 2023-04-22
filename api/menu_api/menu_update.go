package menu_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (MenuApi) MenuUpdateView(c *gin.Context) {
	var mnReq MenuRequest
	id := c.Param("id")
	err := c.ShouldBindJSON(&mnReq)
	// 判断参数是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &mnReq, c)
		return
	}

	var menuModel model.MenuModel
	err = global.Db.First(&menuModel, id).Error
	if err != nil {
		response.FailWithMessage("指定的菜单项不存在", c)
		return
	}

	// 先将该菜单项的关联图片清空
	err = global.Db.Model(&menuModel).Association("Banners").Clear()

	// 如果选择了banner，则需要对中间表进行更新
	if len(mnReq.ImageSortList) > 0 {
		// 操作中间表
		var menuBannerList = make([]model.MenuBanner, len(mnReq.ImageSortList))
		for index, imageSort := range mnReq.ImageSortList {
			menuBannerList[index] = model.MenuBanner{
				MenuID:   menuModel.ID,
				BannerID: imageSort.ImageID,
				Sort:     imageSort.Sort,
			}
		}
		err = global.Db.Create(&menuBannerList).Error
		if err != nil {
			global.Log.Error(err.Error())
			response.FailWithMessage("菜单与图片关联关系创建失败", c)
			return
		}
	}

	// 普通更新
	mnReqMap := structs.Map(&mnReq)
	err = global.Db.Model(&menuModel).Updates(mnReqMap).Error

	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("修改菜单失败", c)
		return
	}
	response.OKWithMessage("修改菜单成功", c)
}
