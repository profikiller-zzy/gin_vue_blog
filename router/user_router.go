package router

import (
	"gin_vue_blog_AfterEnd/api"
	"gin_vue_blog_AfterEnd/middleware"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

var CaptchaStore = cookie.NewStore([]byte("EKfGVamg0nM6WlHK"))

// UserRouter 用户管理路由
func (r RGroup) UserRouter() {
	userApiApp := api.ApiGroupApp.UserApi
	r.Use(sessions.Sessions("captcha", CaptchaStore))
	r.POST("/email_login/", userApiApp.EmailLoginView)
	r.POST("/login/", userApiApp.QQLoginView)
	r.POST("/user_create/", middleware.JwtAdmin(), userApiApp.UserCreateView)
	r.POST("/logout/", middleware.JwtAuth(), userApiApp.UserLogoutView)
	r.GET("/user/", middleware.JwtAuth(), userApiApp.UserListView)
	r.PUT("/user_role/", middleware.JwtAuth(), userApiApp.UserUpdateRoleView)
	r.PUT("/user_password/", middleware.JwtAuth(), userApiApp.UserUpdatePasswordView)
	r.DELETE("/user_remove/", middleware.JwtAuth(), userApiApp.UserRemoveView)
	r.POST("/email_binding/", middleware.JwtAuth(), userApiApp.EmailBindingView)
}
