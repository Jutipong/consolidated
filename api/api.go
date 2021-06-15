package api

import (
	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	router := gin.Default()

	SetupAuthenAPI(router)
	// SetupProductAPI(router)
	// SetupTransactionAPI(router)

	return router
}
