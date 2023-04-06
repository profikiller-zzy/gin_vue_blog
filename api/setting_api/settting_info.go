package setting_api

import (
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

// SettingInfoView 处理请求查看系统设置视图的函数
func (s SettingApi) SettingInfoView(c *gin.Context) {
	response.OKWithData(
		map[string]string{
			"id":   "000",
			"name": "Nick",
		}, c)
}
