package main

import (
	"fmt"
	"gin_vue_blog_AfterEnd/core"
	"gin_vue_blog_AfterEnd/global"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

func main() {
	// 读取配置文件，并将配置文件写入全局变量
	global.Config = core.InitConfig()
	// 初始化日志，并将日志写入全局变量
	global.Log = core.InitLogger()
	client = NewESClient()

	//// 测试创建索引
	//var demo DemoModel
	//demo.CreateIndex()

	//// 插入数据测试
	//var dataDemo = DemoModel{
	//	Title:     "Java基础",
	//	UserID:    2,
	//	CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
	//}
	//InsertData(&dataDemo)
	//fmt.Printf("文档ID：%s; 插入记录为:%s\n", dataDemo.ID, dataDemo.Title)

	// 分页查询测试
	dataDemoList, count := FindList("", 1, 10)
	fmt.Printf("一共查到 %d 条数据：\n", count)
	for _, value := range dataDemoList {
		fmt.Printf("%v\n", value)
	}

	//// 更新记录测试
	//Update("u3y56ocBOyzrYnP975a7", &DemoModel{
	//	Title: "我要学习golang后端",
	//})

	//// 批量删除测试
	//count, err := Remove([]string{
	//	"iE236ocBU9sEaXXeObTE",
	//	"uny46ocBOyzrYnP9KpbG",
	//})
	//if err != nil {
	//	logrus.Error(err)
	//}
	//logrus.Info(fmt.Sprintf("删除%d条记录成功!", count))
}
