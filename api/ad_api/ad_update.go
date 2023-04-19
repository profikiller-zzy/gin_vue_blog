package ad_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/fatih/structs"
	"github.com/gin-gonic/gin"
)

func (AdApi) AdUpdateView(c *gin.Context) {
	var adReq AdRequest
	id := c.Param("id")
	err := c.ShouldBindJSON(&adReq)
	// 判断跳转链接是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &adReq, c)
		return
	}

	var adModel model.AdModel
	err = global.Db.First(&adModel, "id = ?", id).Error
	if err != nil { // 没有找到符合条件的记录
		response.FailWithMessage("该广告不存在", c)
		return
	}

	// 结构体转map
	adReqMap := structs.Map(&adReq)
	err = global.Db.Model(&adModel).Updates(adReqMap).Error

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage("修改广告成功", c)
}
