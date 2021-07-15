package main

import (
	free "consolidated/domain/fee"
	"consolidated/domain/inward"
	"consolidated/domain/outward"
	"consolidated/feature/auth"

	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func api(r *gin.Engine) {

	//## authentication
	_authApi := auth.AuthAPI{}
	_authApi.Setup(r)

	//## InwardFee
	_inwardFee := inward.InwardAPI{}
	_inwardFee.Setup(r)

	//## OutwardFee
	_outwardFee := outward.OutwardAPI{}
	_outwardFee.Setup(r)

	//## Fee
	_fee := free.FeeAPI{}
	_fee.Setup(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
