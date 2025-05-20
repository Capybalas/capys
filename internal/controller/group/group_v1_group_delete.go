package group

import (
	v1 "capys/api/group/v1"
	"capys/internal/service"
	"context"
)

func (c *ControllerV1) GroupDelete(ctx context.Context, req *v1.GroupDeleteReq) (res *v1.GroupDeleteRes, err error) {
	err = service.Group().Delete(ctx, req.Id)
	return
}
