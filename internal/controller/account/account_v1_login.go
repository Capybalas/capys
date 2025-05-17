package account

import (
	v1 "capys/api/account/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	out, err := service.Account().Login(ctx, req.Account, req.Password)
	res = &v1.LoginRes{
		LoginOut: out,
	}
	return
}
