package problem

import (
	v1 "capys/api/problem/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) ProblemGetOne(ctx context.Context, req *v1.ProblemGetOneReq) (res *v1.ProblemGetOneRes, err error) {
	out, err := service.Problem().GetOne(ctx, req.Id)
	res = &v1.ProblemGetOneRes{
		OutProblem: out,
	}
	return
}
