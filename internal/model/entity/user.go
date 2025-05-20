// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id        uint64      `json:"id"         orm:"id"         description:""`                       //
	Email     string      `json:"email"      orm:"email"      description:""`                       //
	Phone     string      `json:"phone"      orm:"phone"      description:""`                       //
	Password  string      `json:"password"   orm:"password"   description:""`                       //
	Safe      string      `json:"safe"       orm:"safe"       description:""`                       //
	Power     uint        `json:"power"      orm:"power"      description:"0-一般用户 1-学生 2-教师 3-管理员"` // 0-一般用户 1-学生 2-教师 3-管理员
	IsBan     bool        `json:"is_ban"     orm:"is_ban"     description:"是否封禁"`                   // 是否封禁
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`                       //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`                       //
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`                       //
}
