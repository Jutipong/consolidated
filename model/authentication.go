package model

type Login struct {
	Username string `form:"Username" binding:"required"`
	Password string `form:"Password" binding:"required"`
}
