package service

import (
	"capys/internal/model"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
)

type IBizCtx interface {
	Init(r *ghttp.Request, customCtx *model.Context)
	GetUser(ctx context.Context) (user *model.Context)
	GetUserId(ctx context.Context) (id string)
}

var localBizCtx IBizCtx

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("IBizCtx接口未实现或注册")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}
