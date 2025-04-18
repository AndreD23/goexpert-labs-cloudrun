package configs

import "github.com/spf13/viper"

var config *Config

type Config struct {
	WeatherAPI string `mapstructure:"WEATHER_API"`
}

func NewConfig() *Config {
	return config
}

func init() {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	config = &Config{
		WeatherAPI: viper.GetString("WEATHER_API"),
	}
}
