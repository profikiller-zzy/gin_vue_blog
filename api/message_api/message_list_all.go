package message_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/common_service"
	"github.com/gin-gonic/gin"
)

func (MessageApi) MessageListAllView(c *gin.Context) {
	// 作为管理员，可以看见所有用户的聊天列表；
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，error：%s", err.Error()), c)
		return
	}

	var messageList = make([]model.MessageModel, 0)
	var count int64
	messageList, count, err = common_service.PagingList(model.MessageModel{}, common_service.PageInfoDebug{
		PageInfo: pageModel,
		Debug:    true,
	})
	response.OKWithPagingData(messageList, count, c)
}
