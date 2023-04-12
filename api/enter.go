package api

import "gin_vue_blog_AfterEnd/api/setting_api"

// ApiGroup 是对整个Api定义的结构体的统合，方便链式调用
type ApiGroup struct {
	SettingApi setting_api.SettingApi
}

// ApiGroupApp 实例化ApiGroup对象
var ApiGroupApp = new(ApiGroup)
