// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Group is the golang structure for table group.
type Group struct {
	Id        uint64      `json:"id"         orm:"id"         description:""`     //
	GroupName string      `json:"group_name" orm:"group_name" description:"分组名称"` // 分组名称
	CreatedAt *gtime.Time `json:"created_at" orm:"created_at" description:""`     //
	UpdatedAt *gtime.Time `json:"updated_at" orm:"updated_at" description:""`     //
	DeletedAt *gtime.Time `json:"deleted_at" orm:"deleted_at" description:""`     //
}
