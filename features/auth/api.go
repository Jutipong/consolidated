package auth

import (
	"consolidated/features/auth/controller"

	"github.com/gin-gonic/gin"
)

type AuthAPI struct{}

func (a *AuthAPI) Setup(r *gin.Engine) {
	g := r.Group("/api")
	{
		auth := controller.AuthenticatedController{}
		g.POST("/login", auth.Login)
	}
}
