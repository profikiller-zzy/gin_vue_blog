package flag

import "gin_vue_blog_AfterEnd/model"

func CreateESIndex() {
	model.ArticleModel{}.CreateIndex()
}
