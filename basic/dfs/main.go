package main

import "fmt"

type Node struct {
	Val   int
	Right *Node
	Left  *Node
	Next  *Node
}

type NodeLevel struct {
	node  *Node
	level int
}

func dfs(root *Node, visited []*Node, queue []NodeLevel) *Node {
	// root checking
	if root == nil {
		return root
	}
	// level := queue[0].level + 1
	queue = queue[1:]

	if root.Left != nil {
		visited = append(visited, root.Left)
		return dfs(root.Left, visited, queue)
	}

	if root.Right != nil {
		visited = append(visited, root.Right)
		return dfs(root.Left, visited, queue)
	}

	// for len(queue) > 0 {
	// 	vertex := queue[0]
	// 	node, level := vertex.node, vertex.level
	// 	visited = append(visited, node)
	// 	queue = queue[1:]

	// 	fmt.Printf("level: %d visited: %v queue: %v\n", level, visited, queue)

	// 	if node.Left != nil {
	// 		leftNode := NodeLevel{
	// 			node:  node.Left,
	// 			level: level + 1,
	// 		}

	// 		queue = append(queue, leftNode)
	// 	}

	// 	if node.Right != nil {
	// 		rightNode := NodeLevel{
	// 			node:  node.Right,
	// 			level: level + 1,
	// 		}

	// 		queue = append(queue, rightNode)
	// 	}
	// }
	// for _, v := range visited {
	// 	fmt.Printf("all visited node: %v\n", v.Val)
	// }

	return root
}

func main() {
	fmt.Println("testing")
}
