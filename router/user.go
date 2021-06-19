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
	g := r.Group("/api")
	{
		g.GET("", helper.JwtVerify, helper.LoggerRequest, controller.FindAll)
		g.GET("/:id", helper.JwtVerify, helper.LoggerRequest, controller.FindAll)
		g.POST("", helper.JwtVerify, helper.LoggerRequest, controller.AddNewCustomer)
		g.PUT("", helper.JwtVerify, helper.LoggerRequest, controller.PutOneCustomer)
		g.DELETE("", helper.JwtVerify, helper.LoggerRequest, controller.DeleteCustomer)
	}
}
