package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() {
	configDb := getConfigDb()
	db, err := gorm.Open(sqlserver.Open(configDb), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("fail to connect to database: %s", configDb))
	}
	DB = db
	// defer DB.Close()
}

func getConfigDb() string {

	if Config.Databaseconfig.Port > 0 {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s:%d?database=%s",
			Config.Databaseconfig.User,
			Config.Databaseconfig.Pass,
			Config.Databaseconfig.Server,
			Config.Databaseconfig.Port,
			Config.Databaseconfig.DatabaseName,
		)
	} else {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=%s",
			Config.Databaseconfig.User,
			Config.Databaseconfig.Pass,
			Config.Databaseconfig.Server,
			Config.Databaseconfig.DatabaseName,
		)
	}
}
