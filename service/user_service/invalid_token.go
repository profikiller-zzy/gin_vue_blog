package user_service

import (
	"fmt"
	"gin_vue_blog_AfterEnd/global"
	"time"
)

// AddInvalidTokenToBlackList 将失效的token加入黑名单
func (UserService) AddInvalidTokenToBlackList(tokenString string, duration time.Duration) error {
	err := global.Redis.Set(fmt.Sprintf("logout_%s", tokenString), "", duration).Err()
	return err
}

// CheckTokenInBlackList 判断token是否在黑名单中(判断该token是否失效)
func (UserService) CheckTokenInBlackList(tokenString string) (bool, error) {
	// 判断该token是否在redis黑名单中
	res, err := global.Redis.Exists(fmt.Sprintf("logout_%s", tokenString)).Result()
	if err != nil {
		global.Log.Error(err.Error())
		return false, err
	}
	if res == 1 { // 该token存在黑名单中
		return true, nil
	}
	return false, nil
}
