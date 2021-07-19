package repository

import "consolidated/domain/outward/model"

type repositoryMock struct {
	dataMock []model.Response
}

func NewRepositoryMock() IRepository {
	dataMock := []model.Response{
		{
			FeeType:     "FeeType 1111",
			FeeAmount:   1111,
			FeeCCY:      "FeeCCY 1111",
			FeeCategory: "FeeCategory 1111",
			FeeRefID:    "FeeRefID 1111",
		},
		{
			FeeType:     "FeeType 2222",
			FeeAmount:   2222,
			FeeCCY:      "FeeCCY 2222",
			FeeCategory: "FeeCategory 2222",
			FeeRefID:    "FeeRefID 2222",
		},
	}
	return &repositoryMock{dataMock: dataMock}
}

func (m *repositoryMock) Inquiry(req *model.Request) ([]model.Response, error) {
	return m.dataMock, nil
}
