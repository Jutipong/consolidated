package router

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	SetupAuthenAPI(router)
	SetupUserAPI(router)

	return router
}
