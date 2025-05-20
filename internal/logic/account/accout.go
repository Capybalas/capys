package account

import (
	"capys/internal/model"
	"capys/internal/service"
	"capys/utility"
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

type sAccount struct{}

// Login implements service.IAccount.
func (s *sAccount) Login(ctx context.Context, account string, password string) (out *model.LoginOut, err error) {
	// user := &entity.User{}

	user, err := service.User().GetOneByAccount(ctx, account, account)
	if err != nil {
		return
	}

	if user == nil {
		err = errors.New("用户不存在")
		return
	}

	if user.IsBan {
		err = errors.New("用户已被封禁")
		return
	}

	userPassword := utility.EncryptPassword(password, user.Safe)
	if userPassword != user.Password {
		err = errors.New("密码错误")
		return
	}

	// 登录成功
	userInfo, err := service.UserInfo().GetOne(ctx, user.Id)
	if err != nil {
		return
	}

	id := strconv.Itoa(int(user.Id))
	currentTime := gtime.Now()

	ExpireTime := currentTime.Add(time.Duration(2) * time.Hour)

	tokenStruct := &model.Token{
		Power: user.Power,
		IsBan: user.IsBan,
	}

	token := utility.NewToken(id)

	if value, _ := g.Redis().Get(ctx, "user."+id); value.IsNil() {
		g.Redis().Set(ctx, "user."+id, tokenStruct)
		g.Redis().Expire(ctx, "user."+id, 7200)
	}

	g.Redis().Set(ctx, token, id)
	g.Redis().Expire(ctx, token, 7200)

	out = &model.LoginOut{
		Id:       id,
		Username: userInfo.Username,
		IdNumber: userInfo.IdNumber,
		Token:    token,
		Expire:   *ExpireTime,
	}

	return

}

func (s *sAccount) Register(ctx context.Context, in *model.RegisterUserInput, userInfo *model.RegisterUserInfoInput) (err error) {

	user, err := service.User().GetOneByAccount(ctx, in.Phone, in.Email)
	if err != nil {
		return
	}

	if user != nil {
		err = errors.New("您的手机号或邮箱已被使用")
		return
	}

	createUserInput := &model.CreateUserInput{}
	createUserInput.RegisterUserInput = in

	createUserInput.Safe = utility.MakeSafeNumber()
	createUserInput.Password = utility.EncryptPassword(in.Password, createUserInput.Safe)
	id, err := service.User().Create(ctx, createUserInput)
	if err != nil {
		return
	}

	err = service.UserInfo().Create(ctx, id, userInfo)

	return
}

func init() {
	service.RegisterAccount(&sAccount{})
}
