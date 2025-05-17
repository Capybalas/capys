package bizctx

import (
	"capys/internal/consts"
	"capys/internal/model"
	"capys/internal/service"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"

)

type sBizCtx struct{}

// GetUser implements service.IBizCtx.
func (s *sBizCtx) GetUser(ctx context.Context) (user *model.Context) {
	user = ctx.Value(consts.ContextKey).(*model.Context)
	return
}

func (s *sBizCtx) GetUserId(ctx context.Context) (id string) {
	id = s.GetUser(ctx).Id
	return
}

func (s *sBizCtx) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(consts.ContextKey, customCtx)
}

func init() {
	service.RegisterBizCtx(&sBizCtx{})
}
