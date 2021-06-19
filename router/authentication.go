package router

import (
	"consolidated/controller"

	"github.com/gin-gonic/gin"
)

//## helper.JwtVerify	  =>  ใช้สำหรับเช็ค token
//## helper.LoggerRequest =>  ใช้สำหรับเขียน logger request
//## controller.xxxxx	  =>  เรียก function Logic....
func SetupAuthenAPI(r *gin.Engine) {
	g := r.Group("/api")
	{
		g.POST("/login", controller.Login)
	}
}
