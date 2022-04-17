package tconv

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculte(t *testing.T) {
	req := require.New(t)
	t.Run("Zero to Fahrenheit", func(t *testing.T) {
		tconv := calculate(0, "-f")
		req.Equal(float64(32), tconv)
	})

	t.Run("Positive to Fahrenheit", func(t *testing.T) {
		tconv := calculate(10, "-f")
		req.Equal(float64(50), tconv)
	})

	t.Run("Negative to Fahrenheit", func(t *testing.T) {
		tconv := calculate(-10, "-f")
		req.Equal(float64(14), tconv)
	})

	t.Run("Zero to Celsius", func(t *testing.T) {
		tconv := calculate(0, "-c")
		req.Equal(float64(-17.78), tconv)
	})

	t.Run("Positive to Celsius", func(t *testing.T) {
		tconv := calculate(10, "-c")
		req.Equal(float64(-12.22), tconv)
	})

	t.Run("Negative to Celsius", func(t *testing.T) {
		tconv := calculate(-10, "-c")
		req.Equal(float64(-23.33), tconv)
	})
}
