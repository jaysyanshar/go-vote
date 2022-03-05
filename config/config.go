package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type Config struct {
	ServicePort             int    `mapstructure:"SERVICE_PORT"`
	DbDriver                string `mapstructure:"DB_DRIVER"`
	DbSource                string `mapstructure:"DB_SOURCE"`
	DbConnMaxLifetimeMinute int    `mapstructure:"DB_CONN_MAX_LIFETIME_MINUTE"`
	DbMaxOpenIdleConn       int    `mapstructure:"DB_MAX_OPEN_IDLE_CONN"`
}

const (
	configPath = "./Config/"
	configName = "Config"
	configType = "env"
)

func LoadConfig() (config *Config, err error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Errorf("failed to read config: %v", err)
		return
	}

	err = viper.Unmarshal(&config)
	log.Infof("config successfully read")
	return
}
