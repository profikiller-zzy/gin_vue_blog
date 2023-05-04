package message_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
)

type MessageRecordRequest struct {
	UserID uint `json:"user_id" binding:"required" msg:"请输入查询的用户id"`
}

func (MessageApi) MessageRecordView(c *gin.Context) {
	var meReq MessageRecordRequest
	err := c.ShouldBindJSON(&meReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &meReq, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var _messageList []model.MessageModel
	var messageList = make([]model.MessageModel, 0)
	// 先查出与自己有关的消息
	global.Db.Order("created_at asc").
		Find(&_messageList, "sender_id = ? or receiver_id = ", claims.UserID, claims.UserID)
	// 再从与自己有关的消息查出与该用户有关的消息，将其加入列表中
	for _, model := range _messageList {
		if model.ReceiverID == meReq.UserID || model.SenderID == meReq.UserID {
			messageList = append(messageList, model)
		}
	}

	// TODO 点开消息，里面的每一条消息，都从未读变成已读

	response.OKWithData(messageList, c)
}
