package outwardfree

import (
	"consolidated/feature/outwardfee/handler"
	"consolidated/util"

	"github.com/gin-gonic/gin"
)

type OutwardFeeAPI struct{}

func (u *OutwardFeeAPI) Setup(r *gin.Engine) {
	g := r.Group("/fee/outward/v1/")
	{
		_user := handler.OutwardHandler{}
		g.POST("branch", util.JwtVerify, _user.Branch)
	}
}
