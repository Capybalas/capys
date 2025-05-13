// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// UserInfo is the golang structure of table user_info for DAO operations like Where/Data.
type UserInfo struct {
	g.Meta   `orm:"table:user_info, do:true"`
	UserId   interface{} //
	Username interface{} //
	IdNumber interface{} // 学号/工号
	Avatar   interface{} //
}
