package user_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

// UserRemoveView 这个函数目前还不完善，删除用户之前需要先将指定用户的全部依赖删除
func (UserApi) UserRemoveView(c *gin.Context) {
	var rmReq model.RemoveRequest
	var userList []model.UserModel
	var count int64 = 0

	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，请确认请求类型为JSON，error：%s", err.Error()), c)
		return
	}

	count = global.Db.Find(&userList, rmReq.IDList).RowsAffected
	if count == 0 { // 需要删除的图片ID没有在数据库中查到
		response.FailWithMessage("文件不存在", c)
		return
	}
	// TODO:删除用户关联的消息、评论、用户收藏文章、用户发布文章
	global.Db.Delete(&model.AdModel{}, rmReq.IDList)
	response.FailWithMessage(fmt.Sprintf("删除 %d 个用户成功", count), c)
}
