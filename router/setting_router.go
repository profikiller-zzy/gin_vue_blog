package router

import (
	"gin_vue_blog_AfterEnd/api"
)

// SettingRouter 系统配置api
func (r RGroup) SettingRouter() {
	settingApiApp := api.ApiGroupApp.SettingApi
	r.GET("/setting/:name", settingApiApp.SettingInfoView)
	r.PUT("/setting/:name", settingApiApp.SettingInfoUpdate)
}
