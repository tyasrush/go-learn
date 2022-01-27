package example

import (
	"fmt"

	"github.com/ericlagergren/decimal"
)

func TestCalculateDecimal() {
	test := decimal.New(1234, 2)
	fmt.Println("test decimal new: ", test)
}
