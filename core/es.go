package core

import (
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
	global.Log.Info("连接到ES成功！")
	return client
}
