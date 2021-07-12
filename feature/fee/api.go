package inwardfree

import (
	"consolidated/feature/fee/handler"
	"consolidated/util"

	"github.com/gin-gonic/gin"
)

type FeeAPI struct{}

func (u *FeeAPI) Setup(r *gin.Engine) {
	g := r.Group("fee/v1/")
	{
		_user := handler.FeeHandler{}
		g.POST("promotioncode", util.JwtVerify, _user.Fee)
	}
}
