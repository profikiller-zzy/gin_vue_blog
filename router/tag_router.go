package router

import "gin_vue_blog_AfterEnd/api"

func (r RGroup) TagRouter() {
	tagApiApp := api.ApiGroupApp.TagApi
	r.GET("/tag/", tagApiApp.TagListView)
	r.PUT("/tag/", tagApiApp.TagCreateView)
	r.DELETE("/tag/", tagApiApp.TagRemoveView)
}
