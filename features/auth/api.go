package auth

import (
	"consolidated/features/auth/handler"

	"github.com/gin-gonic/gin"
)

type AuthAPI struct{}

func (a *AuthAPI) Setup(r *gin.Engine) {
	g := r.Group("/api")
	{
		g.POST("/login", handler.Login)
	}
}
