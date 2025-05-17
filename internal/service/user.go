package service

import (
	"capys/internal/model"
	"capys/internal/model/entity"
	"context"
)

type IUser interface {
	Create(ctx context.Context, in *model.CreateUserInput) (id string, err error)
	GetOneByAccount(ctx context.Context, phone, email string) (user *entity.User, err error)
	GetOne(ctx context.Context, id any) (user *model.User, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("IUser接口未实现或注册")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
