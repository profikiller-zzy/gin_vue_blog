package article_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/es_service"
	"github.com/gin-gonic/gin"
	"github.com/liu-cn/json-filter/filter"
)

type ArticleSearchRequest struct {
	model.PageInfo
	Tag    string `json:"tag" form:"tag"`
	IsUser bool   `json:"is_user" form:"is_user"` // 根据这个参数判断是否显示我收藏的文章列表
}

func (ArticleApi) ArticleListView(c *gin.Context) {
	var cr ArticleSearchRequest
	if err := c.ShouldBindQuery(&cr); err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}
	list, count, err := es_service.CommonList(cr.Key, cr.PageNum, cr.PageSize)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("查询失败", c)
		return
	}
	omitList := filter.Omit("list", list)
	response.OKWithList(omitList, int64(count), c)

	//boolSearch := elastic.NewBoolQuery()
}
