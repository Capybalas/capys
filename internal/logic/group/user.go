package group

import (
	"capys/internal/dao"
	"capys/internal/model/do"
	"capys/internal/model/entity"
	"capys/internal/service"
	"context"

)

type sUserGroup struct{}

func (s *sUserGroup) Create(ctx context.Context, userID any, groupID any) (err error) {
	_, err = dao.UserGroup.Ctx(ctx).Data(do.UserGroup{
		UserId:  userID,
		GroupId: groupID,
	}).Insert()
	return
}

func (s *sUserGroup) Delete(ctx context.Context, userID any, groupID any) (err error) {
	_, err = dao.UserGroup.Ctx(ctx).Where(do.UserGroup{
		UserId:  userID,
		GroupId: groupID,
	}).Delete()
	return
}

func (s *sUserGroup) GetOne(ctx context.Context, userID any, groupID any) (out *entity.UserGroup, err error) {
	err = dao.UserGroup.Ctx(ctx).Where(do.UserGroup{
		UserId:  userID,
		GroupId: groupID,
	}).Scan(&out)
	return
}

func init() {
	service.RegisterUserGroup(&sUserGroup{})
}
