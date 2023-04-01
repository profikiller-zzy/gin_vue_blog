package main

import (
	"gin_vue_blog_AfterEnd/core"
)

func main() {
	// 读取配置文件
	core.InitConfig()
	// 连接mysql数据库
	core.InitGorm()
}
