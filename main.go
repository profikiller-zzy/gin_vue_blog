package main

import (
	"fmt"
	"gin_vue_blog_AfterEnd/core"
	_ "gin_vue_blog_AfterEnd/docs"
	"gin_vue_blog_AfterEnd/flag"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/router"
)

// @title           API 文档
// @version         1.0
// @description     gvb API文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      127.0.0.1:8080
// @BasePath  /

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// 读取配置文件，并将配置文件写入全局变量
	global.Config = core.InitConfig()
	// 初始化日志，并将日志写入全局变量
	global.Log = core.InitLogger()
	// 连接mysql数据库，并将数据库写入全局变量
	global.Db = core.InitGorm()
	// 连接redis数据库，并将数据库写入全局变量
	global.Redis = core.InitRedis()
	// 连接es数据库，并将数据库写入全局变量
	global.ES = core.InitES()

	// 初始化gin路由引擎
	r := router.InitRouter()
	global.Log.Info(fmt.Sprintf("gvb_sever 运行在:%s", global.Config.System.Addr()))

	// 捕获命令行参数，并对不同命令行参数的值来执行不同的操作
	flag.Parse()
	err := r.Run(global.Config.System.Addr())
	if err != nil {
		global.Log.Fatalf(err.Error())
	}
}
