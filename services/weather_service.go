package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
func (s *WeatherAPIService) GetWeather(localidade, estado string) (*models.WeatherResponse, error) {
	if localidade == "" {
		return nil, errors.New("localidade cannot be empty")
	}
	localidadeAjustada := strings.ReplaceAll(localidade, " ", "_")
	estadoAjustado := strings.ReplaceAll(estado, " ", "_")
	localidade = localidadeAjustada + "_" + estadoAjustado

	url := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no", s.BaseURL, s.APIKey, localidade)
	fmt.Println(url)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	fmt.Println("Body:")
	fmt.Println(body)

	var weatherAPIResponse models.WeatherAPIResponse
	if err := json.Unmarshal(body, &weatherAPIResponse); err != nil {
		fmt.Println(err.Error())
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
