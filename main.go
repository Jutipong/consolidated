package main

import (
	"consolidated/config"
	"consolidated/router"
)

func main() {

	//## Initial
	config.ConnectDB()
	// config.DB.Table("Customer").AutoMigrate(&entity.Customer{})

	//## Set up router
	router := router.Setup()

	//## Set port
	port := "8080"
	router.Run(":" + port)
}
