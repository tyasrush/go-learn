package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// example.TestCalculateDecimal()

	zeroError := decimal.NewFromFloat(float64(0.1)).Div(decimal.Zero)
	fmt.Println("zero error : ", zeroError)

	testingString := "lohe lohe"
	if len(testingString) == 0 {
		fmt.Println("kosong dong")
	}
}
