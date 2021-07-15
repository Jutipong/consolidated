package config

import (
	"fmt"

	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var _db *gorm.DB

func InitialDB() {
	configDb := getConfigDb()
	db, err := gorm.Open(sqlserver.Open(configDb+"&parseTime=True"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	_db = db
}

func getConfigDb() string {
	if config.Database.Port > 0 {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s:%d?database=%s",
			config.Database.User,
			config.Database.Pass,
			config.Database.Server,
			config.Database.Port,
			config.Database.DatabaseName,
		)
	} else {
		return fmt.Sprintf(
			"sqlserver://%s:%s@%s?database=%s",
			config.Database.User,
			config.Database.Pass,
			config.Database.Server,
			config.Database.DatabaseName,
		)
	}
}

func Db() *gorm.DB {
	return _db
}
