package admin

import (
	v1 "capys/api/admin/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) ProblemGetList(ctx context.Context, req *v1.ProblemGetListReq) (res *v1.ProblemGetListRes, err error) {
	out, err := service.Problem().GetList(ctx, req.BaseListInput)

	res = &v1.ProblemGetListRes{
		BaseListOut: out,
	}
	return
}
