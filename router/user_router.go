package router

import "gin_vue_blog_AfterEnd/api"

// UserRouter 用户管理路由
func (r RGroup) UserRouter() {
	userApiApp := api.ApiGroupApp.UserApi
	r.POST("/user/", userApiApp.EmailLoginView)
}
