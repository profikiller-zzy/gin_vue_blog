package model

// FeedbackModel 管理用户对网站发出的反馈、建议
type FeedbackModel struct {
	MODEL
	Email        string `gorm:"size:64" json:"email"`          // 反馈用户的邮箱
	Content      string `gorm:"size:128" json:"content"`       // 内容
	ReplyContent string `gorm:"size:128" json:"reply_content"` // 回复的内容
	IsReply      bool   `json:"is_reply"`                      // 是否回复
}
