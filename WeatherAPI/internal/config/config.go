package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

//Viper is used for environment loading
type Config struct {
	AppUrl      string
	AppID       string
	Host        string
	Port        string
	ForecastUrl string
}

const (
	Port        = "app.port"
	Host        = "app.host"
	Url         = "api.url"
	AppID       = "api.appId"
	ForecastUrl = "api.forecastUrl"
)

var EnvConfig *Config

func Load() (*Config, error) {
	env := os.Getenv("env")
	var config Config
	//Setting Location for Viper
	if len(env) <= 0 {
		env = "local"
	}
	fmt.Println("Loading Config")
	viper.SetConfigName("config-" + env)
	viper.AddConfigPath("config")
	viper.AddConfigPath("../config")
	viper.AddConfigPath("../../config")
	viper.AddConfigPath("../../../config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("couldn't read config file: %s", err)
	}
	config.AppUrl = viper.GetString(Url)
	config.AppID = viper.GetString(AppID)
	config.Host = viper.GetString(Host)
	config.Port = viper.GetString(Port)
	config.ForecastUrl = viper.GetString(ForecastUrl)
	EnvConfig = &config
	return EnvConfig, nil

}
