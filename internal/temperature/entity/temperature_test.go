package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTemperature(t *testing.T) {
	celsiusInput := 0.0
	expectedFahrenheit := 32.0
	expectedKelvin := 273.0

	temp := NewTemperature(celsiusInput)

	assert.Equal(t, celsiusInput, temp.Celsius(), "Celsius value should be equal to the input")
	assert.Equal(t, expectedFahrenheit, temp.Fahrenheit(), "Fahrenheit conversion did not match expected value")
	assert.Equal(t, expectedKelvin, temp.Kelvin(), "Kelvin conversion did not match expected value")
}

func TestTemperatureConversions(t *testing.T) {
	tests := []struct {
		celsius            float64
		expectedFahrenheit float64
		expectedKelvin     float64
	}{
		{0, 32, 273},
		{100, 212, 373},
		{-40, -40, 233},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			temp := NewTemperature(tt.celsius)
			assert.Equal(t, tt.expectedFahrenheit, temp.Fahrenheit(), "Fahrenheit conversion did not match expected value")
			assert.Equal(t, tt.expectedKelvin, temp.Kelvin(), "Kelvin conversion did not match expected value")
			assert.Equal(t, tt.celsius, temp.Celsius(), "Celsius value should be equal to the input")
		})
	}
}
