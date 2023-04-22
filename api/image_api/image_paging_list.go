package image_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/common_service"
	"github.com/gin-gonic/gin"
)

// ImagePagingListView 图片列表
//
//		@Tags			图片管理
//		@Summary		获取图片列表
//		@description	获取图片列表
//		@param			adReq query model.PageInfo false "查询参数"
//		@Router			/api/image/ [GET]
//	 	@Success       	200	{object}	response.Response{Data=response.ListResponse[model.BannerModel]}
func (ImageApi) ImagePagingListView(c *gin.Context) {
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，error：%s", err.Error()), c)
		return
	}

	var imageList = make([]model.BannerModel, 0)
	var count int64
	// 对图片列表进行分页
	imageList, count, err = common_service.PagingList(model.BannerModel{}, common_service.PageInfoDebug{
		PageInfo: pageModel,
		Debug:    true,
	})
	response.OKWithPagingData(imageList, count, c)
	return
}
