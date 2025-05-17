package v1

import (
	"capys/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type RegisterReq struct {
	g.Meta `path:"/" method:"post" tags:"用户管理" summary:"创建用户"`
	*model.RegisterUserInput
	*model.RegisterUserInfoInput
}

type RegisterRes struct{}

type LoginReq struct {
	g.Meta   `path:"/login" method:"post" tags:"用户管理" summary:"用户登录"`
	Account  string `json:"account" v:"required#您必须提供账号|min-length:2#账号最少2个字符|max-length:20#账号最多20个字符" dc:"账号"`
	Password string `json:"password" v:"required#您必须提供密码|min-length:6#密码最少6位|max-length:10#密码最多10位" dc:"登录密码"`
}

type LoginRes struct {
	*model.LoginOut
}
