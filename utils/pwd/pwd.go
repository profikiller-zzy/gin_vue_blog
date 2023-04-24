package pwd

import (
	"gin_vue_blog_AfterEnd/global"
	"golang.org/x/crypto/bcrypt"
)

// BcryptPw 加密密码
func BcryptPw(pwd string) string {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		global.Log.Error(err.Error())
	}
	return string(hashPwd)
}

func VerifyPwd(pwd string, hashPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))
	if err != nil {
		global.Log.Warnln(err.Error())
		return false
	}
	return true
}
