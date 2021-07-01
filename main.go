package main

import (
	"consolidated/config"
	"consolidated/router"
	"consolidated/utils"
)

func init() {
	//## 1.Config
	if err := config.SetupConfig("./config"); err != nil {
		panic("fail get config: config.yaml")
	}
	//## 2.Logger File
	if err := utils.SetupLogger(); err != "" {
		utils.LogError(err)
		panic(err)
	}
	//## 3.Database
	// if err := config.SetupDatabase(); err != "" {
	// 	utils.LogError(err)
	// 	panic(err)
	// }
	//## 4.Setup validate

	//## 5.Initial Master Validate Rule
	config.InitMasterRule()
	//## 6. Setup Custom Validation
	config.SetupCustomValidate()
}

func main() {
	//## 7.Router
	router := router.Setup()
	//## 8.Start Server
	router.Run(":" + config.Config.Server.Port)
}
