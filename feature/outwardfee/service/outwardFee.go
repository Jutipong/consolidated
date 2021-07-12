package service

import (
	"consolidated/feature/outwardfee/model"
	"consolidated/feature/outwardfee/repository"
)

var _inwardFee = repository.OutwardFee{}

type InwardFee struct{}

func (i *InwardFee) FindById(u *[]model.Request) (err error) {
	if err = _inwardFee.FindById(u); err != nil {
		return err
	}
	return nil
}
