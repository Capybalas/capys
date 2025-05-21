package service

import (
	"capys/internal/model/entity"
	"context"

)

type IUserGroup interface {
	Create(ctx context.Context, userID, groupID any) (err error)
	Delete(ctx context.Context, userID, groupID any) (err error)
	GetOne(ctx context.Context, userID, groupID any) (out *entity.UserGroup, err error)
}

var localUserGroup IUserGroup

func UserGroup() IUserGroup {
	if localUserGroup == nil {
		panic("IUserGroup接口未实现或注册")
	}
	return localUserGroup
}

func RegisterUserGroup(i IUserGroup) {
	localUserGroup = i
}
