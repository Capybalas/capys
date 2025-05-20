package group

import (
	v1 "capys/api/group/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) GroupGetOne(ctx context.Context, req *v1.GroupGetOneReq) (res *v1.GroupGetOneRes, err error) {
	out, err := service.Group().GetOne(ctx, req.Id)

	res = &v1.GroupGetOneRes{
		Group: out,
	}

	return
}
