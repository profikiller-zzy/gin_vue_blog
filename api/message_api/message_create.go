package message_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type MessageCreateRequest struct {
	SenderID   uint   `json:"sender_id" binding:"required" msg:"请指定发送人"`   // 消息发送人的ID
	ReceiverID uint   `json:"receiver_id" binding:"required" msg:"请指定接收人"` // 消息接收人的ID
	Content    string `json:"content" binding:"required" msg:"输入需要发送的内容"`  // 内容
}

func (MessageApi) MessageCreate(c *gin.Context) {
	var mcReq MessageCreateRequest
	err := c.ShouldBindJSON(&mcReq)
	// 判断跳转链接是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &mcReq, c)
		return
	}

	var sender, receiver model.UserModel
	err = global.Db.First(&sender, mcReq.SenderID).Error
	if err != nil {
		response.FailWithMessage("发送人为空", c)
		return
	}
	err = global.Db.First(&receiver, mcReq.ReceiverID).Error
	if err != nil {
		response.FailWithMessage("收件人为空", c)
		return
	}

	err = global.Db.Create(&model.MessageModel{
		SenderID:         sender.ID,
		SenderNickName:   sender.NickName,
		SenderAvatar:     sender.Avatar,
		ReceiverID:       receiver.ID,
		ReceiverNickName: receiver.NickName,
		ReceiverAvatar:   receiver.Avatar,
		IsRead:           false,
		Content:          mcReq.Content,
	}).Error
	if err != nil {
		response.FailWithMessage("消息发送失败!", c)
		return
	}
	response.OKWithMessage("消息发送成功", c)
}
