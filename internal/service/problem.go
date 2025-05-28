package service

import (
	"capys/internal/model"
	"capys/utility/casql"
	"context"
)

type IProblem interface {
	Create(ctx context.Context, in *model.InProblem) (err error)
	Update(ctx context.Context, id any, in *model.InProblem) (err error)
	GetOne(ctx context.Context, id any) (out *model.OutProblem, err error)
	GetList(ctx context.Context, in *casql.BaseListInput) (out *casql.BaseListOut, err error)
	Delete(ctx context.Context, id any) (err error)
	Run(ctx context.Context, id any, userCode *model.UserSubmitProblem) (out *model.RunInfo, err error)
	AddNumber(ctx context.Context, id any, in *model.AddNumber) (err error)
}

var localProblem IProblem

func Problem() IProblem {
	if localProblem == nil {
		panic("IProblem接口未实现或注册")
	}
	return localProblem
}

func RegisterProblem(i IProblem) {
	localProblem = i
}
