package group

import (
	v1 "capys/api/group/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) GroupGetList(ctx context.Context, req *v1.GroupGetListReq) (res *v1.GroupGetListRes, err error) {
	out, err := service.Group().GetList(ctx, req.BaseListInput)

	res = &v1.GroupGetListRes{
		BaseListOut: out,
	}

	return
}
