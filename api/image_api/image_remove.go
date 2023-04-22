package image_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
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
	global.Db.Delete(&model.BannerModel{}, rmReq.IDList)
	response.FailWithMessage(fmt.Sprintf("删除 %d 张图片成功", count), c)
}
