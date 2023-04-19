package router

import "gin_vue_blog_AfterEnd/api"

func (r RGroup) AdRouter() {
	imageApiApp := api.ApiGroupApp.AdAPI
	r.POST("/advertise/", imageApiApp.AdCreateView)
	r.GET("/advertise/", imageApiApp.AdListView)
	r.PUT("/advertise/:id", imageApiApp.AdUpdateView)
	r.DELETE("/advertise/", imageApiApp.AdRemoveView)
}
