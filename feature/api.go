package api

import (
	"consolidated/feature/auth"
	inwardfree "consolidated/feature/inwardfee"
	outwardfree "consolidated/feature/outwardfee"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {

	//## authentication
	_authApi := auth.AuthAPI{}
	_authApi.Setup(r)

	//## InwardFee
	_inwardFee := inwardfree.InwardFeeAPI{}
	_inwardFee.Setup(r)

	//## OutwardFee
	_outwardFee := outwardfree.OutwardFeeAPI{}
	_outwardFee.Setup(r)
}
