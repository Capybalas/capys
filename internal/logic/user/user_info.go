package user

import (
	"capys/internal/dao"
	"capys/internal/model"
	"capys/internal/model/do"
	"capys/internal/service"
	"context"
)

type sUserInfo struct{}

func (s *sUserInfo) GetOne(ctx context.Context, id any) (out *model.UserInfo, err error) {
	err = dao.UserInfo.Ctx(ctx).WherePri(id).Scan(&out)
	return
}

func (s *sUserInfo) Create(ctx context.Context, id string, in *model.RegisterUserInfoInput) (err error) {
	_, err = dao.UserInfo.Ctx(ctx).Insert(do.UserInfo{
		UserId:   id,
		Username: in.Username,
		IdNumber: in.IdNumber,
	})
	return
}

func init() {
	service.RegisterUserInfo(&sUserInfo{})
}
