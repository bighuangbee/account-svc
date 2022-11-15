# 错误码
	错误码       Reason                        描述

	0         Success                        操作成功
	1         InternalError                  内部错误
	2         DbError                        数据库错误
	7         InvalidParameter               无效的参数
	9         Timeout                        请求超时
	11        Unauthenticated                jwt token无效
	12        RecordNotFound                 记录未找到
	13        RecordIsExists                 记录已经存在
	14        RedisExistError                RedisExist 错误
	15        RedisGetError                  RedisGet 错误
	30        NoAccess                       无访问权限
	50        GrpcError                      前面预留慧飞旧服务错误码使用
	51        HttpError                      http调用失败
	52        CaptchaError                   验证码校验失败
	53        AccountPwdError                账号密码校验失败
