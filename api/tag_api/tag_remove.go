package tag_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

func (TagApi) TagRemoveView(c *gin.Context) {
	var rmReq model.RemoveRequest
	var tagList []model.TagModel
	var count int64 = 0

	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，请确认请求类型为JSON，error：%s", err.Error()), c)
		return
	}

	count = global.Db.Find(&tagList, rmReq.IDList).RowsAffected
	if count == 0 { // 需要删除的图片ID没有在数据库中查到
		response.FailWithMessage("指定的标签不存在", c)
		return
	}
	// 先将需要删除的标签的关联记录删除
	global.Db.Delete(&model.TagModel{}, rmReq.IDList)
	response.FailWithMessage(fmt.Sprintf("删除 %d 条标签记录成功", count), c)
}
