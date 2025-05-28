package v1

import (
	"capys/internal/model"
	"capys/utility/casql"

	"github.com/gogf/gf/v2/frame/g"
)

type ProblemGetOneReq struct {
	g.Meta `path:"/problem/{id}" method:"get" tags:"题库" summary:"用户获取某题"`
	Id     string `json:"id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
}

type ProblemGetOneRes struct {
	*model.OutProblem
}

type ProblemGetListReq struct {
	g.Meta `path:"/problem" method:"get" tags:"题库" summary:"用户获取所有题目"`
	*casql.BaseListInput
}

type ProblemGetListRes struct {
	*casql.BaseListOut
}

type ProblemSubmitReq struct {
	g.Meta `path:"/problem/{id}" method:"post" tags:"题库" summary:"用户提交答案"`
	Id     string `json:"id" v:"required#无效id|IsNumeric#id不正确|min:1#id不正确"`
	*model.UserSubmitProblem
}

type ProblemSubmitRes struct {
}
