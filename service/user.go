package service

import (
	"consolidated/config"
	"consolidated/entity"
)

func FindAll(c *[]entity.User) (err error) {
	if err = config.DB.Table("User").Find(&c).Error; err != nil {
		return err
	}
	return nil
}

func FindID(c *entity.User, id string) (err error) {
	if err = config.DB.Table("User").Where("id = ?", id).First(c).Error; err != nil {
		return err
	}
	return nil
}

func AddNewCustomer(c *entity.User) (err error) {
	if err = config.DB.Table("User").Create(c).Error; err != nil {
		return err
	}
	return nil
}

func PutOneCustomer(c *entity.User, id string) (err error) {
	config.DB.Table("User").Where("id = ?", id).Save(c)
	return nil
}

func DeleteCustomer(c *entity.User, id string) (err error) {
	config.DB.Table("User").Where("id = ?", id).Delete(c)
	return nil
}
