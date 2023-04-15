package core

import (
	"gin_vue_blog_AfterEnd/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"time"
)

// InitGorm gorm连接到mysql数据库
func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		global.Log.Warnln("未配置mysql数据库，取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()
	// 设置mysql日志
	if global.Config.System.Env == "debug" {
		global.MysqlLog = logger.Default.LogMode(logger.Info) // 输出所有sql语句
	} else {
		global.MysqlLog = logger.Default.LogMode(logger.Error) // 只打印产生错误的sql
	}

	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		global.Log.Fatalf("[%s] mysql连接失败", dsn)
	}
	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)               // 最大空闲连接数
	sqlDb.SetMaxOpenConns(100)              // 连接池最大容量
	sqlDb.SetConnMaxLifetime(time.Hour * 4) // 连接最大复用时间，不能超过mysql的wait_timeout
	global.Log.Info("连接到mysql数据库成功！")
	return db
}
