package core

import (
	"gin_vue_blog_AfterEnd/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

// InitGorm gorm连接到mysql数据库
func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		log.Println("未配置mysql数据库，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	// 设置mysql日志
	var mysqlLogger logger.Interface
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info) //
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		log.Fatalf("[%s] mysql连接失败", dsn)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDb.SetMaxOpenConns(100)              // 连接池最大容量
	sqlDb.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	return db
}
