package router

import (
	auth "consolidated/features/authentication"
	"consolidated/features/user"
	"consolidated/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.GinBodyLogMiddleware())
	router.Use(gin.Recovery())

	//## authentication
	authApi := auth.AuthAPI{}
	authApi.Setup(router)

	//## User
	userApi := user.UserAPI{}
	userApi.Setup(router)

	return router
}
