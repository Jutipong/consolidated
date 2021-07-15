package service

import (
	"consolidated/base"
	"consolidated/domain/inward/model"
	"consolidated/domain/inward/repository"
)

type IService interface {
	Iwrm(req *model.Request) ([]model.Response, float32, []string)
}

type service struct {
	repo repository.IRepository
}

func NewService(repo repository.IRepository) IService {
	return service{repo: repo}
}

func (s service) Iwrm(req *model.Request) (result []model.Response, code float32, errs []string) {
	code, errs = req.Validate()
	if code != base.Successfully {
		return result, code, errs
	}

	result, err := s.repo.Inquiry(req)
	if errs != nil {
		errs = append(errs, err.Error())
		return result, base.DataNotFound, errs
	}

	return result, base.Successfully, errs
}
