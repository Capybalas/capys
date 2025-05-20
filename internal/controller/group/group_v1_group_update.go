package group

import (
	v1 "capys/api/group/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) GroupUpdate(ctx context.Context, req *v1.GroupUpdateReq) (res *v1.GroupUpdateRes, err error) {
	err = service.Group().Update(ctx, req.Id, req.CreateGroupInput)
	return
}
