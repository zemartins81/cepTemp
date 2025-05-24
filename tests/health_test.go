package tests

import (
	"cepTemp/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"cepTemp/config"
	"github.com/gin-gonic/gin"
)

func TestHealthEndpoint(t *testing.T) {
	// Configura o modo de teste do Gin
	gin.SetMode(gin.TestMode)

	// Cria uma configuração de teste
	cfg := config.NewConfig()

	// Configura o router
	newRouter := router.SetupRouter(cfg)

	// Cria uma requisição de teste para o endpoint de health check
	req, _ := http.NewRequest("GET", "/health", nil)
	w := httptest.NewRecorder()
	newRouter.ServeHTTP(w, req)

	// Verifica o status code
	if w.Code != http.StatusOK {
		t.Errorf("Status code esperado %d, obtido %d", http.StatusOK, w.Code)
	}

	// Verifica o corpo da resposta
	expected := `{"status":"ok"}`
	if w.Body.String() != expected {
		t.Errorf("Resposta esperada %s, obtida %s", expected, w.Body.String())
	}
}
