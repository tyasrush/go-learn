package main

import "fmt"

func Fib(n int, memo map[int]int) int {
	if _, ok := memo[n]; ok {
		return memo[n]
	}
	if n <= 2 {
		return 1
	}

	// memoization
	memo[n] = Fib(n-1, memo) + Fib(n-2, memo)
	return memo[n]
}

func main() {
	fmt.Println("Hello, 世界")

	// fibNumbers := []int64{1, 1, 2, 3, 5, 8, 13}

	// testFib := Fib(6)
	// penjabaran Fib(6)
	// iterasi 1 : Fib(5) + Fib(4)
	// iterasi 2 : (Fib(4) + Fib(3)) + (Fib(3) + Fib(2))
	// iterasi 3 : ((Fib(3) + Fib(2)) + (Fib(2) + Fib(1))) + ((Fib(2) + Fib(1)) + 1)
	// iterasi 4 : ((((Fib(2) + Fib(1)) + 1)  + (1 + 1 + 1 + 1 + 1)
	maps := make(map[int]int)
	// fmt.Println("fib ke 6 : ", Fib(6))
	fmt.Println("fib ke 50 : ", Fib(50, maps))
}
