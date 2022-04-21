package main

import "fmt"

type Node struct {
	key  string
	next *Node
	prev *Node
}

type LinkedList struct {
	head *Node
	tail *Node
}

func (l *LinkedList) Add(key string) {
	llist := &Node{
		key:  key,
		next: l.head,
	}

	if l.head != nil {
		l.head.prev = llist
	}

	l.head = llist

	curList := l.head
	for curList.next != nil {
		curList = curList.next
	}

	l.tail = curList
}

func main() {
	fmt.Println("testing")
}
