package problem

import (
	"context"
	v1 "capys/api/problem/v1"
	"capys/internal/service"
)

func (c *ControllerV1) ProblemCreate(ctx context.Context, req *v1.ProblemCreateReq) (res *v1.ProblemCreateRes, err error) {
	err = service.Problem().Create(ctx, req.InProblem)
	return
}
