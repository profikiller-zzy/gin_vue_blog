package image_api

import (
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/image_service"
	"github.com/gin-gonic/gin"
	"mime/multipart"
)

// ImageUploadingView 上传图片并将图片保存在uploads文件夹中
//
//		@Tags			图片管理
//		@Summary		上传多个图片
//		@description	上传多个图片
//		@param			image formData file true "需要上传的图片文件"
//		@Router			/api/image/ [POST]
//	 	@Success       	200	{object}	response.Response{Date=[]model.FileUploadResponse}
//		@Failure		500	{object}	response.Response
func (ImageApi) ImageUploadingView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		response.LogFail(err, c)
		return
	}
	var FileHeaderList []*multipart.FileHeader = form.File["image"]
	if len(FileHeaderList) == 0 {
		response.FailWithMessage("没有指定任何文件或者文件不存在", c)
		return
	}

	var upResList []model.FileUploadResponse = make([]model.FileUploadResponse, len(FileHeaderList))
	for index, FileHeader := range FileHeaderList {
		upResList[index] = image_service.ImageService{}.ImageUploadService(FileHeader, c)
	}
	response.OKWithData(upResList, c)
}
