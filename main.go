package main

import (
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
	//# Setup validate
	validate.Config(func(opt *validate.GlobalOption) {
		opt.StopOnError = false
		// opt.SkipOnEmpty = false
	})
}

func main() {
	//## 4.Router
	router := router.Setup()
	//## 5.Start Server
	router.Run(":" + config.Config.Server.Port)
}
