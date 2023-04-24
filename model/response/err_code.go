package response

type ErrorCode int

const (
	SettingsError  = ErrorCode(1001)
	ParameterError = ErrorCode(1002)
)

var CodeMessage = map[ErrorCode]string{
	SettingsError:  "系统错误",
	ParameterError: "参数错误",
}
