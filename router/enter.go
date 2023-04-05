package router

import (
	"gin_vue_blog_AfterEnd/global"
	"github.com/gin-gonic/gin"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	apiRouter := router.Group("/api/")

	apiRouterGroupApp := RouterGroup{
		RouterGroup: apiRouter,
	}
	apiRouterGroupApp.SettingRouter()

	return router
}
