package user_api

import (
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/model/response"
	"gin_vue_blog_AfterEnd/service/common_service"
	"gin_vue_blog_AfterEnd/utils/desen"
	"gin_vue_blog_AfterEnd/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	model.UserModel
	RoleID int `json:"role_id"`
}

type UserListRequest struct {
	model.PageInfo
	Role int `json:"role" form:"role"`
}

// UserListView 为管理员返回全部用户信息
func (UserApi) UserListView(c *gin.Context) {
	_claims, _ := c.Get("claims")
	claims := _claims.(*jwts.CustomClaims)
	var pageModel model.PageInfo
	err := c.ShouldBindQuery(&pageModel)
	if err != nil {
		response.FailWithCode(response.ParameterError, c)
		return
	}
	var userRepList = make([]UserResponse, 0)
	var userList = make([]model.UserModel, 0)
	var count int64
	// 对图片列表进行分页
	userList, count, err = common_service.PagingList(model.UserModel{}, common_service.PageInfoDebug{
		PageInfo: pageModel,
		Debug:    true,
	})

	for _, user := range userList {
		if ctype.Role(claims.Role) != ctype.PermissionAdmin {
			// 非管理员，不显示用户名
			user.UserName = ""
		}
		// 数据脱敏
		user.Tel = desen.DesensitizationTel(user.Tel)
		user.Email = desen.DesensitizationEmail(user.Email)
		userRepList = append(userRepList, UserResponse{
			UserModel: user,
			RoleID:    int(user.Role),
		})
	}
	response.OKWithPagingData(userRepList, count, c)
	return
}
