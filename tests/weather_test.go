package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"cepTemp/api"
	"cepTemp/config"
	"cepTemp/models"
	"github.com/gin-gonic/gin"
)

func TestGetWeatherByCEP(t *testing.T) {
	// Configura o modo de teste do Gin
	gin.SetMode(gin.TestMode)

	// Cria uma configuração de teste
	cfg := config.NewConfig()

	// Configura o router
	router := api.SetupRouter(cfg)

	// Testes para diferentes cenários
	tests := []struct {
		name           string
		cep            string
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name:           "CEP inválido (formato incorreto)",
			cep:            "123",
			expectedStatus: http.StatusUnprocessableEntity,
			expectedBody:   models.ErrorResponse{Message: "invalid zipcode"},
		},
		{
			name:           "CEP não encontrado",
			cep:            "99999999",
			expectedStatus: http.StatusNotFound,
			expectedBody:   models.ErrorResponse{Message: "can not find zipcode"},
		},
		// Nota: O teste de sucesso não é possível sem uma chave de API real
	}

	// Executa os testes
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Cria uma requisição de teste
			req, _ := http.NewRequest("GET", "/weather/"+tt.cep, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			// Verifica o status code
			if w.Code != tt.expectedStatus {
				t.Errorf("Status code esperado %d, obtido %d", tt.expectedStatus, w.Code)
			}

			// Verifica o corpo da resposta para erros
			if tt.expectedStatus != http.StatusOK {
				var response models.ErrorResponse
				err := json.Unmarshal(w.Body.Bytes(), &response)
				if err != nil {
					t.Errorf("Erro ao decodificar resposta: %v", err)
				}

				expectedBody, _ := tt.expectedBody.(models.ErrorResponse)
				if response.Message != expectedBody.Message {
					t.Errorf("Mensagem esperada %s, obtida %s", expectedBody.Message, response.Message)
				}
			}
		})
	}
}
