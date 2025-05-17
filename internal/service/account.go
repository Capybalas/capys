package service

import (
	"capys/internal/model"
	"context"
)

type IAccount interface {
	// Register 注册用户
	Register(ctx context.Context, user *model.RegisterUserInput, userInfo *model.RegisterUserInfoInput) (err error)
	Login(ctx context.Context, account, password string) (out *model.LoginOut, err error)
}

var localAccount IAccount

func Account() IAccount {
	if localAccount == nil {
		panic("IAccount接口未实现或注册")
	}
	return localAccount
}

func RegisterAccount(i IAccount) {
	localAccount = i
}
