package tconv

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCalculte(t *testing.T) {
	req := require.New(t)

	tconv := calculate(0, "-f")
	req.Equal(float64(32), tconv)

	tconv = calculate(10, "-f")
	req.Equal(float64(50), tconv)

	tconv = calculate(-10, "-f")
	req.Equal(float64(14), tconv)

	tconv = calculate(0, "-c")
	req.Equal(float64(-17.78), tconv)

	tconv = calculate(10, "-c")
	req.Equal(float64(-12.22), tconv)

	tconv = calculate(-10, "-c")
	req.Equal(float64(-23.33), tconv)
}
