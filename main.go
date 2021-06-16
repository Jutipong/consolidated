package main

import (
	"consolidated/config"
	"consolidated/router"
)

func init() {
	//## Config
	if err := config.SetupConfg("./config"); err != nil {
		panic("fail get config: config.yaml")
	} else {
		//## Database
		config.SetupDatabase()
	}
}

func main() {

	//## Set up router
	router := router.Setup()

	//## Set port
	port := "8080"
	router.Run(":" + port)
}
