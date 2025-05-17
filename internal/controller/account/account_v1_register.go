package account

import (
	v1 "capys/api/account/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	err = service.Account().Register(ctx, req.RegisterUserInput, req.RegisterUserInfoInput)
	return
}
