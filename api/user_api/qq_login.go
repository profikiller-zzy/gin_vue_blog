package user_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/model/response"
	qq_plugin "gin_vue_blog_AfterEnd/plugin/qq"
	"gin_vue_blog_AfterEnd/utils/jwts"
	"gin_vue_blog_AfterEnd/utils/pwd"
	"gin_vue_blog_AfterEnd/utils/random"
	"github.com/gin-gonic/gin"
)

func (UserApi) QQLoginView(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		response.FailWithMessage("没有code", c)
		return
	}
	qqInfo, err := qq_plugin.NewQQLogin(code)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	openID := qqInfo.OpenID
	// 根据openID判断用户是否存在
	// 根据openID判断用户是否存在
	var userModel model.UserModel
	err = global.Db.Take(&userModel, "token = ?", openID).Error
	if err != nil {
		// 不存在，就注册
		userModel = model.UserModel{
			NickName:   qqInfo.Nickname,
			UserName:   openID,
			Password:   pwd.BcryptPw(random.RandPassword(16)), // 随机生成16位密码
			Avatar:     qqInfo.Avatar,
			Addr:       "", // 需要通过IP地址算出地理位置
			Token:      openID,
			IP:         c.ClientIP(),
			Role:       ctype.PermissionUser,
			SignStatus: ctype.SignQQ,
		}
		err = global.Db.Create(&userModel).Error
		if err != nil {
			global.Log.Error(err.Error())
			response.FailWithMessage("注册失败", c)
			return
		}
	}
	// 登录操作
	tokenString, err := jwts.GenerateToken(jwts.JwtPayLoad{
		UserID:   userModel.ID,
		NickName: userModel.NickName,
		Role:     int(userModel.Role),
		Avatar:   userModel.Avatar,
	})
	//ip, addr := utils.GetAddrByGin(c)
	response.OKWithData(tokenString, c)
}
