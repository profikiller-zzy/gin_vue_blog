package es_service

import (
	"context"
	"encoding/json"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

func CommonList(key string, pageNum, pageSize int) (list []model.ArticleModel, count int, err error) {
	// 创建一个新的布尔查询 boolSearch，用于构建 Elasticsearch 查询条件
	boolSearch := elastic.NewBoolQuery()
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if pageNum <= 0 {
		pageNum = 1
	}

	res, err := global.ESClient.
		Search(model.ArticleModel{}.Index()). // 指定索引
		Query(boolSearch).
		From((pageNum - 1) * pageSize).
		Size(pageSize). // 设置每页的数量
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count = int(res.Hits.TotalHits.Value) //搜索到结果总条数
	for _, hit := range res.Hits.Hits {
		var article model.ArticleModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &article)
		if err != nil {
			logrus.Error(err)
			continue
		}
		article.ID = hit.Id
		list = append(list, article)
	}
	return list, count, err
}
