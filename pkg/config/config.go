package config

import (
	"github.com/spf13/viper"
	"github.com/sirupsen/logrus"
	
)

type Config struct { 
	Enviroment  EnvironmentConfig
    Database 	DatabaseConfig
    Server   	ServerConfig
    Version  	VersionConfig
}

type EnvironmentConfig struct {
	Env string `mapstructure:"env"`
}

type DatabaseConfig struct {
    Host     string `mapstructure:"host"`
    Port     string `mapstructure:"port"`
    Username string `mapstructure:"username"`
    Password string `mapstructure:"password"`
    DBName   string `mapstructure:"dbname"`
    MaxConnections int `mapstructure:"max_connections"`
}

type ServerConfig struct {
    Port string `mapstructure:"port"`
}

type VersionConfig struct {
    Release string `mapstructure:"release"`
}

func SetupEnv() (Config,error) {
	viper.SetConfigName("dev")
    viper.AddConfigPath("./pkg/config")
	if err := viper.ReadInConfig(); err != nil{
		logrus.Errorf("error reading config file: %v", err)
		return Config{}, err
	}

	var config Config
	err := viper.Unmarshal(&config)
	if err != nil{
		logrus.Errorf("error unmarshilling config: %v", err)
		return Config{}, err
	}
	
	return config, nil
}
