// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// UserProblemRecord is the golang structure for table user_problem_record.
type UserProblemRecord struct {
	Id        uint64      `json:"id"         orm:"id"         description:""` //
	UserId    uint64      `json:"user_id"    orm:"user_id"    description:""` //
	ProblemId uint64      `json:"problem_id" orm:"problem_id" description:""` //
	Answer    string      `json:"answer"     orm:"answer"     description:""` //
	IsPass    bool        `json:"is_pass"    orm:"is_pass"    description:""` //
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""` //
}
