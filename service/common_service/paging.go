package common_service

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gorm.io/gorm"
)

type PageInfoDebug struct {
	model.PageInfo
	Debug bool // 是否打印sql语句
}

// PagingList 对不同数据模型的数据项进行分页，返回指定页的所有数据和所有数据项的数量
func PagingList[T any](model T, debug PageInfoDebug) (list []T, count int64, err error) {
	// 对数据模型列表进行分页
	db := global.Db
	if debug.Debug {
		db = global.Db.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	var offset int
	// 使用到model T入参中携带的条件参数
	count = db.Where(model).Select("id").Find(&list).RowsAffected
	if debug.PageNum == 0 { // 如果
		offset = 0
	} else {
		offset = (debug.PageNum - 1) * debug.PageSize
	}

	if debug.Sort == "" { // 默认按照创建时间从新到旧排
		debug.Sort = "created_at desc"
	}
	err = db.Where(model).Limit(debug.PageSize).Offset(offset).Order(debug.Sort).Find(&list).Error
	return list, count, err
}
