package main

import (
	"context"
	"encoding/json"
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"log"
	"os"
)

// NewESClient 初始化一个新的ES连接
func NewESClient() *elastic.Client {
	client, err := elastic.NewClient(
		//elastic 服务地址
		elastic.SetURL(global.Config.ES.URL()),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
		// 设置错误日志输出
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		// 设置info日志输出
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))
	if err != nil {
		logrus.Fatalf("es连接失败 %s", err.Error())
	}
	logrus.Info("连接到ES成功！")
	return client
}

// InsertData 向索引中插入数据
func InsertData(data *DemoModel) (err error) {
	indexResponse, err := client.Index().
		Index(data.Index()).
		BodyJson(data).Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	data.ID = indexResponse.Id
	return nil
}

// FindList 列表查询
func FindList(key string, pageNum, pageSize int) (demoList []DemoModel, count int) {
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

	res, err := client.
		Search(DemoModel{}.Index()).
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
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	return demoList, count
}

// FindSourceList 指定返回的字段，对查询结果进行分页
func FindSourceList(key string, page, limit int) {
	boolSearch := elastic.NewBoolQuery()
	from := page
	if key != "" {
		boolSearch.Must(
			elastic.NewMatchQuery("title", key),
		)
	}
	if limit == 0 {
		limit = 10
	}
	if from == 0 {
		from = 1
	}

	res, err := client.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		Source(`{"_source": ["title"]}`).
		From((from - 1) * limit).
		Size(limit).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return
	}
	count := int(res.Hits.TotalHits.Value) //搜索到结果总条数
	demoList := []DemoModel{}
	for _, hit := range res.Hits.Hits {
		var demo DemoModel
		data, err := hit.Source.MarshalJSON()
		if err != nil {
			logrus.Error(err.Error())
			continue
		}
		err = json.Unmarshal(data, &demo)
		if err != nil {
			logrus.Error(err)
			continue
		}
		demo.ID = hit.Id
		demoList = append(demoList, demo)
	}
	fmt.Println(demoList, count)
}

// Update 更新指定ID文档的title
func Update(id string, data *DemoModel) error {
	_, err := client.
		Update().
		Index(DemoModel{}.Index()).
		Id(id).
		Doc(map[string]string{
			"title": data.Title,
		}).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return err
	}
	logrus.Info("更新demo成功")
	return nil
}

// Remove 批量删除es文档
func Remove(idList []string) (count int, err error) {
	// 创建一个批量操作的服务 bulkService
	bulkService := client.Bulk().Index(DemoModel{}.Index()).Refresh("true")
	for _, id := range idList { // 为每个 ID 创建一个 elastic.NewBulkDeleteRequest().Id(id) 请求
		req := elastic.NewBulkDeleteRequest().Id(id)
		bulkService.Add(req)
	}
	res, err := bulkService.Do(context.Background())
	return len(res.Succeeded()), err
}

func FindDemoByKey(key string) {
	boolSearch := elastic.NewBoolQuery()
	if key != "" {
		boolSearch.Must(
			elastic.NewTermQuery("key", key),
		)
	}
	res, err := client.
		Search(DemoModel{}.Index()).
		Query(boolSearch).
		Size(2).
		Do(context.Background())
	fmt.Println(err)
	fmt.Println(res.Hits.TotalHits.Value)
	fmt.Println(len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		fmt.Println(string(hit.Source))
	}
}
