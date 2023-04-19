package ad_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

func (AdApi) AdRemoveView(c *gin.Context) {
	var rmReq model.RemoveRequest
	var adList []model.AdModel
	var count int64 = 0

	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，请确认请求类型为JSON，error：%s", err.Error()), c)
		return
	}

	count = global.Db.Find(&adList, rmReq.IDList).RowsAffected
	if count == 0 { // 需要删除的图片ID没有在数据库中查到
		response.FailWithMessage("文件不存在", c)
		return
	}
	global.Db.Delete(&model.AdModel{}, rmReq.IDList)
	response.FailWithMessage(fmt.Sprintf("删除 %d 广告记录成功", count), c)
}
