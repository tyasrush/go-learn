package main

import (
	"fmt"
	"math"
)

func main() {
	testNum := 529
	counting := false
	var maxGap float64
	var currentGap float64
	for testNum != 0 {
		if counting == false { // for the first "1"
			if (testNum & 1) == 1 { // note: cannot use n&1 withoug "()"
				counting = true // start to count
			}
		} else { // counting = true
			if (testNum & 1) == 0 { // note: cannot use n&1 withoug "()"
				currentGap += 1
			} else { // N & 1 == 1
				maxGap = math.Max(maxGap, currentGap)
				currentGap = 0 // reset
			}
		}

		// fmt.Println(testNum)
		testNum = testNum >> 1 // shift by one (right side)
		// note: cannot just write "N >> 1"
	}

	fmt.Println(maxGap)
}
