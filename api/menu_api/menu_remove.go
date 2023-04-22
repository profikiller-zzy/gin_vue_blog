package menu_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemove(c *gin.Context) {
	var rmReq model.RemoveRequest
	var menuList []model.MenuModel
	var count int64 = 0

	err := c.ShouldBindJSON(&rmReq)
	if err != nil {
		response.FailWithMessage(fmt.Sprintf("参数绑定失败，error：%s", err.Error()), c)
		return
	}

	count = global.Db.Find(&menuList, rmReq.IDList).RowsAffected
	if count == 0 { // 需要删除的图片ID没有在数据库中查到
		response.FailWithMessage("菜单不存在", c)
		return
	}
	err = global.Db.Transaction(func(tx *gorm.DB) error {
		// 先将该菜单项的关联图片清空
		err = global.Db.Model(&menuList).Association("Banners").Clear()
		if err != nil {
			return err
		}
		// 再删除对应的菜单项列表
		err = global.Db.Delete(&model.MenuModel{}, rmReq.IDList).Error
		if err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(fmt.Sprintf("删除菜单失败, 错误信息:%s", err.Error()), c)
		return
	}
	response.FailWithMessage(fmt.Sprintf("删除 %d 个菜单成功", count), c)
}
