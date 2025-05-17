package cmd

import (
	"capys/internal/controller/account"
	"capys/internal/controller/user"
	"capys/internal/service"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()

			s.Group("/", func(baseRouter *ghttp.RouterGroup) {
				baseRouter.Middleware(ghttp.MiddlewareHandlerResponse)

				baseRouter.Bind(
					account.NewV1(),
				)

				// 需要登陆
				baseRouter.Group("", func(authRouter *ghttp.RouterGroup) {
					authRouter.Middleware(service.Middleware().Ctx)

					authRouter.Group("/user", func(userRouter *ghttp.RouterGroup) {
						userRouter.Bind(
							user.NewV1(),
						)
					})
				})

			})

			s.Run()
			return nil
		},
	}
)
