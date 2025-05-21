package v1

import (
	"github.com/gogf/gf/v2/frame/g"

)

type GroupAddUserReq struct {
	g.Meta `path:"/{id}/add" method:"post" tags:"用户分组" summary:"添加一个用户到分组"`
	ID     string `json:"id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
	UserID string `json:"user_id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
}

type GroupAddUserRes struct{}

type GroupRemoveUserReq struct {
	g.Meta `path:"/{id}/remove" method:"post" tags:"用户分组" summary:"分组移除某个用户"`
	ID     string `json:"id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
	UserID string `json:"user_id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
}

type GroupRemoveUserRes struct{}
