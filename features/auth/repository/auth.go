package repository

type Auth struct {
	Username string `validate:"required"`
	Password string `validate:"required"`
}

type UserRequest struct {
	SystemId     string
	EmpCode      string
	User         string
	Permissions  string
	TransationId string `gorm:"type:uuid;"`
}
