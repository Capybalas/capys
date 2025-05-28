package admin

import (
	v1 "capys/api/admin/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) ProblemCreate(ctx context.Context, req *v1.ProblemCreateReq) (res *v1.ProblemCreateRes, err error) {
	err = service.Problem().Create(ctx, req.InProblem)
	return
}
