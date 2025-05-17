package service

import "github.com/gogf/gf/v2/net/ghttp"

type IMiddleware interface {
	Ctx(r *ghttp.Request)
}

var localMiddleware IMiddleware

func Middleware() IMiddleware {
	if localMiddleware == nil {
		panic("IMiddleware接口未实现或注册")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
