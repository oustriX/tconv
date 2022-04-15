package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var target float64

	args := os.Args

	num, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		panic(err.Error())
	}

	targetUnit := args[2]
	if targetUnit != "-f" && targetUnit != "-c" {
		panic("Incorrect translation unit")
	}

	if targetUnit == "-c" {
		target = (num - 32) / 1.8
	} else {
		target = num*1.8 + 32
	}

	// round up to 2 digits after point
	fmt.Println(math.Round(target*100) / 100)
}
