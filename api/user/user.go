// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"

	"capys/api/user/v1"
)

type IUserV1 interface {
	UserGetOne(ctx context.Context, req *v1.UserGetOneReq) (res *v1.UserGetOneRes, err error)
}
