package tests

import (
	"testing"
)

func TestTemperatureConversion(t *testing.T) {
	// Testes para conversão de temperatura
	// Usando valores conhecidos para validar as fórmulas

	// Celsius para Fahrenheit: F = C * 1,8 + 32
	// Celsius para Kelvin: K = C + 273.15

	tests := []struct {
		name      string
		tempC     float64
		expectedF float64
		expectedK float64
	}{
		{
			name:      "Temperatura 0°C",
			tempC:     0,
			expectedF: 32,
			expectedK: 273.15,
		},
		{
			name:      "Temperatura 25°C",
			tempC:     25,
			expectedF: 77,
			expectedK: 298.15,
		},
		{
			name:      "Temperatura 100°C",
			tempC:     100,
			expectedF: 212,
			expectedK: 373.15,
		},
		{
			name:      "Temperatura negativa -10°C",
			tempC:     -10,
			expectedF: 14,
			expectedK: 263.15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Calculando manualmente as conversões
			tempF := tt.tempC*1.8 + 32
			tempK := tt.tempC + 273.15

			// Verificando se os resultados estão próximos do esperado (considerando arredondamentos)
			if !almostEqual(tempF, tt.expectedF, 0.01) {
				t.Errorf("Conversão para Fahrenheit: obtido %.2f, esperado %.2f", tempF, tt.expectedF)
			}

			if !almostEqual(tempK, tt.expectedK, 0.01) {
				t.Errorf("Conversão para Kelvin: obtido %.2f, esperado %.2f", tempK, tt.expectedK)
			}
		})
	}
}

// Função auxiliar para comparar valores de ponto flutuante com tolerância
func almostEqual(a, b, tolerance float64) bool {
	diff := a - b
	if diff < 0 {
		diff = -diff
	}
	return diff <= tolerance
}
