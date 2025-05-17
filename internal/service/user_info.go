package service

import (
	"capys/internal/model"
	"context"
)

type IUserInfo interface {
	Create(ctx context.Context, id string, in *model.RegisterUserInfoInput) (err error)
	GetOne(ctx context.Context, id any) (out *model.UserInfo, err error)
}

var localUserInfo IUserInfo

func UserInfo() IUserInfo {
	if localUserInfo == nil {
		panic("IUserInfo接口未实现或注册")
	}
	return localUserInfo
}

func RegisterUserInfo(i IUserInfo) {
	localUserInfo = i
}
