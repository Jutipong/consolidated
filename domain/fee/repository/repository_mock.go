package repository

import "consolidated/domain/fee/model"

type repositoryMock struct {
	dataMock []model.Response
}

func NewRepositoryMock() IRepository {
	dataMock := []model.Response{
		{
			ProRefID:      "ProRefID: 1",
			PromotionType: "PromotionType : 1",
			Amount:        1111,
			CCY:           "CCY : 1"},
		{
			ProRefID:      "ProRefID: 2",
			PromotionType: "PromotionType : 2",
			Amount:        2222,
			CCY:           "CCY : 2"},
	}
	return repositoryMock{dataMock: dataMock}
}

func (mc repositoryMock) Inquiry(req *model.Request) ([]model.Response, error) {
	return mc.dataMock, nil
}
