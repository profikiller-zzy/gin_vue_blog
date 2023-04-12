package setting_api

import (
	"gin_vue_blog_AfterEnd/config"
	"gin_vue_blog_AfterEnd/core"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

// SettingInfoUpdate 处理请求查看系统设置视图的函数
func (SettingApi) SettingInfoUpdate(c *gin.Context) {
	var confSiteInfo config.SiteInfo
	err := c.ShouldBindJSON(&confSiteInfo)
	if err != nil {
		//
		response.FailWithCode(response.ParameterError, c)
		return
	}
	global.Config.SiteInfo = confSiteInfo
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("修改成功", c)
}
