package flag

import "flag"

type Options struct {
	DB bool
}

// Parse 解析命令参数，并对不同的命令行参数的值来执行不同的操作
func Parse() {
	dbFlag := flag.Bool("db", false, "auto migrate database")
	flag.Parse()
	if *dbFlag {
		MakeMigration()
	}
}
