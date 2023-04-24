package api

// gin-swagger middleware
// swagger embed files
import (
	"gin_vue_blog_AfterEnd/api/ad_api"
	"gin_vue_blog_AfterEnd/api/image_api"
	"gin_vue_blog_AfterEnd/api/menu_api"
	"gin_vue_blog_AfterEnd/api/setting_api"
	"gin_vue_blog_AfterEnd/api/user_api"
)

// ApiGroup 是对整个Api定义的结构体的统合，方便链式调用
type ApiGroup struct {
	SettingApi setting_api.SettingApi
	ImageApi   image_api.ImageApi
	AdAPI      ad_api.AdApi
	MenuApi    menu_api.MenuApi
	UserApi    user_api.UserApi
}

// ApiGroupApp 实例化ApiGroup对象
var ApiGroupApp = new(ApiGroup)
