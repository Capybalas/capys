package model

type CreateUserInput struct {
	*RegisterUserInput
	Safe string
}

type Token struct {
	Power uint `json:"power"`
	IsBan bool `json:"is_ban"`
}

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Power uint   `json:"power"`
}

type UserInfo struct {
	UserId   string `json:"user_id"`
	Username string `json:"username"`
	IdNumber string `json:"id_number"`
	Avatar   string `json:"avatar"`
}
