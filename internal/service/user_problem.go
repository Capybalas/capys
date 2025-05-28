package service

import (
	"capys/internal/model/entity"
	"context"
)

type IUserProblem interface {
	GetOne(ctx context.Context, id any) (out *entity.UserProblem, err error)
	Create(ctx context.Context, in *entity.UserProblem) (err error)
	Update(ctx context.Context, in *entity.UserProblem) (err error)
}

var localUserProblem IUserProblem

func UserProblem() IUserProblem {
	if localUserProblem == nil {
		panic("IUserProblem接口未实现或注册")
	}
	return localUserProblem
}

func RegisterUserProblem(i IUserProblem) {
	localUserProblem = i
}
