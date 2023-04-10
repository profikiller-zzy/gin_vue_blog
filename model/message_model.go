package model

type MessageModel struct {
	MODEL
	SenderID       uint      `gorm:"primaryKey" json:"sender_id"` // 消息发送人的ID
	SendUser       UserModel `gorm:"foreignKey:SenderID" json:"-"`
	SenderNickName string    `gorm:"size:36" json:"sender_nick_name"` // 消息发送人昵称
	SenderAvatar   string    `gorm:"size:256" json:"sender_avatar"`   // 消息发送人头像地址

	ReceiverID       uint      `gorm:"primaryKey" json:"receiver_id"` // 消息接收人的ID
	Receiver         UserModel `gorm:"foreignKey:SenderID" json:"-"`
	ReceiverNickName string    `gorm:"size:36" json:"receiver_nick_name"` // 消息接收人昵称
	ReceiverAvatar   string    `gorm:"size:256" json:"receiver_avatar"`   // 消息接收人头像地址
	IsRead           bool      `gorm:"default:false"json:"is_read"`       // 是否已读
	Content          string    `json:"content"`                           // 内容
}
