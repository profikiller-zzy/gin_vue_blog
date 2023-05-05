package article_service

import (
	"context"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"github.com/sirupsen/logrus"
)

// InsertArticleToES 像索引中插入数据
func InsertArticleToES(article *model.ArticleModel) (err error) {
	indexResponse, err := global.ESClient.Index().
		Index(article.Index()).
		BodyJson(article).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	article.ID = indexResponse.Id
	return nil
}
