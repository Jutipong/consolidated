package main

import (
	"consolidated/base"
	"consolidated/config"
	"consolidated/middleware"
	"consolidated/router"
	"consolidated/utils"

	"github.com/gin-gonic/gin"
)

//GIN_MODE=debug
//GIN_MODE=release
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

func main() {
	r := gin.Default()
	r.Use(middleware.GinBodyLogMiddleware())
	r.Use(gin.Recovery())
	router.Setup(r)
	r.Run(":" + config.Config.Server.Port)
}
