package router

import (
	"gin_vue_blog_AfterEnd/api"
)

// SettingRouter 系统配置api
func (r RouterGroup) SettingRouter() {
	settingApi := api.ApiGroupApp.SettingApi
	r.GET("/setting/", settingApi.SettingInfoView)
}
