package repository

import (
	"consolidated/config"
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id         uuid.UUID `gorm:"type:uuid;"`
	Name       string
	Last       string
	CreateDate time.Time
	CreateBy   string
	UpdateDate time.Time
	UpdateBy   string
	IsActive   bool
}

func FindAll(u *[]User) (err error) {
	rows, err := config.DB.Table("User").Find(&u).Rows()
	if err = config.DB.Table("User").Find(&u).Error; err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func FindID(u *User, id string) (err error) {
	if err := config.DB.Table("User").Where("Id = ?", strings.ToUpper(id)).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCustomer(u *User) (err error) {
	rows, err := config.DB.Table("User").Create(u).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func PutOneCustomer(u *User, id string) (err error) {
	rows, err := config.DB.Table("User").Where("id = ?", id).Save(u).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func DeleteCustomer(u *User, id string) (err error) {
	rows, err := config.DB.Table("User").Where("id = ?", id).Delete(u).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
