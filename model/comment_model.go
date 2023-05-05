package model

type CommentModel struct {
	MODEL

	ParentCommentID uint            `gorm:"column:parent_comment_id;default:0;not null" json:"parent_comment_id"` // 父评论的ID
	ParentComment   *CommentModel   `gorm:"foreignKey:ParentCommentID" json:"parent_comment"`                     // 父评论
	SubComments     []*CommentModel `gorm:"foreignKey:ParentCommentID" json:"sub_comments"`                       // 子评论列表，父评论和子评论之间是一对多的关系

	Content         string    `gorm:"size:256" json:"content"`                                    // 评论内容
	LikeCount       int       `gorm:"size:8;default:0" json:"like_count"`                         // 点赞量
	SubCommentCount int       `gorm:"size:8;default:0" json:"sub_comment_count"`                  // 子评论数
	ArticleID       string    `gorm:"column:article_id;not null;index;size:32" json:"article_id"` // 关联文章的ID
	User            UserModel `gorm:"foreignKey:UserID" json:"-"`                                 // 关联的用户
	UserID          uint      `json:"user_id"`                                                    // 关联用户的ID
}
