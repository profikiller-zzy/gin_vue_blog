package user_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/model/response"
	"github.com/gin-gonic/gin"
)

type UserRoleRequest struct {
	Role     ctype.Role `json:"role" binding:"required,oneof=1 2 3 4" msg:"输入正确的权限等级"`
	NickName string     `json:"nick_name"`
	UserID   uint       `json:"user_id" binding:"required" msg:"请确认需要修改的用户ID"`
}

// UserUpdateRoleView 用户权限变更
func (UserApi) UserUpdateRoleView(c *gin.Context) {
	var urReq UserRoleRequest
	err := c.ShouldBindJSON(&urReq)
	// 判断参数是否合法
	if err != nil {
		response.FailBecauseOfParamError(err, &urReq, c)
		return
	}

	var userModel model.UserModel
	err = global.Db.First(&userModel, urReq.UserID).Error
	if err != nil {
		response.FailWithMessage("输入的用户ID错误", c)
		return
	}
	err = global.Db.Model(&userModel).Updates(map[string]interface{}{
		"role":      urReq.Role,
		"nick_name": urReq.NickName,
	}).Error
	if err != nil {
		global.Log.Error(err.Error())
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("用权限成功变更为%s,用户昵称变更为'%s'", urReq.Role.String(), urReq.NickName), c)
}
