package repository

import (
	"consolidated/config"
	"time"
)

type AuthUser struct {
	ID         int64
	Username   string
	Password   string
	CreateDate time.Time
	UpdateDate time.Time
}

func (auth *AuthUser) FindByUserName() (err error) {
	if err := config.DB.Table("auth_user").Where("Username = ?", auth.Username).First(&auth).Error; err != nil {
		return err
	}
	return nil
}
