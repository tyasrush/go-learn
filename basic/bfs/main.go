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

// https://faun.pub/2-different-ways-to-implement-bfs-in-golang-8399f5d2452d
func bfs(root *Node) *Node {
	// root checking
	if root == nil {
		return root
	}

	queue := []NodeLevel{
		{
			root,
			0,
		},
	}
	visited := []*Node{}

	for len(queue) > 0 {
		vertex := queue[0]
		node, level := vertex.node, vertex.level
		visited = append(visited, node)
		queue = queue[1:]

		fmt.Printf("level: %d visited: %v queue: %v\n", level, visited, queue)

		if node.Left != nil {
			leftNode := NodeLevel{
				node:  node.Left,
				level: level + 1,
			}

			queue = append(queue, leftNode)
		}

		if node.Right != nil {
			rightNode := NodeLevel{
				node:  node.Right,
				level: level + 1,
			}

			queue = append(queue, rightNode)
		}
	}
	for _, v := range visited {
		fmt.Printf("all visited node: %v\n", v.Val)
	}

	return root
}

func main() {

	node1 := &Node{
		Val: 1,
		Left: &Node{
			Val: 2,
			Left: &Node{
				Val: 3,
				Right: &Node{
					Val: 7,
					Right: &Node{
						Val: 8,
					},
				},
			},
			Right: &Node{
				Val: 4,
				Left: &Node{
					Val: 9,
				},
				Right: &Node{
					Val: 10,
				},
			},
		},
		Right: &Node{
			Val: 5,
			Right: &Node{
				Val: 6,
			},
		},
	}

	bfs(node1)

}
