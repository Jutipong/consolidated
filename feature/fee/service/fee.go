package service

import (
	"consolidated/feature/fee/model"
	"consolidated/feature/fee/repository"
)

var _inwardFee = repository.Fee{}

func FindById(u *[]model.Request) (err error) {
	if err = _inwardFee.FindById(u); err != nil {
		return err
	}
	return nil
}
