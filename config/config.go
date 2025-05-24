package config

// Config representa a configuração da aplicação
type Config struct {
	WeatherAPIKey string
	Port          string
}

// NewConfig cria uma nova instância de configuração
func NewConfig() *Config {
	return &Config{
		WeatherAPIKey: "DUMMY_API_KEY_12345", // Chave fictícia para desenvolvimento
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
