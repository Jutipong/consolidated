package repository

import (
	"consolidated/config"
	"consolidated/entity"
)

func FindAll(u *[]entity.User) (err error) {
	rows, err := config.DB.Table("User").Find(&u).Rows()
	if err = config.DB.Table("User").Find(&u).Error; err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func FindID(u *entity.User, id string) (err error) {
	rows, err := config.DB.Table("User").Where("id = ?", id).First(u).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func AddNewCustomer(u *entity.User) (err error) {
	rows, err := config.DB.Table("User").Create(u).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func PutOneCustomer(u *entity.User, id string) (err error) {
	rows, err := config.DB.Table("User").Where("id = ?", id).Save(u).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}

func DeleteCustomer(u *entity.User, id string) (err error) {
	rows, err := config.DB.Table("User").Where("id = ?", id).Delete(u).Rows()
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
