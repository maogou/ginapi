package errcode

//公共的业务码
var (
	Success                   = NewError(0, "成功")
	NotFound                  = NewError(404, "找不到页面")
	MethodNotAllow            = NewError(405, "方法不允许访问")
	ServerError               = NewError(10000000, "服务内部错误")
	InvalidParams             = NewError(10000001, "入参错误")
	ResourceNotFound          = NewError(10000002, "资源找不到")
	UnauthorizedAuthNotExist  = NewError(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(10000007, "请求过多")
	TokenInvalidParams        = NewError(10000008, "token入参错误")
)
