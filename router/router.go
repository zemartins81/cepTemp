package router

import (
	"cepTemp/config"
	"cepTemp/handlers"
	"cepTemp/middlewares"
	"cepTemp/services"
	"github.com/gin-gonic/gin"
)

// SetupRouter configura as rotas da API
func SetupRouter(cfg *config.Config) *gin.Engine {
	// Inicializa o router
	router := gin.Default()

	// Adiciona middleware de tratamento de erros
	router.Use(middlewares.ErrorHandler())

	// Inicializa os serviços
	viaCEPService := services.NewViaCEPService()
	weatherService := services.NewWeatherAPIService(cfg.GetWeatherAPIKey())

	// Inicializa os handlers
	weatherHandler := handlers.NewWeatherHandler(viaCEPService, weatherService)

	// Configura as rotas
	router.GET("/weather/:cep", weatherHandler.GetWeatherByCEP)

	// Rota de health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return router
}
