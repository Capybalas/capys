package problem

import (
	v1 "capys/api/problem/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) ProblemUpdate(ctx context.Context, req *v1.ProblemUpdateReq) (res *v1.ProblemUpdateRes, err error) {
	err = service.Problem().Update(ctx, req.Id, req.InProblem)
	return
}
