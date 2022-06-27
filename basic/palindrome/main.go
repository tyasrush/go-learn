package main

import (
	"log"
)

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}

	var (
		maxStr     = ""
		currStr    = ""
		maxLength  = 0
		currLength = 0
	)

	for i := 0; i < len(s); i++ {
		currStr = expand(s, i, i)
		currLength = len(currStr)

		// update maximum length palindromic substring if odd length
		// palindrome has a greater length

		if currLength > maxLength {
			maxLength = currLength
			maxStr = currStr
		}

		// Find the longest even length palindrome with `str[i]` and `str[i+1]`
		// as midpoints. Note that an even length palindrome has two
		// midpoints.

		currStr = expand(s, i, i+1)
		currLength = len(currStr)

		// update maximum length palindromic substring if even length
		// palindrome has a greater length

		if currLength > maxLength {
			maxLength = currLength
			maxStr = currStr
		}
	}

	return maxStr
}

func expand(str string, low, high int) string {
	for low >= 0 && high < len(str) && string(str[low]) == string(str[high]) {
		low--
		high++
	}

	return str[low+1 : high]
}

func main() {
	testCases := []struct {
		param          string
		expectedResult string
	}{

		{
			"babad",
			"bab",
		},
		{
			"cbbd",
			"bb",
		},
		{
			"ccc",
			"ccc",
		},
		{
			"bb",
			"bb",
		},
	}

	for _, test := range testCases {
		if test.expectedResult != longestPalindrome(test.param) {
			log.Fatalf("error on param %s exp result %s\n", test.param, test.expectedResult)
		}
	}
}
