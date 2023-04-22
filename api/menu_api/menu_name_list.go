package menu_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

// MenuNameList 图片名称列表
//
//		@Tags			菜单管理
//		@Summary		获取图片名称列表
//		@description	获取图片名称列表
//		@Router			/api/menu_name/ [GET]
//	 	@Success       	200	{object}	response.Response{Data=[]MenuNameResponse}
func (MenuApi) MenuNameList(c *gin.Context) {
	var menuNameRepList []MenuNameResponse
	global.Db.Model(model.MenuModel{}).Select("id", "title", "path").Scan(&menuNameRepList)
	response.OKWithData(menuNameRepList, c)
}
