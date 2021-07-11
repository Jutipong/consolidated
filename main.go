package main

import (
	"consolidated/base"
	"consolidated/config"
	api "consolidated/features"
	"consolidated/middleware"
	"consolidated/utils"

	"github.com/gin-gonic/gin"
)

//GIN_MODE=debug
//GIN_MODE=release
func init() {
	config.InitialConfig()
	config.InitTimeZone()
	config.InitialDB()
	utils.SetupLogger()
	base.InitMasterRule()
}

func main() {
	r := gin.Default()
	r.Use(middleware.LoggerFile())
	r.Use(gin.Recovery())
	api.Setup(r)
	r.Run(":" + config.Server().Port)
}
