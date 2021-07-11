package api

import (
	"consolidated/features/auth"
	inwardfree "consolidated/features/inwardfee"
	outwardfree "consolidated/features/outwardfee"
	"consolidated/features/user"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	//## authentication
	_authApi := auth.AuthAPI{}
	_authApi.Setup(r)

	//## User
	_userApi := user.UserAPI{}
	_userApi.Setup(r)

	//## InwardFee
	_inwardFee := inwardfree.InwardFeeAPI{}
	_inwardFee.Setup(r)

	//## OutwardFee
	_outwardFee := outwardfree.OutwardFeeAPI{}
	_outwardFee.Setup(r)
}
