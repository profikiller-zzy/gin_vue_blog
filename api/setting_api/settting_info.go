package setting_api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SettingInfoView 处理请求查看系统设置视图的函数
func (s SettingApi) SettingInfoView(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "ok",
	})
}
