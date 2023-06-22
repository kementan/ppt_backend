package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	APPKey               string        `mapstructure:"APP_KEY"`
	HTTPServerAddress    string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	AccessTokenDuration  time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
	RefreshTokenDuration time.Duration `mapstructure:"REFRESH_TOKEN_DURATION"`
	PSQLDBDriver         string        `mapstructure:"PSQL_DB_DRIVER"`
	PSQLDBSource         string        `mapstructure:"PSQL_DB_SOURCE"`
	RedisDBAddress       string        `mapstructure:"REDIS_DB_ADDRESS"`
	RedisDBPassword      string        `mapstructure:"REDIS_DB_PASSWORD"`
	RedisDBIndex         int           `mapstructure:"REDIS_DB_INDEX"`
	ElasticDBAddress     string        `mapstructure:"ELASTIC_DB_ADDRESS"`
	ElasticDBUser        string        `mapstructure:"ELASTIC_DB_USER"`
	ElasticDBPassword    string        `mapstructure:"ELASTIC_DB_PASSWORD"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
