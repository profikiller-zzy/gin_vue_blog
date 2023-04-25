package middleware

import (
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
)

// JwtAuth 管理用户登录的中间件
func JwtAuth() gin.HandlerFunc {
	// 如何判断发送请求的是admin还是user
	// 从浏览器请求头中获取token，使用token判断是不是管理员
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("token")
		if tokenString == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.VerifyToken(tokenString)
		if err != nil {
			response.FailWithMessage("非法token", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}

func JwtAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("token")
		if tokenString == "" {
			response.FailWithMessage("未携带token", c)
			c.Abort()
			return
		}
		claims, err := jwts.VerifyToken(tokenString)
		if err != nil {
			response.FailWithMessage("非法token", c)
			c.Abort()
			return
		}
		if claims.Role != int(ctype.PermissionAdmin) { // 不是管理员
			response.FailWithMessage("没有权限", c)
			c.Abort()
			return
		}
		c.Set("claims", claims)
	}
}
