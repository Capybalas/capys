package group

import (
	v1 "capys/api/group/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) GroupRemoveUser(ctx context.Context, req *v1.GroupRemoveUserReq) (res *v1.GroupRemoveUserRes, err error) {
	err = service.UserGroup().Delete(ctx, req.UserID, req.ID)
	return
}
