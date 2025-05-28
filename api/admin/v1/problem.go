package v1

import (
	"capys/internal/model"
	"capys/utility/casql"

	"github.com/gogf/gf/v2/frame/g"
)

type ProblemCreateReq struct {
	g.Meta `path:"/problem" method:"post" tags:"题库" summary:"创建题目"`
	*model.InProblem
}

type ProblemCreateRes struct{}

type ProblemUpdateReq struct {
	g.Meta `path:"/problem/{id}" method:"put" tags:"题库" summary:"修改一条题目"`
	Id     string `json:"id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
	*model.InProblem
}

type ProblemUpdateRes struct{}

type ProblemGetOneReq struct {
	g.Meta `path:"/problem/{id}" method:"get" tags:"题库" summary:"获取一条题目"`
	Id     string `json:"id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
}

type ProblemGetOneRes struct {
	*model.OutProblem
}

type ProblemGetListReq struct {
	g.Meta `path:"/problem" method:"get" tags:"题库" summary:"获取题目"`
	*casql.BaseListInput
}

type ProblemGetListRes struct {
	*casql.BaseListOut
}

type ProblemDeleteReq struct {
	g.Meta `path:"/problem/{id}" method:"delete" tags:"题库" summary:"删除题目"`
	Id     string `json:"id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
}

type ProblemDeleteRes struct{}
