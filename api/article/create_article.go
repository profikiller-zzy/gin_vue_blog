package article

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/article_service"
	"gin_vue_blog_AfterEnd/utils"
	"gin_vue_blog_AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

type CreateArticleRequest struct {
	Title    string      `json:"title" structs:"title" binding:"required" meg:"文章标题不能为空！"`                // 文章标题
	Keyword  string      `json:"keyword,omit(list)" structs:"keyword"`                                    // 关键字
	Abstract string      `json:"abstract" structs:"abstract"`                                             // 文章简介
	Content  string      `json:"content,omit(list)" structs:"content" binding:"required" meg:"文章内容不能为空！"` // 文章内容
	Category string      `json:"category" structs:"category"`                                             // 文章分类
	Source   string      `json:"source" structs:"source"`                                                 // 文章来源
	Link     string      `json:"link" structs:"link"`                                                     // 原文链接
	BannerID uint        `json:"banner_id" structs:"banner_id"`                                           // 文章封面id
	Tags     ctype.Array `json:"tags" structs:"tags"`                                                     // 文章标签
}

func (ArticleApi) CreateArticle(c *gin.Context) {
	var caReq CreateArticleRequest
	err := c.ShouldBindJSON(&caReq)
	if err != nil {
		response.FailBecauseOfParamError(err, &caReq, c)
		return
	}
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	userID := claims.UserID
	userNickName := claims.NickName

	if caReq.Abstract == "" { // 如果文章的简介为空，则默认截取文章中前50个字符作为简介
		caReq.Abstract = utils.TruncateText(caReq.Content, 50)
	}

	// 查对应banner_id对应的图片URL
	var bannerURL string
	err = global.Db.Model(&model.BannerModel{}).Where("id = ?", caReq.BannerID).Select("path").Scan(&bannerURL).Error
	if err != nil {
		response.FailWithMessage("指定的banner不存在", c)
		return
	}

	// 查询用户头像的URL
	var userAvatar string
	err = global.Db.Model(&model.UserModel{}).Where("id = ?", userID).Select("avatar").Scan(&userAvatar).Error
	if err != nil {
		response.FailWithMessage("用户不存在", c)
		return
	}

	// 实例化Article,将文章写入es
	now := time.Now().Format("2006-01-02 15:04:05")
	article := model.ArticleModel{
		CreatedAt:    now,
		UpdatedAt:    now,
		Title:        caReq.Title,
		Keyword:      caReq.Keyword,
		Abstract:     caReq.Abstract,
		Content:      caReq.Content,
		UserID:       userID,
		UserNickName: userNickName,
		UserAvatar:   userAvatar,
		Category:     caReq.Category,
		Source:       caReq.Source,
		Link:         caReq.Link,
		BannerID:     caReq.BannerID,
		BannerUrl:    bannerURL,
		Tags:         caReq.Tags,
	}
	err = article_service.InsertArticleToES(&article)
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage("文章写入ES服务器失败", c)
		return
	}
	response.OKWithMessage("文章上传成功！", c)
	return
}
