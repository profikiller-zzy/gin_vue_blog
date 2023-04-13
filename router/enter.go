package router

import (
	"gin_vue_blog_AfterEnd/global"
	"github.com/gin-gonic/gin"
)

type RGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	apiRouter := router.Group("/api/")

	apiRouterGroupApp := RGroup{
		RouterGroup: apiRouter,
	}
	apiRouterGroupApp.SettingRouter()
	apiRouterGroupApp.ImageRouter()
	return router
}
