package main

import (
	"fmt"

	"github.com/shopspring/decimal"
)

func main() {
	// example.TestCalculateDecimal()

	// zeroError := decimal.NewFromFloat(float64(0.1)).Div(decimal.Zero)
	// fmt.Println("zero error : ", zeroError)

	roundDec := decimal.NewFromFloat(0.989797).Round(2)
	fmt.Println("round testing: ", roundDec)

	testingDiv := decimal.Zero.Mul(decimal.Zero).Round(2)
	fmt.Println("zero ", testingDiv)
}
