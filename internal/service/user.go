package service

type IUser interface {
	Create() (content string, err error)
}

var localUser IUser

func User() IUser {
	if localUser == nil {
		panic("IUser接口未实现或注册")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
