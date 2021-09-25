package main

import (
	"fmt"
	"time"
)

func main() {
	n := time.Now()
	monthInt := fmt.Sprintf("%d", int(n.Month()))
	if n.Month() < 10 {
		monthInt = fmt.Sprintf("0%d", int(n.Month()))
	}
	strDateJoin := fmt.Sprintf("%d%s%d", n.Year(), monthInt, n.Day())
	fmt.Println("time - ", strDateJoin)
}
