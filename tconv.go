package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var result float64

	args := os.Args

	num, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		panic(err.Error())
	}

	resultUnit := args[2]
	if resultUnit != "-f" && resultUnit != "-c" {
		panic("Incorrect translation unit")
	}

	if resultUnit == "-c" {
		result = (num - 32) / 1.8
	} else {
		result = num*1.8 + 32
	}

	// round up to 2 digits after point
	fmt.Println(math.Round(result*100) / 100)
}
