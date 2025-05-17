package user

import (
	"capys/internal/dao"
	"capys/internal/model"
	"capys/internal/model/do"
	"capys/internal/model/entity"
	"capys/internal/service"
	"context"
	"strconv"

	"github.com/gogf/gf/v2/frame/g"

)

type sUser struct{}

// ReloadPower implements service.IUser.
func (s *sUser) ReloadPower(ctx context.Context, id string) (err error) {
	user := &model.Token{}
	err = dao.User.Ctx(ctx).WherePri(id).Scan(user)
	g.Redis().Set(ctx, "user."+id, &user)
	g.Redis().Expire(ctx, "user."+id, 2*60*60) // 2小时过期
	return
}

func (s *sUser) GetOne(ctx context.Context, id any) (user *model.User, err error) {
	err = dao.User.Ctx(ctx).WherePri(id).Scan(&user)
	return
}

// GetOneByAccount implements service.IUser.
func (s *sUser) GetOneByAccount(ctx context.Context, phone, email string) (user *entity.User, err error) {
	err = dao.User.Ctx(ctx).WhereOr("phone", phone).WhereOr("email", email).Scan(&user)
	return
}

func (s *sUser) Create(ctx context.Context, in *model.CreateUserInput) (id string, err error) {
	idInt64, err := dao.User.Ctx(ctx).Data(do.User{
		Email:    in.Email,
		Phone:    in.Phone,
		Password: in.Password,
		Safe:     in.Safe,
	}).InsertAndGetId()

	if err != nil {
		return
	}

	id = strconv.Itoa(int(idInt64))
	return
}

func init() {
	service.RegisterUser(&sUser{})
}
