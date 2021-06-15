package main

import (
	// "init/config"
	// "init/entity"
	// "init/routers"
	// "os"

	"consolidated/api"
	"consolidated/config"
)

func main() {

	//Initial
	// router := gin.Default()

	config.ConnectDB()
	// config.DB.Table("Customer").AutoMigrate(&entity.Customer{})

	//Set up router
	router := api.Setup()

	//Set startup port
	port := "8080"
	router.Run(":" + port)
}
