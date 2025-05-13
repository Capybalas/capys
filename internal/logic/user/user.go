package user

import "capys/internal/service"

type sUser struct{}

// Create implements service.IUser.
func (s *sUser) Create() (content string, err error) {
	return "user", nil
}

func init() {
	service.RegisterUser(&sUser{})
}
