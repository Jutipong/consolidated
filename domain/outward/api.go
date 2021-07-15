package outward

import (
	"consolidated/feature/outwardfee/handler"
	"consolidated/util"

	"github.com/gin-gonic/gin"
)

type OutwardAPI struct{}

func (u *OutwardAPI) Setup(r *gin.Engine) {
	g := r.Group("/fee/outward/v1/")
	{
		_user := handler.OutwardHandler{}
		g.POST("branch", util.JwtVerify, _user.Branch)
	}
}
