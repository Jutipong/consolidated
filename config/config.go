package config

import (
	"github.com/spf13/viper"
)

//## Server
type server struct {
	Env         string
	RunMode     string
	Port        string
	SecretKey   string
	TokenExpire int
}

//## Database
type database struct {
	Server       string
	Port         int
	User         string
	Pass         string
	DatabaseName string
}

//## Logger File
type loggerFile struct {
	RootPath string
}

var config appConfig

type appConfig struct {
	Server     server
	Database   database
	LoggerFile loggerFile
}

func InitialConfig() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("")
	v.AutomaticEnv()
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	v.Unmarshal(&config)
}

func Server() server {
	return config.Server
}

func Database() database {
	return config.Database
}

func LogFile() loggerFile {
	return config.LoggerFile
}
