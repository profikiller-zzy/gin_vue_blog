package flag

import (
	"flag"
	"gin_vue_blog_AfterEnd/global"
)

type Options struct {
	DB   bool
	User string // -u admin -u user
	ES   string // -es create -es delete
}

// Parse 解析命令参数，并对不同的命令行参数的值来执行不同的操作
func Parse() {
	dbFlag := flag.Bool("db", false, "auto migrate database")
	userFlag := flag.String("u", "", "create user or admin")
	esFlag := flag.String("es", "", "create or delete index")
	flag.Parse()
	var option = Options{
		DB:   *dbFlag,
		User: *userFlag,
		ES:   *esFlag,
	}
	Execute(option)
}

func Execute(options Options) {
	if options.DB {
		MakeMigration()
		return
	}
	// 输入非法
	if options.User != "" && options.User != "user" && options.User != "admin" {
		global.Log.Error("Invalid user type. Please use \"user\" or \"admin\".")
		return
	}
	switch options.User {
	case "user":
		CreateUser("user")
	case "admin":
		CreateUser("admin")
	}

	if options.ES != "" && options.ES != "create" && options.ES != "delete" {
		global.Log.Error("Invalid es type. Please use \"create\" or \"delete\".")
		return
	}
	switch options.ES {
	case "create": // 创建索引
		CreateESIndex()
	case "delete":

	}
}
