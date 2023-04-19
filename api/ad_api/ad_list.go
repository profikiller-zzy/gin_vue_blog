package ad_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/common_service"
	"github.com/gin-gonic/gin"
	"strings"
)

func (AdApi) AdListView(c *gin.Context) {
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，error：%s", err.Error()), c)
		return
	}

	var adList []model.AdModel
	var count int64
	referer := c.GetHeader("Referer")
	isShow := true
	if strings.Contains(referer, "admin") { // 请求是从admin来的
		isShow = false
	}
	adList, count, err = common_service.PagingList(model.AdModel{IsShow: isShow}, common_service.PageInfoDebug{
		PageInfo: pageModel,
		Debug:    true,
	})

	// 判断 referer (来源)中是否包含 admin(管理员)，如果包含则将所有广告返回，如果不是，则只需要将 is_show=true 的广告返回即可
	response.OKWithPagingData(adList, count, c)
	return
}
