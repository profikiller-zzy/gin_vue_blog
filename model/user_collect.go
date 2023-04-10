package model

import "time"

// UserCollect 此表记录用户和文章之间的收藏关系，为自定义多对多，新增了`创建时间`的属性
type UserCollect struct {
	UserID    uint         `gorm:"primaryKey"`           // 收藏用户的ID
	User      UserModel    `gorm:"foreignKey:UserID"`    // 收藏用户
	ArticleID uint         `gorm:"primaryKey"`           // 收藏文章的ID
	Article   ArticleModel `gorm:"foreignKey:ArticleID"` // 收藏的文章
	CreatedAt time.Time    `json:"created_at"`           // 创建时间
}
