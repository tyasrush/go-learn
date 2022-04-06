package main

import (
	"fmt"
	"math"
)

// Reference : https://stackoverflow.com/questions/51146294/reverse-a-number-without-making-it-a-string-in-javascript
func ReverseNumber(param int64) int64 {
	var revNumber int64 = 0
	for param > 0 {
		revNumber = (revNumber * 10) + (param % 10)
		divParam := param / 10
		param = int64(math.Floor(float64(divParam)))
	}
	return revNumber
}

func main() {
	fmt.Println("testing")

	// 789 => reverse process
	// first => (0 * 10) + (789 % 10)
	// first => 0 + 9
	// floor the param number
	// 789 / 10 => floor => 78.9 jadi 78

	// second => (9 * 10) + (78 % 10)
	// second => 90 + 8
	// floor the param number
	// 78 / 10 => floor => 7.8 jadi 7

	// third => (98 * 10) + (7 % 10)
	// third => 980 + 7
	// floor the param number
	// 7 / 10 => floor => 0.7 jadi 0

	// finish
	testReverseNumber := ReverseNumber(789)
	fmt.Printf("reverse number: %v\n", testReverseNumber)
}
