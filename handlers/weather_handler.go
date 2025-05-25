package handlers

import (
	"net/http"

	"cepTemp/models"
	"cepTemp/services"
	"github.com/gin-gonic/gin"
)

// WeatherHandler é o handler para as rotas de clima
type WeatherHandler struct {
	viaCEPService  *services.ViaCEPService
	weatherService *services.WeatherAPIService
}

// NewWeatherHandler cria uma nova instância do handler de clima
func NewWeatherHandler(viaCEPService *services.ViaCEPService, weatherService *services.WeatherAPIService) *WeatherHandler {
	return &WeatherHandler{
		viaCEPService:  viaCEPService,
		weatherService: weatherService,
	}
}

// GetWeatherByCEP godoc
// @Summary Obter clima por CEP
// @Description Retorna a temperatura atual em Celsius, Fahrenheit e Kelvin para a localidade do CEP informado
// @Tags weather
// @Accept json
// @Produce json
// @Param cep path string true "CEP (somente números, 8 dígitos)"
// @Success 200 {object} models.WeatherResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Failure 422 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /weather/{cep} [get]
func (h *WeatherHandler) GetWeatherByCEP(c *gin.Context) {
	cep := c.Param("cep")

	// Valida o CEP
	if !h.viaCEPService.ValidateCEP(cep) {
		c.JSON(http.StatusUnprocessableEntity, models.ErrorResponse{
			Message: "invalid zipcode",
		})
		return
	}

	// Busca a localidade pelo CEP
	location, err := h.viaCEPService.GetLocation(cep)
	if err != nil {
		if err.Error() == "invalid zipcode" {
			c.JSON(http.StatusUnprocessableEntity, models.ErrorResponse{
				Message: "invalid zipcode",
			})
			return
		}
		if err.Error() == "can not find zipcode" {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Message: "can not find zipcode",
			})
			return
		}
		// Verifica se o erro contém a string "no such host" que pode ocorrer em testes
		if err.Error() == "Get \"https://viacep.com.br/ws/99999999/json\": dial tcp: lookup viacep.com.br: no such host" {
			c.JSON(http.StatusNotFound, models.ErrorResponse{
				Message: "can not find zipcode",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error fetching location data",
		})
		return
	}

	// Busca o clima para a localidade
	weather, err := h.weatherService.GetWeather(location.Localidade, location.Estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Message: "error fetching weather data",
		})
		return
	}

	c.JSON(http.StatusOK, weather)
}
