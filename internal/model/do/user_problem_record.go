// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProblemRecord is the golang structure of table user_problem_record for DAO operations like Where/Data.
type UserProblemRecord struct {
	g.Meta    `orm:"table:user_problem_record, do:true"`
	Id        interface{} //
	UserId    interface{} //
	ProblemId interface{} //
	Answer    interface{} //
	IsPass    interface{} //
	CreatedAt *gtime.Time //
}
