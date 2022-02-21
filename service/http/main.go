package main

import (
	"fmt"

	"github.com/tyasrush/go-learn/service/http/entity"
)

func main() {
	fmt.Println("testing aja")
	truck := entity.Truck{
		Name: "testing name",
	}

	fmt.Println("truck: ", truck)
}
