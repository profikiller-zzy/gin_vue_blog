package model

import "gin_vue_blog_AfterEnd/model/ctype"

// AuthModel 管理用户信息，由于gorm.Model结构体中有一个字段为
// deletedAt，该字段用于实现记录的逻辑删除(软删除)，本项目不用软
// 删除，所以自定义了MODEL取代了gorm.Model
type AuthModel struct {
	MODEL
	NickName      string           `gorm:"size:36" json:"nick_name"`                                                     // 昵称
	UserName      string           `gorm:"size:36" json:"user_name"`                                                     // 用户名
	Password      string           `gorm:"size:128" json:"password"`                                                     // 密码
	Avatar        string           `gorm:"size:256" json:"avatar"`                                                       // 头像地址
	Email         string           `gorm:"size:128" json:"email"`                                                        // 邮箱地址
	Tel           string           `gorm:"size:18" json:"tel"`                                                           // 电话号码
	Addr          string           `gorm:"size:64" json:"addr"`                                                          // 地址
	Token         string           `gorm:"size:128" json:"token"`                                                        // 其他平台的登录认证
	IP            string           `gorm:"size:20" json:"ip"`                                                            // ip地址
	Role          ctype.Role       `gorm:"size:4;default:1" json:"role"`                                                 // 角色：1 管理员 2 普通用户 3 游客
	SignStatus    ctype.SignStatus `gorm:"smallint(6)" json:"sign_status"`                                               // 注册方式
	ArticleModels []ArticleModel   `gorm:"foreignKey:AuthID" json:"-"`                                                   // 发布的文章列表
	CollectModels []ArticleModel   `gorm:"many2many:user_collect;foreignKey:AuthID;foreignReference:ArticleID" json:"-"` // 收藏的文章列表
}
