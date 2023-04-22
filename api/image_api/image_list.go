package image_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   uint   `json:"id"`
	Path string `json:"path"` // 图片URL，如果存储在本地则为图片路径，存储在云服务器上则是图片链接
	Name string `json:"name"` // 图片的名称
}

// ImageList 返回信息简略的图片列表
//
//		@Tags			图片管理
//		@Summary		获取信息简略的图片列表
//		@description	获取信息简略的图片列表
//		@Router			/api/imageList/ [GET]
//	 	@Success       	200	{object}	response.Response{Data=[]ImageResponse}
func (ImageApi) ImageList(c *gin.Context) {
	var imageList = make([]ImageResponse, 0)

	global.Db.Model(&model.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	response.OKWithData(imageList, c)
}
