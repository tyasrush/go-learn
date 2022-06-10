package main

import (
	"fmt"
	"math"
)

func main() {
	// A[0] = -5
	//   A[1] = -3
	//   A[2] = -1
	//   A[3] =  0
	//   A[4] =  3
	//   A[5] =  6
	param := []int{-5, -3, -1, 0, 3, 6}

	var mapInteger = make(map[float64]bool)
	for i := 0; i < len(param); i++ {
		cursor := math.Abs(float64(param[i]))
		if _, ok := mapInteger[cursor]; !ok {
			mapInteger[cursor] = false
		}
	}

	fmt.Println(len(mapInteger))
}
