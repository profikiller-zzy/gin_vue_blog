package router

import (
	"gin_vue_blog_AfterEnd/api"
	"gin_vue_blog_AfterEnd/middleware"
)

// UserRouter 用户管理路由
func (r RGroup) UserRouter() {
	userApiApp := api.ApiGroupApp.UserApi
	r.POST("/email_login/", userApiApp.EmailLoginView)
	r.GET("/user/", middleware.JwtAuth(), userApiApp.UserListView)
	r.PUT("user_role", middleware.JwtAuth(), userApiApp.UserUpdateRoleView)
	r.PUT("user_password", middleware.JwtAuth(), userApiApp.UserUpdatePasswordView)
}
