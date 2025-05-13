package v1

import "github.com/gogf/gf/v2/frame/g"

type RegisterReq struct {
	g.Meta   `path:"/" method:"post" tags:"用户管理" summary:"创建用户"`
	Email    string `json:"email" v:"required|email#您必须提供邮箱地址|您提供的邮箱地址格式不正确"`
	Phone    string `json:"phone" v:"required|phone#您必须提供手机号|您提供的手机号格式不正确"`
	Password string `json:"password" v:"required|min-length:6|max-length:10|password3#重复密码必须填写|密码最小6位|密码最多10位|密码必须包含大小写字母、数字和特殊字符" dc:"登录密码"`
}

type RegisterRes struct{}
