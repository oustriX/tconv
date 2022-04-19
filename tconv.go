package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("you missed a few arguments")
		fmt.Printf("must: 2. you entered: %d\n", len(os.Args)-1)
		return
	}

	num, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("can't parse number in your first argument")
		fmt.Printf("error message: %s\n", err.Error())
		return
	}

	resultUnit, err := getResultUnit()
	if err != nil {
		fmt.Println("error when trying to convert")
		fmt.Printf("error message: %d\n", err.Error())
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
