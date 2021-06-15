package router

import (
	"consolidated/controller"

	"github.com/gin-gonic/gin"
)

func SetupAuthenAPI(r *gin.Engine) {
	g := r.Group("/api")
	{
		g.POST("/login", controller.Login)
	}
}
