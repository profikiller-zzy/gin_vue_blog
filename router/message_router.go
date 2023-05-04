package router

import (
	"gin_vue_blog_AfterEnd/api"
	"gin_vue_blog_AfterEnd/middleware"
)

func (r RGroup) MessageRouter() {
	messageApiApp := api.ApiGroupApp.MessageApi
	r.POST("message", messageApiApp.MessageCreate)
	r.GET("message_all", middleware.JwtAdmin(), messageApiApp.MessageListAllView)
	r.GET("message", middleware.JwtAuth(), messageApiApp.MessageListView)
	r.GET("message_record", middleware.JwtAuth(), messageApiApp.MessageRecordView)
}
