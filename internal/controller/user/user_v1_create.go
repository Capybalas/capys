package user

import (
	"context"
	"fmt"

	v1 "capys/api/user/v1"
	"capys/internal/service"
)

func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error) {
	content, err := service.User().Create()
	fmt.Println(content)
	return
}
