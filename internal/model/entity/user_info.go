// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UserInfo is the golang structure for table user_info.
type UserInfo struct {
	UserId   uint64 `json:"user_id"   orm:"user_id"   description:""`      //
	Username string `json:"username"  orm:"username"  description:""`      //
	IdNumber string `json:"id_number" orm:"id_number" description:"学号/工号"` // 学号/工号
	Avatar   string `json:"avatar"    orm:"avatar"    description:""`      //
}
