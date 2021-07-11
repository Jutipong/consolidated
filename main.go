package main

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/middleware"
	"consolidated/router"
	"consolidated/utils"

	"github.com/gin-gonic/gin"
)

//GIN_MODE=debug
//GIN_MODE=release
func init() {
	config.InitialConfig()
	config.InitialDB()
	utils.SetupLogger()
	base.InitMasterRule()
}

func main() {
	r := gin.Default()
	r.Use(middleware.GinBodyLogMiddleware())
	r.Use(gin.Recovery())
	router.Setup(r)
	r.Run(":" + config.Server().Port)
}
