package group

import (
	v1 "capys/api/group/v1"
	"capys/internal/service"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
)

func (c *ControllerV1) GroupAddUser(ctx context.Context, req *v1.GroupAddUserReq) (res *v1.GroupAddUserRes, err error) {

	out, err := service.UserGroup().GetOne(ctx, req.UserID, req.ID)
	if err != nil {
		return
	}

	if out != nil {
		err = gerror.New("用户已在该组内")
		return
	}

	err = service.UserGroup().Create(ctx, req.UserID, req.ID)
	return
}
