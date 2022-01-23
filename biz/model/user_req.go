package model

type RegisterPostReq struct {
	Username   string `form:"username"`
	Password   string `form:"password"`
	RePassword string `form:"rePassword"`
}

type LoginPostReq struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
