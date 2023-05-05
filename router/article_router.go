package router

import (
	"gin_vue_blog_AfterEnd/api"
	"gin_vue_blog_AfterEnd/middleware"
)

func (r RGroup) ArticleRouter() {
	articleApiApp := api.ApiGroupApp.ArticleApi
	r.PUT("/create_article/", middleware.JwtAuth(), articleApiApp.CreateArticle)
}
