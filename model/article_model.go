package model

import "gin_vue_blog_AfterEnd/model/ctype"

// ArticleModel 管理文章数据。最后三个数据项是冗余数据，是为了避免反复查表
// 而使用空间换时间的做法。
type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`                       // 文章标题
	Abstract      string         `json:"abstract"`                                   // 文章简介
	Content       string         `json:"content"`                                    // 文章内容
	PageView      int            `json:"page_view"`                                  // 浏览量
	CommentCount  int            `json:"comment_count"`                              // 评论量
	LikeCount     int            `json:"like_count"`                                 // 点赞量
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"tag_models"`    // 文章标签 文章-标签为gorm普通多对多
	Tags          ctype.Array    `gorm:"type:string;size:64" json:"tags"`            // 文章标签，这里的标签是用"\n"隔开的字符串
	CommentModels []CommentModel `gorm:"foreignKey:CommentID" json:"comment_models"` // 文章的评论列表
	User          UserModel      `gorm:"foreignKey:UserID" json:"-"`                 // 文章的作者
	UserID        uint           `json:"user_id"`                                    // 文章作者的ID
	NickName      string         `gorm:"size:36" json:"nick_name"`                   // 文章作者的昵称
	Category      CategoryModel  `gorm:"foreignKey:CategoryID" json:"category"`      // 文章分类，文章-分类为多对一
	Source        string         `json:"source"`                                     // 文章来源
	Link          string         `json:"link"`                                       // 原文链接
	Banner        BannerModel    `gorm:"foreignKey:BannerID" json:"-"`               // 文章封面
	BannerID      uint           `json:"cover_id"`                                   // 文章封面ID
	BannerPath    string         `json:"banner_path"`                                // 文章封面
}
