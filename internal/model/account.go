package model

import "github.com/gogf/gf/v2/os/gtime"

type RegisterUserInput struct {
	Email      string `json:"email" v:"required|email#您必须提供邮箱地址|您提供的邮箱地址格式不正确"`
	Phone      string `json:"phone" v:"required|phone#您必须提供手机号|您提供的手机号格式不正确"`
	Password   string `json:"password" v:"required|min-length:6|max-length:20|password3#密码必须填写|密码最小6位|密码最多20位|密码必须包含大小写字母、数字和特殊字符" dc:"登录密码"`
	RePassword string `json:"re_password" v:"required|same:password#您必须提供确认密码|两次输入的密码不一致" dc:"确认密码"`
}

type RegisterUserInfoInput struct {
	Username string `json:"username" v:"required#您必须提供昵称|min-length:2#昵称最少是2个字|max-length:9#昵称最多是9个字" dc:"昵称"`
	IdNumber string `json:"id_number" v:"min-length:8#学号/工号最少8个字符|max-length:16#学号/工号最多16个字符" dc:"学号/工号"`
}

type LoginOut struct {
	Id       string     `json:"id"`
	Username string     `json:"username"`
	IdNumber string     `json:"id_number"`
	Token    string     `json:"token"`
	Expire   gtime.Time `json:"expire"`
}
