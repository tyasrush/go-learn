package main

import "fmt"

type Node struct {
	Value string
	Next  interface{}
}

func PrintAllNode(param Node) {
	idx := 0
	for param.Value != "" {
		fmt.Println("index: ", idx)
		fmt.Println(param.Value)
		param = param.Next.(Node)
		idx += 1
	}
}

func main() {
	fmt.Println("testing")
	a := Node{Value: "A"}
	b := Node{Value: "B"}
	c := Node{Value: "C"}

	a.Next = Node{Value: "lah"}
	b.Next = c

	fmt.Println(a.Next)
	fmt.Println(b.Next)

	PrintAllNode(a)
}
