package user_api

import (
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/common_service"
	"github.com/gin-gonic/gin"
)

func (UserApi) UserPagingListView(c *gin.Context) {
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}

	var imageList = make([]model.UserModel, 0)
	var count int64
	// 对图片列表进行分页
	imageList, count, err = common_service.PagingList(model.UserModel{}, common_service.PageInfoDebug{
		PageInfo: pageModel,
		Debug:    true,
	})
	response.OKWithPagingData(imageList, count, c)
	return
}
