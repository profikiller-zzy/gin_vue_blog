package model

type ArticleModel struct {
	MODEL
	Title         string         `gorm:"size:32" json:"title"`                       // 文章标题
	Abstract      string         `json:"abstract"`                                   //  文章简介
	Content       string         `json:"content"`                                    // 文章内容
	PageView      int            `json:"page_view"`                                  // 浏览量
	CommentVolume int            `json:"comment_volume"`                             // 评论量
	CollectVolume int            `json:"collect_volume"`                             // 点赞量
	TagModels     []TagModel     `gorm:"many2many:article_tag" json:"tag_models"`    // 文章标签
	CommentModels []CommentModel `gorm:"foreignKey:ArticleID" json:"comment_models"` // 文章的评论列表
	AuthModel     AuthModel      `gorm:"foreignKey:ArticleID" json:"-"`              // 文章的作者
	Category      string         `gorm:"size:20" json:"category"`                    // 文章分类
	Source        string         `json:"source"'`                                    // 文章来源
	Link          string         `json:"link"`                                       // 原文链接
	Cover         ImageModel     `json:"-"`                                          // 文章封面图片
	CoverID       uint           `json:"cover_id"`                                   // 与封面的一对一关系
	NickName      string         `gorm:"size:36" json:"nick_name"`                   // 文章作者的昵称
	CoverPath     string         `json:"cover_path"`                                 // 文章封面

}
