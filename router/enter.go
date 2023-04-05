package router

import (
	"gin_vue_blog_AfterEnd/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	router.GET("/user/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	return router
}
