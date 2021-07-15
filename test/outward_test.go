package test

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/domain/outward/model"
	"consolidated/domain/outward/repository"
	"consolidated/domain/outward/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Branch(t *testing.T) {
	config.InitTimeZone()
	base.InitMasterRule()
	repositoryMock := repository.NewRepositoryMock()
	service := service.NewService(repositoryMock)

	req := model.Request{}
	_, code, _ := service.Branch(&req)
	assert.NotNil(t, code)
	// if code != base.Successfully {
	// 	_rult := base.GetRule(code)
	// 	str := util.JsonSerialize(_rult)
	// 	t.Error(fmt.Sprintf("err code: %v", str))
	// }

	// if len(err) > 0 {
	// 	t.Error(fmt.Sprintf("errs : %v", err))
	// }

}
