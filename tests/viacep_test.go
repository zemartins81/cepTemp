package tests

import (
	"testing"

	"cepTemp/services"
)

func TestViaCEPValidation(t *testing.T) {
	service := services.NewViaCEPService()

	tests := []struct {
		name     string
		cep      string
		expected bool
	}{
		{
			name:     "CEP válido",
			cep:      "01001000",
			expected: true,
		},
		{
			name:     "CEP inválido - muito curto",
			cep:      "123",
			expected: false,
		},
		{
			name:     "CEP inválido - muito longo",
			cep:      "123456789",
			expected: false,
		},
		{
			name:     "CEP inválido - contém letras",
			cep:      "1234567a",
			expected: false,
		},
		{
			name:     "CEP inválido - vazio",
			cep:      "",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := service.ValidateCEP(tt.cep)
			if result != tt.expected {
				t.Errorf("ValidateCEP(%s) = %v, esperado %v", tt.cep, result, tt.expected)
			}
		})
	}
}
