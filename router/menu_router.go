package router

import "gin_vue_blog_AfterEnd/api"

func (r RGroup) MenuRouter() {
	menuApiApp := api.ApiGroupApp.MenuApi
	r.POST("/menu/", menuApiApp.CreateMenuView)
	r.GET("/menu/", menuApiApp.MenuListView)
	r.GET("/menu/:id", menuApiApp.MenuDetailView)
	r.GET("/menu_name/", menuApiApp.MenuNameList)
	r.PUT("/menu/:id", menuApiApp.MenuUpdateView)
	r.DELETE("/menu/", menuApiApp.MenuRemove)
}
