package main

import (
	"consolidated/base"
	"consolidated/config"
	api "consolidated/feature"
	"consolidated/util"

	"github.com/gin-gonic/gin"
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

func main() {
	r := gin.Default()
	r.Use(util.LoggerFile())
	r.Use(gin.Recovery())
	api.Setup(r)
	r.Run(":" + config.Server().Port)
}
