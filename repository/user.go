package repository

import (
	"consolidated/config"
	"consolidated/entity"
)

func FindAll(u *[]entity.User) (err error) {
	if err = config.DB.Table("User").Find(&u).Error; err != nil {
		return err
	}
	return nil
}

func FindID(u *entity.User, id string) (err error) {
	if err = config.DB.Table("User").Where("id = ?", id).First(u).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCustomer(u *entity.User) (err error) {
	if err = config.DB.Table("User").Create(u).Error; err != nil {
		return err
	}
	return nil
}

func PutOneCustomer(u *entity.User, id string) (err error) {
	config.DB.Table("User").Where("id = ?", id).Save(u)
	return nil
}

func DeleteCustomer(u *entity.User, id string) (err error) {
	config.DB.Table("User").Where("id = ?", id).Delete(u)
	return nil
}
