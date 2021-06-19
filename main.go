package main

import (
	"consolidated/config"
	"consolidated/helper"
	"consolidated/router"
)

func init() {
	//## Config
	if err := config.SetupConfg("./config"); err != nil {
		panic("fail get config: config.yaml")
	} else {
		//## Database
		config.SetupDatabase()
		//## Logger
		helper.SetupLogger()
	}
}

func main() {

	//## Set up router
	router := router.Setup()

	//## Set port
	router.Run(":" + config.Config.Server.Port)
}
