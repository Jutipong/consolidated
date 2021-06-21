package service

import (
	"consolidated/entity"
	"consolidated/repository"
)

func FindAll(u *[]entity.User) (err error) {
	if err = repository.FindAll(u); err != nil {
		return err
	}
	return nil
}

func FindID(u *entity.User, id string) (err error) {
	if err = repository.FindID(u, id); err != nil {
		return err
	}
	return nil
}

func AddNewCustomer(u *entity.User) (err error) {
	if err = repository.AddNewCustomer(u); err != nil {
		return err
	}
	return nil
}

func PutOneCustomer(u *entity.User, id string) (err error) {
	if err = repository.PutOneCustomer(u, id); err != nil {
		return err
	}
	return nil
}

func DeleteCustomer(u *entity.User, id string) (err error) {
	if err = repository.DeleteCustomer(u, id); err != nil {
		return err
	}
	return nil
}
