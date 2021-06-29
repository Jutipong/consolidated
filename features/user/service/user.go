package service

import "consolidated/features/user/repository"

func FindAll(u *[]repository.User) (err error) {
	if err = repository.FindAll(u); err != nil {
		return err
	}
	return nil
}

func FindID(u *repository.User, id string) (err error) {
	if err = repository.FindID(u, id); err != nil {
		return err
	}
	return nil
}

func AddNewCustomer(u *repository.User) (err error) {
	if err = repository.AddNewCustomer(u); err != nil {
		return err
	}
	return nil
}

func PutOneCustomer(u *repository.User, id string) (err error) {
	if err = repository.PutOneCustomer(u, id); err != nil {
		return err
	}
	return nil
}

func DeleteCustomer(u *repository.User, id string) (err error) {
	if err = repository.DeleteCustomer(u, id); err != nil {
		return err
	}
	return nil
}
