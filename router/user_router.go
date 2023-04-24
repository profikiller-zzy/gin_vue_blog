package router

import "gin_vue_blog_AfterEnd/api"

// UserRouter 用户管理路由
func (r RGroup) UserRouter() {
	userApiApp := api.ApiGroupApp.UserApi
	r.POST("/email_login/", userApiApp.EmailLoginView)
	r.GET("/user/", userApiApp.UserPagingListView)
}
