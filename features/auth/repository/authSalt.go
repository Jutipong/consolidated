package repository

import (
	"consolidated/config"
	"time"
)

type AuthSalt struct {
	ID         int64
	Salt       string
	CreateDate time.Time
	UpdateDate time.Time
}

func (salt *AuthSalt) FindById() (err error) {
	if err := config.DB.Table("auth_salt").Where("ID=?", salt.ID).First(salt).Error; err != nil {
		return err
	}
	return err
}
