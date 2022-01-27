package example

import (
	"fmt"
	"math/big"
	"strconv"

	"github.com/cockroachdb/apd"
	"github.com/ericlagergren/decimal"
)

func TestCalculateDecimal() {
	test := decimal.New(1234, 2)
	fmt.Println("test decimal new: ", test)

	simpleFloatFromBuiltIn, err := strconv.ParseFloat("4.940656458412465411112222233333", 128)
	if err != nil {
		panic(err)
	}

	amt, cond, err := apd.NewFromString("4.940656458412465411112222233333")
	if err != nil {
		panic(err)
	}

	a := apd.BaseContext.WithPrecision(30)

	bigint := apd.NewWithBigInt(big.NewInt(3451), -8)

	res := new(apd.Decimal)

	if _, err := a.Add(res, amt, bigint); err != nil {
		panic(err)
	}

	fmt.Println("parse float built in: ", simpleFloatFromBuiltIn)
	fmt.Printf("parse from apd: %v condition: %v apd bigint: %v result: %v\n", amt, cond, bigint, res)
}
