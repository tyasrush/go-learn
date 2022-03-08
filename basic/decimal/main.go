package main

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/tyasrush/go-learn/basic/decimal/example"
)

func main() {
	example.TestCalculateDecimal()

	zeroError := decimal.NewFromFloat(float64(0.1)).Div(decimal.Zero)
	fmt.Println("zero error : ", zeroError)
}
