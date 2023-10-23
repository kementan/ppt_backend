package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPServerAddress      string        `mapstructure:"HTTP_SERVER_ADDRESS"`
	WebSocketServerAddress string        `mapstructure:"WEBSOCKET_SERVER_ADDRESS"`
	APPKey                 string        `mapstructure:"APP_KEY"`
	SecretKey              string        `mapstructure:"SECRET_KEY"`
	OWM_API                string        `mapstructure:"OWM_API"`
	AllowOrigin            string        `mapstructure:"ALLOW_ORIGIN"`
	TokenAccessDuration    time.Duration `mapstructure:"TOKEN_ACCESS_DURATION"`
	TokenRefreshDuration   time.Duration `mapstructure:"TOKEN_REFRESH_DURATION"`
	GSiteKey               string        `mapstructure:"GRECAPTCHA_SITE_KEY"`
	GSecretKey             string        `mapstructure:"GRECAPTCHA_SECRET_KEY"`
	PSQLDBDriver           string        `mapstructure:"PSQL_DB_DRIVER"`
	PSQLDBSource           string        `mapstructure:"PSQL_DB_SOURCE"`
	MYSQLDBUsername        string        `mapstructure:"MYSQL_DB_USERNAME"`
	MYSQLDBPassword        string        `mapstructure:"MYSQL_DB_PASSWORD"`
	MYSQLDBHost            string        `mapstructure:"MYSQL_DB_HOST"`
	MYSQLDBPort            int           `mapstructure:"MYSQL_DB_PORT"`
	MYSQLDBName            string        `mapstructure:"MYSQL_DB_NAME"`
	RedisDBAddress         string        `mapstructure:"REDIS_DB_ADDRESS"`
	RedisDBPassword        string        `mapstructure:"REDIS_DB_PASSWORD"`
	RedisDBIndex           int           `mapstructure:"REDIS_DB_INDEX"`
	ElasticDBAddress       string        `mapstructure:"ELASTIC_DB_ADDRESS"`
	ElasticDBUser          string        `mapstructure:"ELASTIC_DB_USER"`
	ElasticDBPassword      string        `mapstructure:"ELASTIC_DB_PASSWORD"`
	ElasticIVKey           string        `mapstructure:"ELASTIC_IV_KEY"`
	KafkaBroker            string        `mapstructure:"KAFKA_BROKER"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app.dev")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
