package middleware

import (
	"capys/internal/model"
	"capys/internal/service"
	"capys/utility"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/net/ghttp"
)

type sMiddleware struct{}

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	token := r.GetHeader("Authorization")

	tokenValue, err := utility.ParseToken(r.GetCtx(), token)
	if err != nil {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeNotAuthorized.Code(),
			Message: err.Error(),
		})
		r.Exit()
	}

	if tokenValue.IsBan {
		r.Response.WriteJson(ghttp.DefaultHandlerResponse{
			Code:    gcode.CodeNotAuthorized.Code(),
			Message: "用户已被封禁",
		})
		r.Exit()
	}

	user := &model.Context{
		Id:    tokenValue.Id,
		Power: tokenValue.Power,
	}

	service.BizCtx().Init(r, user)

	r.Middleware.Next()

}

func init() {
	service.RegisterMiddleware(&sMiddleware{})
}
