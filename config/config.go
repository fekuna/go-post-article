package config

import (
	"errors"
	"log"
	"time"

	"github.com/spf13/viper"
)

// App config struct
type Config struct {
	Server ServerConfig
	Logger LoggerConfig
	Mysql  MysqlConfig
}

type ServerConfig struct {
	AppVersion        string
	Port              string
	Mode              string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	CtxDefaultTimeout time.Duration
	SSL               string
	Debug             bool
}

// Logger config
type LoggerConfig struct {
	Development       bool
	DisableCaller     bool
	DisableStackTrace bool
	Encoding          string
	Level             string
}

// Postgresql config
type MysqlConfig struct {
	MysqlHost     string
	MysqlPort     string
	MysqlUsername string
	MysqlPassword string
	MysqlDbName   string
	MysqlSslMode  bool
	Driver        string
}

// Load config file from given path
func LoadConfig(fileName string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(fileName)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}

		return nil, err
	}

	return v, nil
}

// Parse config file
func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	err := v.Unmarshal(&c)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)
		return nil, err
	}

	return &c, nil
}
