package image_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" msg:"请选择图片ID"`
	Name string `json:"name" binding:"required" msg:"请输入图片名称"`
}

// ImageUpdateView 更新图片
//
//		@Tags			图片管理
//		@Summary		更新特定图片信息
//		@description	更新特定图片信息
//		@param			iuReq body ImageUpdateRequest true "更新后的图片信息"
//		@Router			/api/image/ [PUT]
//	 	@Success       	200	{object}	response.Response
//		@Failure		500	{object}	response.Response
func (ImageApi) ImageUpdateView(c *gin.Context) {
	var iuReq ImageUpdateRequest
	err := c.ShouldBindJSON(&iuReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &iuReq, c)
		return
	}
	var imageModel model.BannerModel
	err = global.Db.Take(&imageModel, iuReq.ID).Error
	if err != nil { // 没有找到
		response.FailWithMessage("文件不存在", c)
		return
	}
	// 找到了就将传入的name替换掉旧的name
	err = global.Db.Model(&imageModel).Update("name", iuReq.Name).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("修改广告成功", c)
		return
	}
	response.OKWithMessage("图片名称修改成功", c)
}
