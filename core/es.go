package core

import (
	"gin_vue_blog_AfterEnd/global"
	"github.com/olivere/elastic/v7"
)

func InitES() *elastic.Client {
	var err error
	options := []elastic.ClientOptionFunc{
		elastic.SetURL(global.Config.ES.URL()),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(global.Config.ES.User, global.Config.ES.Password),
	}
	client, err := elastic.NewClient(options...)
	if err != nil {
		global.Log.Error("es连接失败 %s", err.Error())
		return nil
	}
	global.Log.Info("es连接成功!")
	return client
}
