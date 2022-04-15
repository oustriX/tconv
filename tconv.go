package tconv

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func tconv() {
	num, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		panic(err.Error())
	}

	resultUnit, err := getResultUnit()
	if err != nil {
		panic(err.Error())
	}

	result := calculate(num, resultUnit)

	fmt.Println(result)

}

func calculate(num float64, resultUnit string) float64 {
	var result float64

	if resultUnit == "-c" {
		result = (num - 32) / 1.8
	} else {
		result = num*1.8 + 32
	}

	// round up to 2 digits after point
	return math.Round(result*100) / 100
}

func getResultUnit() (string, error) {
	unit := os.Args[2]
	if unit != "-f" && unit != "-c" {
		return unit, errors.New("incorrect translation unit")
	}
	return unit, nil
}
