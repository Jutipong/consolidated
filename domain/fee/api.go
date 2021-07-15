package fee

import (
	"consolidated/config"
	"consolidated/domain/fee/handler"
	"consolidated/domain/fee/repository"
	"consolidated/domain/fee/service"
	"consolidated/util"

	"github.com/gin-gonic/gin"
)

type FeeAPI struct{}

func (u *FeeAPI) Setup(r *gin.Engine) {
	repository := repository.NewRepository(config.DB)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)
	g := r.Group("fee/v1/")
	{
		g.POST("promotioncode", util.JwtVerify, handler.PromotionCode)
	}
}
