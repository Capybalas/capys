package user

import (
	v1 "capys/api/user/v1"
	"capys/internal/service"
	"context"
	"errors"
	"fmt"
)

func (c *ControllerV1) UserGetOne(ctx context.Context, req *v1.UserGetOneReq) (res *v1.UserGetOneRes, err error) {

	fmt.Printf("\n%+v\n", service.BizCtx().GetUser(ctx))

	user, err := service.User().GetOne(ctx, req.Id)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("用户不存在")
		return
	}

	userInfo, err := service.UserInfo().GetOne(ctx, user.Id)
	if err != nil {
		return
	}

	res = &v1.UserGetOneRes{
		User:     user,
		UserInfo: userInfo,
	}

	return
}
