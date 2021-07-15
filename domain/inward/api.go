package inward

import (
	"consolidated/config"
	"consolidated/domain/inward/handler"
	"consolidated/domain/inward/repository"
	"consolidated/domain/inward/service"
	"consolidated/util"

	"github.com/gin-gonic/gin"
)

type InwardFeeAPI struct{}

func (u *InwardFeeAPI) Setup(r *gin.Engine) {
	g := r.Group("fee/inward/v1/")
	{
		repository := repository.NewRepository(config.DB)
		service := service.NewService(repository)
		handler := handler.NewHandler(service)

		g.POST("iwrm", util.JwtVerify, handler.Iwrm)
	}
}
