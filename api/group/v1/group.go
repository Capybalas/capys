package v1

import (
	"capys/internal/model"
	"capys/utility/casql"

	"github.com/gogf/gf/v2/frame/g"

)

type GroupCreateReq struct {
	g.Meta `path:"/" method:"post" tags:"用户分组" summary:"创建用户分组"`
	*model.CreateGroupInput
}

type GroupCreateRes struct{}

type GroupUpdateReq struct {
	g.Meta `path:"/{id}" method:"put" tags:"用户分组" summary:"修改一条用户分组"`
	Id     string `v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
	*model.CreateGroupInput
}

type GroupUpdateRes struct{}

type GroupGetOneReq struct {
	g.Meta `path:"/{id}" method:"get" tags:"用户分组" summary:"获取一条用户分组"`
	Id     string `v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
}

type GroupGetOneRes struct {
	*model.Group
}

type GroupGetListReq struct {
	g.Meta `path:"/" method:"get" tags:"用户分组" summary:"获取用户分组"`
	*casql.BaseListInput
}

type GroupGetListRes struct {
	*casql.BaseListOut
}

type GroupDeleteReq struct {
	g.Meta `path:"/{id}" method:"delete" tags:"用户分组" summary:"删除用户分组"`
	Id     string `v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
}

type GroupDeleteRes struct{}
