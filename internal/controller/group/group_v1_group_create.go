package group

import (
	v1 "capys/api/group/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error) {
	err = service.Group().Create(ctx, req.CreateGroupInput)
	return
}
