package util

import (
	"math/rand"
	"time"
)

// RandNum 生成随机数
func RandNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100000000)
}
