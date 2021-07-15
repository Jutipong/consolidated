package repository

import (
	"consolidated/domain/fee/model"
	"errors"

	"gorm.io/gorm"
)

type IRepository interface {
	Inquiry(req *model.Request) ([]model.Response, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) IRepository {
	return repository{db: db}
}

func (f repository) Inquiry(req *model.Request) ([]model.Response, error) {
	var result []model.Response
	return result, errors.New("Inquiry")
}
