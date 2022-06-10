package main

import (
	"log"
)

func solution(param []int) int {
	//special case
	if len(param) == 0 {
		return 1
	}

	// Using "set" to check if an element has appeared
	// note: need to "import java.util.*" (important)
	var maps = make(map[int]bool)

	// add elements into the set
	for _, val := range param {
		maps[val] = false
	}

	// note: the missing number is not possible bigger than (A.length)
	//       because there are only (A.length) numbers
	for i := 1; i <= len(param); i++ {
		if _, ok := maps[i]; !ok {
			return i
		}
	}

	// means: there are no missing numbers from 1 to A.length
	// Therefore, the missing number is "A.length+1" (important)
	return len(param) + 1
}

func main() {
	testCases := []struct {
		param    []int
		expected int
	}{
		{
			[]int{1, 3, 6, 4, 1, 2},
			5,
		},
		{
			[]int{1, 2, 3},
			4,
		},
		{
			[]int{-1, -3},
			1,
		},
	}

	for _, test := range testCases {
		res := solution(test.param)
		if res != test.expected {
			log.Fatalf("expected: %d actual: %d", test.expected, res)
		}
	}
}
