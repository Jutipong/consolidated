package service

import (
	"consolidated/base"
	"consolidated/domain/fee/model"
	"consolidated/domain/fee/repository"
)

type IService interface {
	PromotionCode(req *model.Request) (model.Response, float32, []string)
}

type service struct {
	repository repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return service{repository: repo}
}

func (s service) PromotionCode(req *model.Request) (result model.Response, code float32, err []string) {
	code, err = req.Validate()
	if code != base.Successfully {
		return result, code, err
	}

	return result, base.Successfully, err
}
