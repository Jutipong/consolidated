package inwardfree

import (
	"consolidated/feature/inwardfee/handler"
	"consolidated/util"

	"github.com/gin-gonic/gin"
)

type InwardFeeAPI struct{}

func (u *InwardFeeAPI) Setup(r *gin.Engine) {
	g := r.Group("fee/inward/v1/")
	{
		_user := handler.InwardHandler{}
		g.POST("iwrm", util.JwtVerify, _user.InwardFee)
	}
}
