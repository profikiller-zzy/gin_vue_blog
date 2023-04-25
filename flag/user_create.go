package flag

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"gin_vue_blog_AfterEnd/model"
	"gin_vue_blog_AfterEnd/model/ctype"
	"gin_vue_blog_AfterEnd/service/user_service"
)

func CreateUser(permission string) {
	//
	var (
		nickName   string
		userName   string
		password   string
		rePassword string
		email      string
		err        error
	)
	fmt.Println("请输入昵称：")
	fmt.Scan(&nickName)
	for {
		fmt.Println("请输入用户名：")
		fmt.Scan(&userName)
		// 查询用户名是否存在
		var user model.UserModel
		err := global.Db.First(&user, "user_name = ?", userName).Error
		if err == nil {
			fmt.Println("用户名已存在，请重新输入：")
		} else {
			break
		}
	}
	fmt.Println("请输入密码：")
	fmt.Scan(&password)
	fmt.Println("请再次输入密码以确认：")
	fmt.Scan(&rePassword)
	// 校验二次密码
	for password != rePassword {
		fmt.Println("你输入确认密码有错，请重新输入：")
		fmt.Scan(&rePassword)
	}
	fmt.Println("请输入邮箱地址：(没有的话按回车直接跳过)")
	fmt.Scan(&email)
	var role ctype.Role = ctype.PermissionUser
	if permission == "admin" {
		role = ctype.PermissionAdmin
	}
	err = user_service.UserService{}.CreateUser(userName, nickName, password, role, email, "127.0.0.1")
	if err != nil {
		global.Log.Error(err.Error())
		return
	}
	global.Log.Info("创建成功!")
}
