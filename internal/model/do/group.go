// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Group is the golang structure of table group for DAO operations like Where/Data.
type Group struct {
	g.Meta    `orm:"table:group, do:true"`
	Id        interface{} //
	GroupName interface{} // 分组名称
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
