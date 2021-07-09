package inwardfree

import (
	"consolidated/features/inwardfee/controller"
	"consolidated/utils"

	"github.com/gin-gonic/gin"
)

type InwardFeeAPI struct{}

func (u *InwardFeeAPI) Setup(r *gin.Engine) {
	g := r.Group("fee/inward/v1/")
	{
		_user := controller.InwardController{}
		g.POST("iwrm", utils.JwtVerify, _user.InwardFee)
	}
}
