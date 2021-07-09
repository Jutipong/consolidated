package inwardfree

import (
	"consolidated/features/inwardfee/controller"
	"consolidated/utils"

	"github.com/gin-gonic/gin"
)

type InwardFeeAPI struct{}

func (u *InwardFeeAPI) Setup(r *gin.Engine) {
	g := r.Group("/api/v1/inwardfee")
	{
		_user := controller.UserController{}
		g.POST("", utils.JwtVerify, _user.InwardFee)
	}
}
