package model

type CreateGroupInput struct {
	GroupName string `json:"group_name"`
}

type Group struct {
	Id        uint64 `json:"id"        orm:"id"         description:""`
	GroupName string `json:"group_name" orm:"group_name" description:"分组名称"`
}
