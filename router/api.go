package router

import (
	"consolidated/middleware"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {

	router := gin.Default()
	router.Use(middleware.GinBodyLogMiddleware())
	router.Use(gin.Recovery())

	SetupAuthenAPI(router)
	SetupUserAPI(router)

	return router
}
