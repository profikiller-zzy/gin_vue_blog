package main

import (
	"fmt"
	"gin_vue_blog_AfterEnd/core"
	"gin_vue_blog_AfterEnd/flag"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/router"
)

func main() {
	// 读取配置文件，并将配置文件写入全局变量
	global.Config = core.InitConfig()
	// 初始化日志，并将日志写入全局变量
	global.Log = core.InitLogger()
	// 连接mysql数据库，并将数据库写入全局变量
	global.Db = core.InitGorm()
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
