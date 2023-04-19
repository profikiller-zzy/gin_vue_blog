package ad_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type AdRequest struct {
	Title     string `json:"title"  binding:"required" msg:"请输入标题" structs:"title"`
	Href      string `json:"href"  binding:"required,url" msg:"跳转链接非法" structs:"href"`             // 标识了这个字段必填，且为合法的URL
	ImagePath string `json:"image_path"  binding:"required,url" msg:"图片地址非法" structs:"image_path"` // 标识了这个字段必填，且为合法的URL
	IsShow    bool   `json:"is_show"  default:"false" msg:"选择是否展示" structs:"is_show"`
}

func (AdApi) AdCreateView(c *gin.Context) {
	var adReq AdRequest
	err := c.ShouldBindJSON(&adReq)
	// 判断跳转链接是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &adReq, c)
		return
	}

	// 需要判断广告是否重复，联合广告的标题和跳转链接
	var adModel model.AdModel
	err = global.Db.First(&adModel, "title = ? and href = ?", adReq.Title, adReq.Href).Error
	if err == nil { // 找到了符合条件的记录
		response.FailWithMessage("相同的广告已存在", c)
		return
	}

	err = global.Db.Create(&model.AdModel{
		Title:     adReq.Title,
		Href:      adReq.Href,
		ImagePath: adReq.ImagePath,
		IsShow:    adReq.IsShow,
	}).Error

	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("添加广告失败", c)
		return
	}
	response.OKWithMessage("添加广告成功", c)
}
