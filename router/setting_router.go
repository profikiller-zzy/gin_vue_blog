package router

import (
	"gin_vue_blog_AfterEnd/api"
)

// SettingRouter 系统配置api
func (r RGroup) SettingRouter() {
	settingApi := api.ApiGroupApp.SettingApi
	r.GET("/setting/:name", settingApi.SettingInfoView)
	r.PUT("/setting/:name", settingApi.SettingInfoUpdate)
}
