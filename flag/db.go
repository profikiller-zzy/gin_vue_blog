package flag

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
)

func MakeMigration() {
	var err error
	// 自定义多对多关系表
	err = global.Db.SetupJoinTable(&model.UserModel{}, "CollectModels", &model.UserCollect{})
	if err != nil {
		global.Log.Warn(err.Error())
	}
	err = global.Db.SetupJoinTable(&model.MenuModel{}, "MenuBanner", &model.MenuBanner{})
	if err != nil {
		global.Log.Warn(err.Error())
	}
	// 对模型自动迁移
	err = global.Db.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&model.UserModel{},
			&model.TagModel{},
			&model.ArticleModel{},
			&model.BannerModel{},
			&model.MessageModel{},
			&model.AdModel{},
			&model.CommentModel{},
			&model.MenuBanner{},
			model.FeedbackModel{},
			model.CategoryModel{},
			model.LoginDataModel{})
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	global.Log.Info("数据表迁移成功！")
}
