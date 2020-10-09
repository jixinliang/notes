package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

var root = new(Node)

func addNode(t *Node, v int) int {
	if root == nil {
		t = &Node{v, nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exists:", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &Node{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverse(node *Node) {
	if node == nil {
		fmt.Println("-> Empty list!")
		return
	}

	for node != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func lookUp(node *Node, value int) bool {
	if root == nil {
		node = &Node{value, nil}
		root = node
		return false
	}

	if value == node.Value {
		return true
	}
	if node.Next == nil {
		return false
	}

	return lookUp(node.Next, value)
}

func size(node *Node) int {
	if node == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	i := 0
	for node != nil {
		i++
		node = node.Next
	}
	return i
}

func main() {
	fmt.Println("root:", root)
	root = nil
	traverse(root)

	addNode(root, 1)
	addNode(root, -1)
	addNode(root, -1)
	traverse(root)

	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverse(root)

	addNode(root, 100)
	traverse(root)

	if lookUp(root, -100) {
		fmt.Println("Node exists!")
	} else {
		fmt.Println("Node does not exist!")
	}

	fmt.Println("size:", size(root))
}
