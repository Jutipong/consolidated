package outwardfree

import (
	"consolidated/features/outwardfee/controller"
	"consolidated/utils"

	"github.com/gin-gonic/gin"
)

type OutwardFeeAPI struct{}

func (u *OutwardFeeAPI) Setup(r *gin.Engine) {
	g := r.Group("/fee/outward/v1/")
	{
		_user := controller.OutwardController{}
		g.POST("branch", utils.JwtVerify, _user.Branch)
	}
}
