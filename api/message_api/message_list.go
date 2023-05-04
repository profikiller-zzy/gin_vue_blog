package message_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

type Message struct {
	SenderID         uint      `json:"send_user_id"` // 发送人id
	SenderNickName   string    `json:"send_user_nick_name"`
	SenderAvatar     string    `json:"send_user_avatar"`
	ReceiverID       uint      `json:"rev_user_id"` // 接收人id
	ReceiverNickName string    `json:"rev_user_nick_name"`
	ReceiverAvatar   string    `json:"rev_user_avatar"`
	Content          string    `json:"content"`       // 消息内容
	CreatedAt        time.Time `json:"created_at"`    // 最新的消息时间
	MessageCount     int       `json:"message_count"` // 消息条数
}

type MessageGroup map[uint]*Message

// MessageListView 用户与其他人的消息列表
// @Tags 消息管理
// @Summary 用户与其他人的消息列表
// @Description 用户与其他人的消息列表
// @Router /api/messages [get]
// @Param token header string  true  "token"
// @Produce json
// @Success 200 {object} res.Response{data=[]Message}
func (MessageApi) MessageListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var messageGroup = MessageGroup{}
	var messageList []model.MessageModel
	var messages = make([]Message, 0)

	global.Db.Order("created_at asc").
		Find(&messageList, "sender_id = ? or receiver_id = ?", claims.UserID, claims.UserID)
	for _, model := range messageList {
		// 判断是一个组的条件
		// send_user_id 和 rev_user_id 其中一个
		// 1 2  2 1
		// 1 3  3 1 是一组
		message := Message{
			SenderID:         model.SenderID,
			SenderNickName:   model.SenderNickName,
			SenderAvatar:     model.SenderAvatar,
			ReceiverID:       model.ReceiverID,
			ReceiverNickName: model.ReceiverNickName,
			ReceiverAvatar:   model.ReceiverAvatar,
			Content:          model.Content,
			CreatedAt:        model.CreatedAt,
			MessageCount:     1,
		}
		idNum := model.SenderID + model.ReceiverID
		val, ok := messageGroup[idNum]
		if !ok {
			// 不存在
			messageGroup[idNum] = &message
			continue
		}
		message.MessageCount = val.MessageCount + 1
		messageGroup[idNum] = &message
	}
	for _, message := range messageGroup {
		messages = append(messages, *message)
	}

	response.OKWithData(messages, c)
	return
}
