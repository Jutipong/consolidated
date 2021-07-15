package main

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/util"

	_ "consolidated/docs"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
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

// @title Consolidated API
// @version 1.0
// @description (◍•ᴗ•◍)❤

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in Bearer token
// @name Authorization
func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(util.LoggerFile())
	r.Use(gin.Recovery())
	api(r)
	r.Run(":" + config.Server().Port)
}
