package main

import (
	"consolidated/config"
	"consolidated/helper"
	"consolidated/router"
	"fmt"
	"time"
)

func init() {
	//## Config
	if err := config.SetupConfg("./config"); err != nil {
		panic("fail get config: config.yaml")
	} else {
		//## Database
		// config.SetupDatabase()
		//## Logger
		helper.SetupLogger()
	}
}

func main() {

	//## Set up router
	router := router.Setup()

	i := 0
	for {
		i = i + 1

		helper.LogInfo(fmt.Sprint("Count: ",i))
		time.Sleep(time.Duration(1) * time.Second)
		// log.Info(fmt.Sprint("Count: ", i))
		// time.Sleep(time.Duration(1) * time.Second)
		// log.Error(fmt.Sprint("Count: ", i+2))
		// log.Warning(fmt.Sprint("Count: ", i+5))
		// time.Sleep(time.Duration(1) * time.Second)
	}
	//## Set port
	port := "8080"
	router.Run(":" + port)
}
