package image_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ImageRemoveView 删除图片
//
//		@Tags			广告管理
//		@Summary		删除广告
//		@description	删除广告
//		@param			rmReq body model.RemoveRequest true "需要删除的广告ID列表"
//		@Router			/api/advertise/ [DELETE]
//	 	@Success       	200	{object}	response.Response
//		@Failure		500	{object}	response.Response
func (ImageApi) ImageRemoveView(c *gin.Context) {
	var rmReq model.RemoveRequest
	var imageList []model.BannerModel
	var count int64 = 0

	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，error：%s", err.Error()), c)
		return
	}

	count = global.Db.Find(&imageList, rmReq.IDList).RowsAffected
	if count == 0 { // 需要删除的图片ID没有在数据库中查到
		response.FailWithMessage("文件不存在", c)
		return
	}

	var IDList = make([]uint, count)
	for index, image := range imageList {
		IDList[index] = image.ID
	}
	err = global.Db.Transaction(func(tx *gorm.DB) error {
		// 先将该菜单项的关联图片清空
		err = global.Db.Where("banner_id in ?", IDList).Delete(&model.MenuBanner{}).Error
		if err != nil {
			return err
		}
		// 再删除对应的菜单项列表
		err = global.Db.Delete(&model.BannerModel{}, IDList).Error
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("删除图片失败, 错误信息:%s", err.Error()), c)
		return
	}
	response.FailWithMessage(fmt.Sprintf("删除 %d 张图片成功", count), c)
}
