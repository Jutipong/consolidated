package main

import (
	"consolidated/config"
	"consolidated/helper"
	"consolidated/router"
)

func main() {

	//## Set up router
	router := router.Setup()

	//## Set port
	router.Run(":" + config.Config.Server.Port)
}

func init() {
	//## Config
	if err := config.SetupConfg("./config"); err != nil {
		panic("fail get config: config.yaml")
	} else {
		//## Logger
		if err := helper.SetupLogger(); err != "" {
			helper.LogError(err)
			panic(err)
		}
		//## Database
		if err := config.SetupDatabase(); err != "" {
			helper.LogError(err)
			panic(err)
		}
	}
}
