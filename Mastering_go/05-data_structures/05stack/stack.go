package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var stack = new(Node)

func pushStack(value int) bool {
	if stack == nil {
		stack = &Node{value, nil}
		size = 1
		return true
	}

	tmp := &Node{value, nil}
	tmp.Next = stack
	stack = tmp
	size++
	return true
}

func popStack(node *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}
	if size == 1 {
		stack = nil
		size = 0
		return node.Value, true
	}

	stack = stack.Next
	size--
	return node.Value, true
}

func traverse(node *Node) {
	if size == 0 {
		fmt.Println("Empty stack!")
		return
	}

	for node != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println()
}



func main() {
	stack = nil
	v, b := popStack(stack)
	if b {
		fmt.Println(v)
	} else {
		fmt.Println("Pop() Failed!")
	}

	pushStack(100)
	pushStack(10)
	traverse(stack)
}
