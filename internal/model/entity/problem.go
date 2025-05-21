// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Problem is the golang structure for table problem.
type Problem struct {
	Id                  uint64      `json:"id"                   orm:"id"                   description:""`                    //
	Title               string      `json:"title"                orm:"title"                description:"简单的题目"`               // 简单的题目
	Topic               string      `json:"topic"                orm:"topic"                description:"详细题目"`                // 详细题目
	Category            uint        `json:"category"             orm:"category"             description:"1 html"`              // 1 html
	Difficulty          uint        `json:"difficulty"           orm:"difficulty"           description:"1 入门 2简单 3中等 4难 5困难"` // 1 入门 2简单 3中等 4难 5困难
	Answer              string      `json:"answer"               orm:"answer"               description:""`                    //
	PassingNumber       uint        `json:"passing_number"       orm:"passing_number"       description:"通过人数"`                // 通过人数
	ParticipationNumber uint        `json:"participation_number" orm:"participation_number" description:"参与人数"`                // 参与人数
	CreatedAt           *gtime.Time `json:"created_at"           orm:"created_at"           description:""`                    //
	UpdatedAt           *gtime.Time `json:"updated_at"           orm:"updated_at"           description:""`                    //
	DeletedAt           *gtime.Time `json:"deleted_at"           orm:"deleted_at"           description:""`                    //
}
