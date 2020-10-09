package main

import "fmt"

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

var root = new(Node)

func addNode(node *Node, value int) int {
	if root == nil {
		node = &Node{value, nil, nil}
		root = node
		return 0
	}

	if value == node.Value {
		fmt.Println("Already existed value:", value)
		return -1
	}

	if node.Next == nil {
		tmp := node
		node.Next = &Node{value, tmp, nil}
		return -2
	}
	return addNode(node.Next, value)
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

func reverse(node *Node) {
	if node == nil {
		fmt.Println("-> Empty list!")
		return
	}

	// keep the value to reverse
	tmp := node
	for node != nil {
		tmp = node
		node = node.Next
	}

	for tmp.Prev != nil {
		fmt.Printf("%d -> ", tmp.Value)
		tmp = tmp.Prev
	}

	fmt.Printf("%d -> ", tmp.Value)
	fmt.Println()
}

func size(node *Node) int {
	if node == nil {
		fmt.Println("-> Empty list!")
		return 0
	}

	n := 0
	for node != nil {
		n++
		node = node.Next
	}
	return n
}

func lookUp(node *Node, value int) bool {
	if root == nil {
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

func main() {
	fmt.Println("root:", root)
	root = nil
	traverse(root)

	addNode(root, 1)
	addNode(root, 1)
	traverse(root)

	addNode(root, 10)
	addNode(root, 15)
	traverse(root)

	fmt.Println("size:",size(root))
	addNode(root,100)
	addNode(root,0)
	addNode(root,0)
	traverse(root)
	reverse(root)

}
