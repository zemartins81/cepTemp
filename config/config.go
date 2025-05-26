package config

import (
	"os"

	"github.com/joho/godotenv"
)

// Config representa a configuração da aplicação
type Config struct {
	WeatherAPIKey string
	Port          string
}

// NewConfig cria uma nova instância de configuração
func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		panic("Erro ao carregar o arquivo .env")
	}
	key := os.Getenv("WEATHER_API_KEY")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Define uma porta padrão se não estiver configurada
	}
	return &Config{
		WeatherAPIKey: key,
		Port:          port,
	}
}

// GetWeatherAPIKey retorna a chave da API de clima
func (c *Config) GetWeatherAPIKey() string {
	return c.WeatherAPIKey
}

// GetPort retorna a porta do servidor
func (c *Config) GetPort() string {
	return c.Port
}
