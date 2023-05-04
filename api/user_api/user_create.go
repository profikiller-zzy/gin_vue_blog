package user_api

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/user_service"
	"github.com/gin-gonic/gin"
)

type UserCreateRequest struct {
	NickName string     `json:"nick_name" binding:"required" msg:"请输入昵称"`  // 昵称
	UserName string     `json:"user_name" binding:"required" msg:"请输入用户名"` // 用户名
	Password string     `json:"password" binding:"required" msg:"请输入密码"`   // 密码
	Role     ctype.Role `json:"role" binding:"required" msg:"请选择权限"`       // 权限  1 管理员  2 普通用户  3 游客
}

func (UserApi) UserCreateView(c *gin.Context) {
	// TODO 后续需要对该函数进行逻辑调整，因为登录之后对某人发消息，即本人是发送人
	var ucReq UserCreateRequest
	if err := c.ShouldBindJSON(&ucReq); err != nil {
		response.FailBecauseOfParamError(err, &ucReq, c)
		return
	}
	err := user_service.UserService{}.CreateUser(ucReq.UserName, ucReq.NickName, ucReq.Password, ucReq.Role, "", c.ClientIP())
	if err != nil {
		global.Log.Error(err)
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OKWithMessage(fmt.Sprintf("用户%s创建成功!", ucReq.UserName), c)
	return
}
