package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
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

	logger := log.New(os.Stdout, "app: ", log.Lshortfile)
	logger.Print("testing trace")

	var varitf interface{}
	varitf = 1

	fmt.Println(varitf)

	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("not oke gaes")
	}

	fmt.Printf("pc - %v file - %s line - %d\n", pc, file, line)
}
