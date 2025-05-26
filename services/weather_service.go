package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
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
		BaseURL: "http://api.weatherapi.com/v1",
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

	// Use url.QueryEscape para encoding adequado
	localidadeEncoded := url.QueryEscape(localidade)
	urlStr := fmt.Sprintf("%s/current.json?key=%s&q=%s&aqi=no", s.BaseURL, s.APIKey, localidadeEncoded)

	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, fmt.Errorf("erro na requisição HTTP: %w", err)
	}
	defer resp.Body.Close()

	// Verificar status HTTP
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API retornou status %d: %s", resp.StatusCode, string(body))
	}

	// Ler o body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler response body: %w", err)
	}

	// Verificar se o body não está vazio
	if len(body) == 0 {
		return nil, errors.New("response body está vazio")
	}

	var weatherAPIResponse models.WeatherAPIResponse
	if err := json.Unmarshal(body, &weatherAPIResponse); err != nil {
		fmt.Printf("Erro no JSON Unmarshal: %v\n", err)
		fmt.Printf("Body recebido: %s\n", string(body))
		return nil, fmt.Errorf("erro ao fazer parse do JSON: %w", err)
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
