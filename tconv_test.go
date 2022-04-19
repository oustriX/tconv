package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculte(t *testing.T) {
	req := require.New(t)
	t.Run("Zero to Fahrenheit", func(t *testing.T) {
		tconv := calculate(0, "-f")
		req.Equal(32.0, tconv)
	})

	t.Run("Positive to Fahrenheit", func(t *testing.T) {
		tconv := calculate(10, "-f")
		req.Equal(50.0, tconv)
	})

	t.Run("Negative to Fahrenheit", func(t *testing.T) {
		tconv := calculate(-10, "-f")
		req.Equal(14.0, tconv)
	})

	t.Run("Zero to Celsius", func(t *testing.T) {
		tconv := calculate(0, "-c")
		req.Equal(-17.78, tconv)
	})

	t.Run("Positive to Celsius", func(t *testing.T) {
		tconv := calculate(10, "-c")
		req.Equal(-12.22, tconv)
	})

	t.Run("Negative to Celsius", func(t *testing.T) {
		tconv := calculate(-10, "-c")
		req.Equal(-23.33, tconv)
	})
}
