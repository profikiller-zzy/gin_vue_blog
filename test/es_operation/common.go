package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
)

type DemoModel struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	UserID   string `json:"user_id"`
	CreateAt string `json:"create_at"`
}

func (DemoModel) Index() string {
	return "demo_index"
}

func (DemoModel) Mapping() string {
	return `
{
  "settings": {
    "index":{
      "max_result_window": "100000"
    }
  }, 
  "mappings": {
    "properties": {
      "title": { 
        "type": "text"
      },
      "key": { 
        "type": "keyword"
      },
      "user_id": {
        "type": "integer"
      },
      "created_at":{
        "type": "date",
        "null_value": "null",
        "format": "[yyyy-MM-dd HH:mm:ss]"
      }
    }
  }
}
`
}

// CreateIndex 创建索引
func (demo DemoModel) CreateIndex() error {
	if demo.IndexExists() {
		// 有索引
		demo.DeleteIndex()
	}

	// 没有索引
	createIndex, err := client.
		CreateIndex(demo.Index()).
		BodyString(demo.Mapping()).
		Do(context.Background())
	if err != nil {
		logrus.Error(fmt.Sprintf("索引创建失败，报错信息：%s", err.Error()))
		return err
	}
	// 索引创建成功但未得到确认或是索引创建被es节点拒绝
	if !createIndex.Acknowledged {
		logrus.Error("索引创建失败")
		return err
	}
	logrus.Info("索引创建成功!")
	return nil
}

// DeleteIndex 删除索引
func (demo DemoModel) DeleteIndex() error {
	logrus.Info("索引存在，删除索引")
	// 删除索引
	indexDelete, err := client.DeleteIndex(demo.Index()).Do(context.Background())
	if err != nil {
		logrus.Error("删除索引失败")
		logrus.Error(err.Error())
		return err
	}
	if !indexDelete.Acknowledged {
		logrus.Error("删除索引失败")
		return err
	}
	logrus.Info("索引删除成功")
	return nil
}

// IndexExists 索引是否存在
func (demo DemoModel) IndexExists() bool {
	isExist, err := client.
		IndexExists(demo.Index()).
		Do(context.Background())
	if err != nil {
		logrus.Error(err.Error())
		return isExist
	}
	return isExist
}
