package problem

import (
	"context"
	v1 "capys/api/problem/v1"
	"capys/internal/service"
)

func (c *ControllerV1) ProblemDelete(ctx context.Context, req *v1.ProblemDeleteReq) (res *v1.ProblemDeleteRes, err error) {
	err = service.Problem().Delete(ctx, req.Id)
	return
}
