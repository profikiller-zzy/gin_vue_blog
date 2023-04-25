package user_api

import (
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service"
	"gin_vue_blog_AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
	"time"
)

// UserLogoutView 用户注销，通过使用redis维护jwt黑名单来解决jwt的有效性问题
func (UserApi) UserLogoutView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)

	tokenString := c.Request.Header.Get("token")

	// 需要现在距离过期时间还有多久，以这个计算出的时间来当作该记录的有效时间
	expiresAt := time.Unix(claims.ExpiresAt, 0)
	duration := expiresAt.Sub(time.Now())
	err := service.ServiceApp.UserServiceApp.AddInvalidTokenToBlackList(tokenString, duration)
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage("注销失败", c)
		return
	}
	response.OKWithMessage("注销成功", c)
}
