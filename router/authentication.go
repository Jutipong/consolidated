package router

import (
	"consolidated/controller"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(r *gin.Engine) {
	g := r.Group("/api")
	{
		auth := controller.AuthenticatedController{}
		g.POST("/login", auth.Login)
	}
}
