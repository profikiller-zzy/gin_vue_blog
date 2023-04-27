package user_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/plugin/email"
	"gin_vue_blog_AfterEnd/utils/jwts"
	"gin_vue_blog_AfterEnd/utils/random"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
)

type EmailCaptcha struct {
	Email string  `json:"email" binding:"required,email" msg:"邮箱非法"`
	Code  *string `json:"code"`
}

// EmailBindingView 用户绑定邮箱，第一阶段，输入邮箱获取验证码；第二阶段，用户输入绑定邮箱和验证码，后台校验验证码通过后，邮箱绑定成功
func (UserApi) EmailBindingView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	var ec EmailCaptcha
	err := c.ShouldBindJSON(&ec)
	if err != nil {
		response.FailBecauseOfParamError(err, &ec, c)
		return
	}

	// 用户绑定邮箱，第一阶段，输入邮箱获取验证码
	// 获取会话
	session := sessions.Default(c)
	// 验证码为空，用户第一次注册邮箱，后台给邮箱发验证码
	// 生成4位验证码， 将生成的验证码存入session
	if ec.Code == nil {
		code := random.RandCode(4)
		// 写入session
		session.Set("valid_code", code)
		// 设置验证码的有效期限为五分钟
		session.Options(sessions.Options{
			MaxAge: int(5 * time.Minute.Seconds())})
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("会话保存出错，请重试!", c)
			return
		}
		err = email.NewCode().Send(ec.Email, "你的验证码是 "+code)
		if err != nil {
			global.Log.Error(err)
		}
		response.OKWithMessage("验证码已发送，请查收", c)
		return
	} else {
		// 第二次请求 验证验证码
		validCode := session.Get("valid_code")
		err = session.Save()
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("会话保存出错，请重试!", c)
			return
		}
		// 判断是否过期
		if validCode == nil {
			response.FailWithMessage("验证码已过期，请重新获取", c)
			return
		}
		// 校验验证码
		if validCode != *ec.Code {
			response.FailWithMessage("验证码错误", c)
			return
		}
		// 修改用户的邮箱
		var user model.UserModel
		err = global.Db.Take(&user, claims.UserID).Error
		if err != nil {
			response.FailWithMessage("用户不存在", c)
			return
		}
		// 第一次的邮箱，和第二次的邮箱也要做一致性校验
		err = global.Db.Model(&user).Update("email", ec.Email).Error
		if err != nil {
			global.Log.Error(err)
			response.FailWithMessage("绑定邮箱失败，发送验证码的邮箱和需要绑定的邮箱不一致", c)
			return
		}
		// 完成绑定
		response.OKWithMessage("邮箱绑定成功", c)
		return
	}
}
