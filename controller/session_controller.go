package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// IsSessionExist session 是否存在
func IsSessionExist(c *gin.Context) bool {
	s := sessions.Default(c)
	loginUser := s.Get("loginUser")
	if loginUser != nil {
		return true
	} else {
		return false
	}
}

// GetSession 拿到 session
func GetSession(c *gin.Context, str string) interface{} {
	s := sessions.Default(c)
	return s.Get(str)
}
