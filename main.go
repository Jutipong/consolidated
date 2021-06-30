package main

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/router"
	"consolidated/utils"

	"github.com/gookit/validate"
)

func init() {
	//## 1.Config
	if err := config.SetupConfg("./config"); err != nil {
		panic("fail get config: config.yaml")
	}
	//## 2.Logger File
	if err := utils.SetupLogger(); err != "" {
		utils.LogError(err)
		panic(err)
	}
	//## 3.Database
	if err := config.SetupDatabase(); err != "" {
		utils.LogError(err)
		panic(err)
	}
	//## 4.Setup validate
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		// opt.SkipOnEmpty = false
	})

	//## 5.Initial Master Validate Rule
	base.InitMasterRule()
}

func main() {
	//## 6.Router
	router := router.Setup()
	//## 7.Start Server
	router.Run(":" + config.Config.Server.Port)
}
