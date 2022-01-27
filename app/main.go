package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()
	recv := now.Add(-15 * time.Minute)
	fmt.Printf("now: %v recv: %v \n", now.UnixMicro(), recv.UnixMilli())
}
