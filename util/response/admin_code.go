package response

const (
	SUCCESS = 200
	ERROR = 500
	INVALID_PARAMS = 400
	ERROR_AUTH_CHECK_TOKEN_FAIL = 401		// token认证失败

	// token相关
	ERROR_CREATE_TOKEN_FAIL = 1007			// token创建失败

	ERROR_CREATE_FAIL 		= 1001 			// 创建失败
	ERROR_UPDATE_FAIL 		= 1002			// 更新失败
	ERROR_DELETE_FAIL 		= 1003			// 删除失败
	ERROR_NOT_FOUND			= 1004			// 未找到数据
	ERROR_EXIST 			= 1005			// 数据已存在

	ERROR_USER_NOT_FOUND 	= 2001			// 用户未找到
	ERROR_USER_EXIST 		= 2002			// 用户已存在
	ERROR_USER_ACCOUNT		= 2003			// 用户名或密码错误
)