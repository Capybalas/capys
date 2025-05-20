package service

import (
	"capys/internal/model"
	"capys/utility/casql"
	"context"
)

type IGroup interface {
	Create(ctx context.Context, in *model.CreateGroupInput) (err error)
	Update(ctx context.Context, id any, in *model.CreateGroupInput) (err error)
	GetOne(ctx context.Context, id any) (out *model.Group, err error)
	GetList(ctx context.Context, in *casql.BaseListInput) (out *casql.BaseListOut, err error)
	Delete(ctx context.Context, id any) (err error)
}

var localGroup IGroup

func Group() IGroup {
	if localGroup == nil {
		panic("IGroup接口未实现或注册")
	}
	return localGroup
}

func RegisterGroup(i IGroup) {
	localGroup = i
}
