package config

import (
	"fmt"

	"consolidated/model"

	"github.com/spf13/viper"
)

var Config appConfig

type appConfig struct {
	Server         model.ServerConfiguration
	Databaseconfig model.DatabaseConfiguration
	LoggerFile     model.LoggerFile
}

// LoadConfig loads config from files
func SetupConfg(configPaths ...string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.SetEnvPrefix("")
	v.AutomaticEnv()
	for _, path := range configPaths {
		v.AddConfigPath(path)
	}
	if err := v.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read the configuration file: %s", err)
	}
	return v.Unmarshal(&Config)
}