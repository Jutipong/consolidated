package router

import (
	"consolidated/controller"
	"consolidated/service"

	"github.com/gin-gonic/gin"
)

func SetupUserAPI(r *gin.Engine) {
	g := r.Group("/api")
	// g := r.Group("/api/user")
	{
		g.GET("", service.JwtVerify, controller.FindAll)
		g.GET("/:id", service.JwtVerify, controller.FindAll)
		g.POST("", controller.AddNewCustomer)
		g.PUT("", controller.PutOneCustomer)
		g.DELETE("", controller.DeleteCustomer)

	}
}
