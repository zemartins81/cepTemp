package config

import (
	"github.com/joho/godotenv"
	"os"
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
	return &Config{
		WeatherAPIKey: key, // Chave fictícia para desenvolvimento
		Port:          "8080",
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
