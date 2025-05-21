// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package group

import (
	"context"

	"capys/api/group/v1"
)

type IGroupV1 interface {
	GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error)
	GroupUpdate(ctx context.Context, req *v1.GroupUpdateReq) (res *v1.GroupUpdateRes, err error)
	GroupGetOne(ctx context.Context, req *v1.GroupGetOneReq) (res *v1.GroupGetOneRes, err error)
	GroupGetList(ctx context.Context, req *v1.GroupGetListReq) (res *v1.GroupGetListRes, err error)
	GroupDelete(ctx context.Context, req *v1.GroupDeleteReq) (res *v1.GroupDeleteRes, err error)
	GroupAddUser(ctx context.Context, req *v1.GroupAddUserReq) (res *v1.GroupAddUserRes, err error)
	GroupRemoveUser(ctx context.Context, req *v1.GroupRemoveUserReq) (res *v1.GroupRemoveUserRes, err error)
}
