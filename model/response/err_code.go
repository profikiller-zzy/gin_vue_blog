package response

type ErrorCode int

const (
	SettingsError = ErrorCode(1001)
)

var CodeMessage = map[ErrorCode]string{
	SettingsError: "系统错误",
}
