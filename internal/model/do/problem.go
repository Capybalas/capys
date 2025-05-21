// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Problem is the golang structure of table problem for DAO operations like Where/Data.
type Problem struct {
	g.Meta              `orm:"table:problem, do:true"`
	Id                  interface{} //
	Title               interface{} // 简单的题目
	Topic               interface{} // 详细题目
	Category            interface{} // 1 html
	Difficulty          interface{} // 1 入门 2简单 3中等 4难 5困难
	Answer              interface{} //
	PassingNumber       interface{} // 通过人数
	ParticipationNumber interface{} // 参与人数
	CreatedAt           *gtime.Time //
	UpdatedAt           *gtime.Time //
	DeletedAt           *gtime.Time //
}
