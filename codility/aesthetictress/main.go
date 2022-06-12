package main

import (
	"log"
)

type ComparedFunc func(a, b int) bool

var ComparedFuncMap = map[string]ComparedFunc{
	"gt": func(a, b int) bool { return a > b },
	"lt": func(a, b int) bool { return a < b },
}

func IsPleasant(numbers []int) bool {
	cursor := "lt"
	if ComparedFuncMap["lt"](numbers[0], numbers[1]) {
		cursor = "gt"
	}

	slicedArray := numbers[1:]
	for index, value := range slicedArray {
		if cursor == "gt" && ComparedFuncMap["lt"](numbers[index], value) {
			cursor = "lt"
			continue
		}

		if cursor == "lt" && ComparedFuncMap["gt"](numbers[index], value) {
			cursor = "gt"
			continue
		}

		return false
	}

	return true
}

func Solution(numbers []int) int {
	if IsPleasant(numbers) {
		return 0
	}

	result := 0
	for i := 0; i < len(numbers); i++ {
		if (i + 1) == len(numbers)-1 {
			break
		}

		newArr := make([]int, len(numbers[0:i])+len(numbers[i+1:]))
		idx := 0
		for _, val := range numbers[0:i] {
			newArr[idx] = val
			idx += 1
		}

		for _, val := range numbers[i+1:] {
			newArr[idx] = val
			idx += 1
		}
		if IsPleasant(newArr) {
			result = result + 1
		}

	}

	if result > 0 {
		return result
	}

	return -1
}

func main() {
	testCases := []struct {
		param    []int
		expected int
	}{
		{
			[]int{3, 4, 5, 3, 7},
			3,
		},
		{
			[]int{1, 2, 3, 4},
			-1,
		},
		{
			[]int{1, 3, 1, 2},
			0,
		},
		{
			[]int{1, 1, 1, 1},
			-1,
		},
	}

	for _, test := range testCases {
		if Solution(test.param) != test.expected {
			log.Fatal("error gan")
		}
	}

}
