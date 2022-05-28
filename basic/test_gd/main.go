package main

import "fmt"

func sumOfPair(numbers []int, expectedSum int) [][2]int {
	result := [][2]int{}
	for i, num := range numbers {
		itemNumber := [2]int{num}
		isCorrect := false
		for j := i + 1; j < len(numbers); j++ {
			if num+numbers[j] == expectedSum {
				itemNumber[1] = numbers[j]
				isCorrect = true
				break
			}
		}

		if isCorrect {
			result = append(result, itemNumber)
			isCorrect = false
		}
	}

	return result
}

func TwoSumDynamic(numbers []int, expectedSum int) [][2]int {
	result := [][2]int{}
	sumsMap := make(map[int]int)
	for _, number := range numbers {
		sumsMinNumber := expectedSum - number

		if _, ok := sumsMap[sumsMinNumber]; ok {
			result = append(result, [2]int{number, sumsMinNumber})
		}

		sumsMap[number] = number
	}
	return result
}

func main() {
	testCases := []struct {
		name           string
		numbers        []int
		expectedNumber int
	}{
		{
			"initial case",
			[]int{4, -2, -4, 2, 1, 5, -3},
			-2,
		},
	}

	for _, test := range testCases {
		result := TwoSumDynamic(test.numbers, test.expectedNumber)
		fmt.Println(result)
	}
}
