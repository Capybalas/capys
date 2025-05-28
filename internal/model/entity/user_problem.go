// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserProblem is the golang structure for table user_problem.
type UserProblem struct {
	UserId    uint64 `json:"user_id"    orm:"user_id"    description:""` //
	ProblemId uint64 `json:"problem_id" orm:"problem_id" description:""` //
	IsPass    bool   `json:"is_pass"    orm:"is_pass"    description:""` //
}
