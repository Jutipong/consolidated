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
		g.GET("", helper.JwtVerify, controller.FindAll)
		g.GET("/:id", helper.JwtVerify, controller.FindAll)
		g.POST("", controller.AddNewCustomer)
		g.PUT("", controller.PutOneCustomer)
		g.DELETE("", controller.DeleteCustomer)

	}
}
