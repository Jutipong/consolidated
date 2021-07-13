package auth

import (
	"consolidated/feature/auth/handler"

	"github.com/gin-gonic/gin"
)

type AuthAPI struct{}

func (a *AuthAPI) Setup(r *gin.Engine) {
	g := r.Group("/")
	{
		g.GET("login", handler.Login)
	}
}
