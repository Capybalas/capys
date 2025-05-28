// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package problem

import (
	"context"

	"capys/api/problem/v1"
)

type IProblemV1 interface {
	ProblemGetOne(ctx context.Context, req *v1.ProblemGetOneReq) (res *v1.ProblemGetOneRes, err error)
	ProblemGetList(ctx context.Context, req *v1.ProblemGetListReq) (res *v1.ProblemGetListRes, err error)
	ProblemSubmit(ctx context.Context, req *v1.ProblemSubmitReq) (res *v1.ProblemSubmitRes, err error)
}
