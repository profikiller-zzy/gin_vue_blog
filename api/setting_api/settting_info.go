package setting_api

import (
	"gin_vue_blog_AfterEnd/config"
	"gin_vue_blog_AfterEnd/core"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type SettingUri struct {
	Name string `uri:"name"` // 添加`uri`标签，可以指定该字段对应的请求参数的名称，以便在路由函数中进行解析。
}

// SettingInfoView 处理请求查看相应模块视图的函数
func (SettingApi) SettingInfoView(c *gin.Context) {
	var uri SettingUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}

	switch uri.Name {
	case "site":
		response.OKWithData(global.Config.SiteInfo, c)
	case "qq":
		response.OKWithData(global.Config.QQ, c)
	case "email":
		response.OKWithData(global.Config.Email, c)
	case "qi_niu":
		response.OKWithData(global.Config.QiNiu, c)
	case "jwt":
		response.OKWithData(global.Config.Jwt, c)
	default:
		response.FailWithMessage("请输入正确的uri，“site”、“qq”、“email”、“qi_niu“或”jwt“", c)
	}
}

// SettingInfoUpdate 处理修改相应模块设置参数的函数
// (注意事项) 通过指定uri以获取和修改不同模块的做法，可以减少接口数量
// 也有弊端，弊端就是接口不能统一
func (SettingApi) SettingInfoUpdate(c *gin.Context) {
	var uri SettingUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	switch uri.Name {
	case "site":
		SiteUpdate(c)
	case "qq":
		QQUpdate(c)
	case "email":
		EmailUpdate(c)
	case "qi_niu":
		QiNiuUpdate(c)
	case "jwt":
		JwtUpdate(c)
	default:
		response.FailWithMessage("请输入正确的uri，“site”、“qq”、“email”、“qi_niu“或”jwt“", c)
	}
	core.SetYaml()
}

// SiteUpdate 包含以下共5个函数分别对应修改本地全局变量`global.Config`中的不同模块
func SiteUpdate(c *gin.Context) {
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

func QQUpdate(c *gin.Context) {
	var confQQInfo config.QQ
	err := c.ShouldBindJSON(&confQQInfo)
	if err != nil {
		//
		response.FailWithCode(response.ParameterError, c)
		return
	}
	global.Config.QQ = confQQInfo
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("修改成功", c)
}

func EmailUpdate(c *gin.Context) {
	var confEmailInfo config.Email
	err := c.ShouldBindJSON(&confEmailInfo)
	if err != nil {
		//
		response.FailWithCode(response.ParameterError, c)
		return
	}
	global.Config.Email = confEmailInfo
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("修改成功", c)
}

func QiNiuUpdate(c *gin.Context) {
	var confQINiuInfo config.QiNiu
	err := c.ShouldBindJSON(&confQINiuInfo)
	if err != nil {
		//
		response.FailWithCode(response.ParameterError, c)
		return
	}
	global.Config.QiNiu = confQINiuInfo
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("修改成功", c)
}

func JwtUpdate(c *gin.Context) {
	var confJwtInfo config.Jwt
	err := c.ShouldBindJSON(&confJwtInfo)
	if err != nil {
		//
		response.FailWithCode(response.ParameterError, c)
		return
	}
	global.Config.Jwt = confJwtInfo
	err = core.SetYaml()
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("修改成功", c)
}
