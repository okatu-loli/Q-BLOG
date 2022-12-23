package errcode

const (
	SUCCESS = 200
	ERROR   = 500
	// code = 100x	用户模块错误
	ERROR_USERNAME         = 1001
	ERROR_PASSWORD         = 1002
	ERROR_USER_NOT_EXIST   = 1003
	ERROR_TOKEN_EXIST      = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WAONG      = 1006
	ERROR_TOKEN_TYPE_ERROR = 1007
	// code = 200x	文章模块错误
	ERROR_CATENAME_USED = 2001
	// code = 300x	分类模块错误
)

var codeMsg = map[int]string{
	SUCCESS:                "OK",
	ERROR:                  "FAIL",
	ERROR_USERNAME:         "用户名已存在",
	ERROR_PASSWORD:         "密码错误",
	ERROR_USER_NOT_EXIST:   "用户不存在",
	ERROR_TOKEN_EXIST:      "TOKEN不存在",
	ERROR_TOKEN_RUNTIME:    "TOKEN已过期",
	ERROR_TOKEN_WAONG:      "TOKEN不正确",
	ERROR_TOKEN_TYPE_ERROR: "TOKEN格式错误",
	ERROR_CATENAME_USED:    "分类已存在",
}

func GetErrMsg(code int) string {
	return codeMsg[code]
}
