package tag_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type TagRequest struct {
	Title string `json:"title"  binding:"required" msg:"请输入标签标题" structs:"title"`
}

func (TagApi) TagCreateView(c *gin.Context) {
	var tagReq TagRequest
	err := c.ShouldBindJSON(&tagReq)
	// 判断跳转链接是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &tagReq, c)
		return
	}

	// 需要判断广告是否重复，联合广告的标题和跳转链接
	var tagModel model.TagModel
	err = global.Db.First(&tagModel, "title = ?", tagReq.Title).Error
	if err == nil { // 找到了符合条件的记录
		response.FailWithMessage("同样标签的文章已经存在", c)
		return
	}

	err = global.Db.Create(&model.AdModel{
		Title: tagReq.Title,
	}).Error

	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("添加标签失败", c)
		return
	}
	response.OKWithMessage("添加标签成功", c)
}
