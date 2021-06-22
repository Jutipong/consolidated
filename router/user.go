package router

import (
	"consolidated/controller"
	"consolidated/helper"

	"github.com/gin-gonic/gin"
)

func SetupUserAPI(r *gin.Engine) {
	g := r.Group("/api/user")
	{
		user := controller.UserController{}
		g.GET("", helper.JwtVerify, user.FindAll)
		g.GET("/:id", helper.JwtVerify, user.FindID)
		g.POST("", helper.JwtVerify, user.AddNewCustomer)
		g.PUT("", helper.JwtVerify, user.PutOneCustomer)
		g.DELETE("", helper.JwtVerify, user.DeleteCustomer)
	}
}
