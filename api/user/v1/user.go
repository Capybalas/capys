package v1

import (
	"capys/internal/model"

	"github.com/gogf/gf/v2/frame/g"
)

type UserGetOneReq struct {
	g.Meta `path:"/{id}" method:"get" tags:"用户" summary:"获取一条用户"`
	Id     string `in:"path" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
}

type UserGetOneRes struct {
	*model.User
	*model.UserInfo
}
