package main

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/util"

	"github.com/gin-gonic/gin"
	_ "consolidated/docs"
)

//GIN_MODE=debug
//GIN_MODE=release
func init() {
	config.InitialConfig()
	config.InitTimeZone()
	config.InitialDB()
	config.SetupLogger()
	base.InitMasterRule()
}


// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	r := gin.Default()
	r.Use(util.LoggerFile())
	r.Use(gin.Recovery())
	api(r)
	r.Run(":" + config.Server().Port)
}
