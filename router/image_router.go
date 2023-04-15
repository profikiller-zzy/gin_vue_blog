package router

import "gin_vue_blog_AfterEnd/api"

func (r RGroup) ImageRouter() {
	imageApiApp := api.ApiGroupApp.ImageApi
	r.POST("/image/", imageApiApp.ImageUploadingView)
	r.GET("/image/", imageApiApp.ImageListView)
}
