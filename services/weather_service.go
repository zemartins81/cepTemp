package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"cepTemp/models"
)

// WeatherAPIService é o serviço para consulta de clima
type WeatherAPIService struct {
	BaseURL string
	APIKey  string
}

// NewWeatherAPIService cria uma nova instância do serviço WeatherAPI
func NewWeatherAPIService(apiKey string) *WeatherAPIService {
	return &WeatherAPIService{
		BaseURL: "https://api.weatherapi.com/v1",
		APIKey:  apiKey,
	}
}

// GetWeather busca o clima para uma localidade
func (s *WeatherAPIService) GetWeather(city string) (*models.WeatherResponse, error) {
	if city == "" {
		return nil, errors.New("city cannot be empty")
	}

	url := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no", s.BaseURL, s.APIKey, city)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherAPIResponse models.WeatherAPIResponse
	if err := json.Unmarshal(body, &weatherAPIResponse); err != nil {
		return nil, err
	}

	// Converte as temperaturas
	tempC := weatherAPIResponse.Current.TempC
	tempF := tempC*1.8 + 32
	tempK := tempC + 273.15

	return &models.WeatherResponse{
		TempC: tempC,
		TempF: tempF,
		TempK: tempK,
	}, nil
}
