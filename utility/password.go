package utility

import (
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/google/uuid"
)

func MakeSafeNumber() string {
	return uuid.New().String()
}

func EncryptPassword(in, safe string) (password string) {
	password, _ = gmd5.Encrypt(in + safe)
	return
}
