package utility

import (
	"capys/internal/model"
	"capys/internal/service"
	"context"
	"encoding/base64"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

)

func NewToken(id string) string {
	// 生成一个新的token
	return base64.StdEncoding.EncodeToString([]byte(id + MakeSafeNumber()))
}

func ParseToken(ctx context.Context, token string) (user *model.Context, err error) {

	if token == "" {
		err = gerror.New("请登录")
		return
	}

	value, err := g.Redis().Get(ctx, token)
	if err != nil {
		return
	}

	if value.IsNil() {
		err = gerror.New("登录凭证已过期，请重新登录")
		return
	}

	powerValue, err := g.Redis().Get(ctx, "user."+value.String())
	if err != nil {
		return
	}

	if powerValue.IsNil() {
		service.User().ReloadPower(ctx, value.String())
	}

	power := &model.Token{}
	powerValue.Scan(&power)

	user = &model.Context{
		Id:    value.String(),
		Power: power.Power,
	}

	return
}
