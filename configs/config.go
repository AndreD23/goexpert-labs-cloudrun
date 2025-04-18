package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

var config *Config

type Config struct {
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
}

func NewConfig() *Config {
	return config
}

func init() {
	var err error
	config, err = loadConfig()
	if err != nil {
		panic(fmt.Sprintf("Erro ao carregar configurações: %v", err))
	}
}

func loadConfig() (*Config, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetConfigFile(".env")

	// Permite override por variáveis de ambiente
	viper.AutomaticEnv()

	// Tenta ler o arquivo .env, mas não falha se não encontrar
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			// Só retorna erro se não for um erro de arquivo não encontrado
			return nil, fmt.Errorf("erro ao ler arquivo de configuração: %w", err)
		}
		fmt.Println("Arquivo .env não encontrado. Usando variáveis de ambiente.")
	}

	// Define valores padrão
	viper.SetDefault("WEATHER_API_KEY", "") // Você pode definir um valor padrão se desejar

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		return nil, fmt.Errorf("erro ao processar configurações: %w", err)
	}

	// Validação das configurações obrigatórias
	if config.WeatherAPIKey == "" {
		return nil, fmt.Errorf("WEATHER_API_KEY é obrigatória")
	}

	return config, nil
}
