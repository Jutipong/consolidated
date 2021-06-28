package authentication

import (
	"consolidated/features/authentication/controller"

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
