package problem

import (
	"context"
	v1 "capys/api/problem/v1"
	"capys/internal/service"
)

func (c *ControllerV1) ProblemGetList(ctx context.Context, req *v1.ProblemGetListReq) (res *v1.ProblemGetListRes, err error) {
	out, err := service.Problem().GetList(ctx, req.BaseListInput)

	res = &v1.ProblemGetListRes{
		BaseListOut: out,
	}
	return
}
