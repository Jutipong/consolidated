package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupDatabase() string {
	configDb := getConfigDb()
	db, err := gorm.Open(sqlserver.Open(configDb), &gorm.Config{})
	if err != nil {
		return fmt.Sprintf("fail to connect to database config: %s error: %v", configDb, err.Error())
	}
	DB = db
	// defer DB.Close()
	return ""
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
