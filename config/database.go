package config

import (
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := "sqlserver://sa:p@ssw0rd@localhost?database=Consolidated"
	db, err := gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		//fmt.Println("statuse: ", err)
		panic("failed to connect database: ")
	}
	// defer DB.Close()
	DB = db
}
