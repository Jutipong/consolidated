package router

import (
	"consolidated/controller"

	"github.com/gin-gonic/gin"
)

func SetupUserAPI(r *gin.Engine) {
	// g := r.Group("/api", middleware.JwtVerify())
	g := r.Group("/api/user")
	{
		g.GET("", controller.FindAll)
		g.GET("/:id", controller.FindAll)
		g.POST("", controller.AddNewCustomer)
		g.PUT("", controller.PutOneCustomer)
		g.DELETE("", controller.DeleteCustomer)

	}
}
