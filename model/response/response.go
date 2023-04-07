package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 封装了一些响应成功和失败的response
type Response struct {
	Code int    `json:"code"`
	Data any    `json:"data"`
	Msg  string `json:"msg"`
}

const (
	Success = 0
	Error   = 7
)

func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code: code,
		Data: data,
		Msg:  msg,
	})
}

func OK(data any, msg string, c *gin.Context) {
	Result(Success, data, msg, c)
}

func OKWithData(data any, c *gin.Context) {
	Result(Success, data, "操作成功", c)
}

func OKWithMessage(msg string, c *gin.Context) {
	Result(Success, map[string]interface{}{}, msg, c)
}

func FailWithCode(code ErrorCode, c *gin.Context) {
	msg, ok := CodeMessage[ErrorCode(code)]
	if ok {
		Result(int(code), map[string]interface{}{}, msg, c)
		return
	}
	Result(int(code), map[string]interface{}{}, "未知错误", c)
}

func Fail(c *gin.Context) {
	Result(Error, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(msg string, c *gin.Context) {
	Result(Error, map[string]interface{}{}, msg, c)
}
