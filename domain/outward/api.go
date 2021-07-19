package outward

import (
	"consolidated/config"
	"consolidated/domain/outward/handler"
	"consolidated/domain/outward/repository"
	"consolidated/domain/outward/service"
	"consolidated/util"

	"github.com/gin-gonic/gin"
)

type OutwardAPI struct{}

func (u *OutwardAPI) Setup(r *gin.Engine) {
	g := r.Group("/fee/outward/v1/")
	{
		repository := repository.NewRepository(config.Db())
		service := service.NewService(repository)
		handler := handler.NewHandler(service)
		g.POST("branch", util.JwtVerify, handler.Branch)
	}
}
