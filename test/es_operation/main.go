package main

import (
	"gin_vue_blog_AfterEnd/core"
	"gin_vue_blog_AfterEnd/global"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func InitES() *elastic.Client {
	var err error
	options := []elastic.ClientOptionFunc{
		elastic.SetURL("http://localhost:9200/"),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
	}
	client, err := elastic.NewClient(options...)
	if err != nil {
		global.Log.Error("es连接失败，错误信息: %s", err.Error())
		return nil
	}
	global.Log.Info("es连接成功!")
	return client
}

func main() {
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	client = InitES()
	var demo DemoModel
	demo.CreateIndex()
}
