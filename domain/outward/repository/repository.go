package repository

import (
	"consolidated/domain/outward/model"
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
	return &repository{db: db}
}

func (r *repository) Inquiry(req *model.Request) ([]model.Response, error) {
	var res []model.Response
	return res, errors.New("Inquiry")
}
