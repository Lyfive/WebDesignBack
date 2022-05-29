/**
 @author: 李映飞
 @date:   2022/4/28
 @note:
**/
package e

var MsgFlags = map[int]string{
	SUCCESS:                   "ok",
	ERROR:                     "fail",
	INVALID_PARAMS:            "请求参数错误",
	CONTAINING_SENSITIVE_WORD: "包含敏感词",

	ERROR_EXIST_STUDENT:    "已存在该学生",
	ERROR_NOT_EXIST_NUMBER: "学号不存在",
	ERROR_NOT_EXIST_CLASS:  "班级不存在",
	ERROR_NOT_EXIST_COURSE: "课程不存在",
	ERROR_EXIST_GRADE:      "成绩已存在",

	ERROR_AUTH_CHECK_TOKEN_FAIL:      "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT:   "Token已超时",
	ERROR_AUTH_TOKEN:                 "Token生成失败",
	ERROR_AUTH:                       "Token错误",
	ERROR_INSUFFICIENT_ACCESS_RIGHTS: "访问权限不足",
	ERROR_EXIST_USER:                 "已存在该用户",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
