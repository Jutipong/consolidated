package test

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/domain/fee/model"
	"consolidated/domain/fee/repository"
	"consolidated/domain/fee/service"
	"consolidated/util"
	"fmt"
	"testing"
)

func Test_PromotionCode(t *testing.T) {
	config.InitTimeZone()
	base.InitMasterRule()
	repositoryMock := repository.NewRepositoryMock()
	service := service.NewService(repositoryMock)

	req := model.Request{}
	_, code, _ := service.PromotionCode(&req)

	if code != base.Successfully {
		_rult := base.GetRule(code)
		str := util.JsonSerialize(_rult)
		t.Error(fmt.Sprintf("err code: %v", str))
	}

	// if len(err) > 0 {
	// 	t.Error(fmt.Sprintf("errs : %v", err))
	// }

}
