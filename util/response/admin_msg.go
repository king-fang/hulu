package response

var MsgFlags = map[int]string {
	SUCCESS 					:"success",
	ERROR 						:"服务器错误",
	INVALID_PARAMS 				: "参数错误",

	ERROR_AUTH_CHECK_TOKEN_FAIL :"Token鉴权失败",
	ERROR_CREATE_TOKEN_FAIL 	:"Token创建失败",

	ERROR_CREATE_FAIL			:"创建失败",
	ERROR_UPDATE_FAIL			:"更新失败",
	ERROR_DELETE_FAIL			:"删除失败",
	ERROR_NOT_FOUND				:"未找到",
	ERROR_EXIST					:"已存在",

	ERROR_USER_NOT_FOUND		:"用户不存在",
	ERROR_USER_EXIST			:"用户已存在",
	ERROR_USER_ACCOUNT			:"用户名或密码错误",
}

func GetMsg(code int,message string) string  {
	msg,ok := MsgFlags[code]
	if ok {
		return message + msg
	}
	return MsgFlags[ERROR]
}
