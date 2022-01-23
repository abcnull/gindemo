package util

import (
	"strconv"
)

var SecretKey []byte

func InitSecretKey() {
	SecretKey = []byte("some string")
}

// ChangeSecretKey 改变 secretKey
func ChangeSecretKey() {
	r := RandNum()
	SecretKey = []byte(strconv.Itoa(r))
}

func GetSecretKey() []byte {
	return SecretKey
}
