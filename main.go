package main

import (
	"gin_vue_blog_AfterEnd/core"
	"gin_vue_blog_AfterEnd/global"
)

func main() {
	// 读取配置文件
	global.Config = core.InitConfig()
	// 连接mysql数据库
	global.Db = core.InitGorm()
}
