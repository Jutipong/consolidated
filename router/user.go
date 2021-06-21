package router

import (
	"consolidated/controller"
	"consolidated/helper"

	"github.com/gin-gonic/gin"
)

//## helper.JwtVerify	  =>  ใช้สำหรับเช็ค token
//## helper.LoggerRequest =>  ใช้สำหรับเขียน logger request
//## controller.xxxxx	  =>  เรียก function Logic....
func SetupUserAPI(r *gin.Engine) {
	g := r.Group("/api/user")
	{
		g.GET("", helper.JwtVerify, controller.FindAll)
		g.GET("/:id", helper.JwtVerify, controller.FindID)
		g.POST("", helper.JwtVerify, controller.AddNewCustomer)
		g.PUT("", helper.JwtVerify, controller.PutOneCustomer)
		g.DELETE("", helper.JwtVerify, controller.DeleteCustomer)
	}
}
