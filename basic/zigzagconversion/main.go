package main

import "strings"

func convert(s string, numRows int) string {
	// Base conditions
	if s == "" || numRows <= 0 {
		return ""
	}
	if numRows == 1 {
		return s
	}
	// Resultant string
	var sb strings.Builder
	// Step size
	step := int(2*numRows - 2)
	// Loop for each row
	for i := 0; i < numRows; i++ {
		// Loop for each character in the row
		for j := i; j < len(s); j += step {
			sb.WriteString(string(s[j]))
			if i != 0 && i != numRows-1 && (j+step-2*i) < len(s) {
				sb.WriteString(string(s[j+step-2*i]))
			}
		}
	}
	return sb.String()
}

func main() {
}
