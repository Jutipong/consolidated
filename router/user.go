package router

import (
	"consolidated/controller"
	"consolidated/helper"

	"github.com/gin-gonic/gin"
)

func SetupUserAPI(r *gin.Engine) {
	g := r.Group("/api")
	// g := r.Group("/api/user")
	{
		g.GET("", helper.JwtVerify, helper.LoggerRequest, controller.FindAll)
		g.GET("/:id", helper.JwtVerify, helper.LoggerRequest, controller.FindAll)
		g.POST("", controller.AddNewCustomer, helper.LoggerRequest)
		g.PUT("", controller.PutOneCustomer, helper.LoggerRequest)
		g.DELETE("", controller.DeleteCustomer, helper.LoggerRequest)

	}
}
