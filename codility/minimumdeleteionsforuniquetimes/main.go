package main

import "log"

func minDeletions(s string) int {
	// Add to frequency map
	m := make(map[rune]int)
	for _, v := range s {
		m[v]++
	}

	ans := 0
	// Store the unique frequencies
	set := make(map[int]struct{})
	for _, v := range m {

		// if the frequenct exists, then delete one char
		for Contains(set, v) {
			ans++
			v--
		}

		// Add the frequency after deletion if it is not zero
		if v != 0 {
			set[v] = struct{}{}
		}
	}

	return ans
}

func Contains(m map[int]struct{}, key int) bool {
	if _, ok := m[key]; ok {
		return true
	}

	return false
}

func main() {
	testCases := []struct {
		param    string
		expected int
	}{
		{
			"aab",
			0,
		},
		{
			"aaabbbcc",
			2,
		},
		{
			"ceabaacb",
			2,
		},
		{
			"example",
			4,
		},
	}

	for _, test := range testCases {
		if minDeletions(test.param) != test.expected {
			log.Fatal("error - ", test.param)
		}
	}
}
