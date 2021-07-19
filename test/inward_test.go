package test

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/domain/inward/model"
	"consolidated/domain/inward/repository"
	"consolidated/domain/inward/service"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Iwrm(t *testing.T) {
	config.InitTimeZone()
	base.InitMasterRule()
	repositoryMock := repository.NewRepositoryMock()
	service := service.NewService(repositoryMock)
	req := model.Request{}
	_, code, _ := service.Iwrm(&req)
	assert.NotNil(t, code)
	// if code != base.Successfully {
	// 	_rult := base.GetRule(code)
	// 	str := util.JsonSerialize(_rult)
	// 	t.Error(fmt.Sprintf("err code: %v", str))
	// }

	// assert.NotNil(t, code)
	// a := model.Response{
	// 	FeeAmount: 123,
	// }
	// b := model.Response{
	// 	FeeAmount: 123,
	// }
	// assert.Equal(t, a, b)
	// assert.Equal(t, TITLE, firstVideo.Title)
	// assert.Equal(t, DESCRIPTION, firstVideo.Description)
	// assert.Equal(t, URL, firstVideo.URL)

	// if len(err) > 0 {
	// 	t.Error(fmt.Sprintf("errs : %v", err))
	// }

}
