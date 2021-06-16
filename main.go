package main

import (
	"consolidated/config"
	"consolidated/router"
)

func init() {
	//## Get config
	if err := config.LoadConfig("./config"); err != nil {
		// log.Println(fmt.Errorf("invalid application configuration: %s", err))
		panic("fail get config: config.yaml")
	} else {
		//## Initial
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
