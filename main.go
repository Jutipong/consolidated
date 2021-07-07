package main

import (
	"consolidated/base"
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
		panic(err)
	}
	//## 3.Database
	if err := config.SetupDatabase(); err != "" {
		utils.LogError(err)
		panic(err)
	}

	//## 4.Initial Master Validate Rule
	base.InitMasterRule()
	//## 5. Setup Custom Validation ยังไม่ได้ใช้งาน
	// base.SetupCustomValidate()
}

//GIN_MODE=debug
//GIN_MODE=release
func main() {
	//## 6.Router
	router := router.Setup()
	//## 7.Start Server
	router.Run(":" + config.Config.Server.Port)
}
