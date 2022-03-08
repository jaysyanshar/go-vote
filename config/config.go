package config

import (
	"github.com/labstack/gommon/log"
	"github.com/spf13/viper"
)

type Config struct {
	ServicePort int `mapstructure:"SERVICE_PORT"`

	AccessSecret          string `mapstructure:"ACCESS_SECRET"`
	TokenExpirationMinute int    `mapstructure:"TOKEN_EXPIRATION_MINUTE"`

	DbDriver                string `mapstructure:"DB_DRIVER"`
	DbSource                string `mapstructure:"DB_SOURCE"`
	DbConnMaxLifetimeMinute int    `mapstructure:"DB_CONN_MAX_LIFETIME_MINUTE"`
	DbMaxOpenIdleConn       int    `mapstructure:"DB_MAX_OPEN_IDLE_CONN"`

	RedisAddress  string `mapstructure:"REDIS_ADDRESS"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDb       int    `mapstructure:"REDIS_DB"`
}

const (
	configPath = "./Config/"
	configName = "Config"
	configType = "env"
)

var config *Config

func Get() *Config {
	return config
}

func Init() (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(configName)
	viper.SetConfigType(configType)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Errorf("failed to read config: %v", err)
		return nil, err
	}

	err = viper.Unmarshal(&config)
	log.Infof("config successfully read")
	return Get(), nil
}
