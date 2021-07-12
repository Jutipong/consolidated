package service

import (
	"consolidated/feature/inwardfee/model"
	"consolidated/feature/inwardfee/repository"
)

var _inwardFee = repository.InwardFee{}

type InwardFee struct{}

func (i *InwardFee) FindById(u *[]model.Request) (err error) {
	if err = _inwardFee.FindById(u); err != nil {
		return err
	}
	return nil
}
