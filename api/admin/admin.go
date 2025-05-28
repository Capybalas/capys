// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package admin

import (
	"context"

	"capys/api/admin/v1"
)

type IAdminV1 interface {
	ProblemCreate(ctx context.Context, req *v1.ProblemCreateReq) (res *v1.ProblemCreateRes, err error)
	ProblemUpdate(ctx context.Context, req *v1.ProblemUpdateReq) (res *v1.ProblemUpdateRes, err error)
	ProblemGetOne(ctx context.Context, req *v1.ProblemGetOneReq) (res *v1.ProblemGetOneRes, err error)
	ProblemGetList(ctx context.Context, req *v1.ProblemGetListReq) (res *v1.ProblemGetListRes, err error)
	ProblemDelete(ctx context.Context, req *v1.ProblemDeleteReq) (res *v1.ProblemDeleteRes, err error)
}
