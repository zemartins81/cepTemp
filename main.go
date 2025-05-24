package main

import (
	"cepTemp/router"
	"log"
	"os"

	"cepTemp/config"
)

// @title Weather CEP API
// @version 1.0
// @description API para consulta de clima por CEP
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	// Carrega configurações
	cfg := config.NewConfig()

	// Verifica se existe uma chave de API definida como variável de ambiente
	if apiKey := os.Getenv("WEATHER_API_KEY"); apiKey != "" {
		cfg.WeatherAPIKey = apiKey
	}

	// Verifica se existe uma porta definida como variável de ambiente
	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = port
	}

	// Configura o router
	newRouter := router.SetupRouter(cfg)

	// Inicia o servidor
	log.Printf("Servidor iniciado na porta %s", cfg.Port)
	if err := newRouter.Run(":" + cfg.Port); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
