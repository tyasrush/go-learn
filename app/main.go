package main

import (
	"fmt"

	"github.com/tyasrush/go-learn/config"
)

func main() {
	testCfg := config.GetPgConfig("config.yaml")
	fmt.Println(testCfg)
}
