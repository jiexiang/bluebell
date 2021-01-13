package models

type ParamSignUp struct {
	Username   string `form:"username" binding:"required"`
	Password   string `form:"password" binding:"required"`
	RePassword string `form:"re_password" binding:"required"`
}
