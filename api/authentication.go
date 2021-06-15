package api

import (
	"consolidated/controller"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(r *gin.Engine) {
	authenAPI := r.Group("/api")
	{
		authenAPI.POST("/login", controller.Login)

		// func(ctx *gin.Context) {
		// 	controller.Login(ctx)
		// })

	}
}
