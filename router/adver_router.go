package router

import "gin_vue_blog_AfterEnd/api"

func (r RGroup) AdRouter() {
	imageApiApp := api.ApiGroupApp.AdAPI
	r.POST("/advertise/", imageApiApp.AdCreateView)
}
