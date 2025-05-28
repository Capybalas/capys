// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserProblem is the golang structure of table user_problem for DAO operations like Where/Data.
type UserProblem struct {
	g.Meta    `orm:"table:user_problem, do:true"`
	UserId    interface{} //
	ProblemId interface{} //
	IsPass    interface{} //
}
