package router

import (
	"gin_vue_blog_AfterEnd/global"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type RGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	err := router.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		global.Log.Warnln(err.Error())
	}

	// use ginSwagger middleware to serve the API docs
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiRouter := router.Group("/api/")

	apiRouterGroupApp := RGroup{
		RouterGroup: apiRouter,
	}
	apiRouterGroupApp.SettingRouter()
	apiRouterGroupApp.ImageRouter()
	apiRouterGroupApp.AdRouter()
	apiRouterGroupApp.MenuRouter()
	apiRouterGroupApp.UserRouter()
	apiRouterGroupApp.TagRouter()
	apiRouterGroupApp.MessageRouter()
	apiRouterGroupApp.ArticleRouter()

	return router
}
