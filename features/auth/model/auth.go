package model

type UserLogin struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type UserRequest struct {
	UserId   string
	UserName string
}
