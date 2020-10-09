package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

var size = 0
var queue = new(Node)

func pushNode(node *Node, value int) bool {
	if queue == nil {
		queue = &Node{value, nil}
		size++
		return true
	}

	// queue not null , append queue to the next of node
	// create a new node that is placed in front of the current queue. After that,
	// the head of the queue becomes the node that was just created
	node = &Node{value, nil}
	node.Next = queue
	queue = node
	size++

	return true
}

func popNode(node *Node) (int, bool) {
	if size == 0 {
		return 0, false
	}

	if size == 1 {
		queue = nil
		size--
		return node.Value, true
	}

	tmp := node
	for node.Next != nil {
		tmp = node
		node = node.Next
	}

	value := tmp.Next.Value
	tmp.Next = nil
	size--

	return value, true
}

func traverse(node *Node) {
	if size == 0 {
		fmt.Println("Empty list!")
		return
	}
	for node != nil {
		fmt.Printf("%d -> ", node.Value)
		node = node.Next
	}
	fmt.Println()
}

func main() {
	fmt.Println("size:", size)
	queue = nil
	pushNode(queue, 10)
	fmt.Println("size:", size)
	traverse(queue)

	v, b := popNode(queue)
	if b {
		fmt.Println("Poped:", v)
	}
	fmt.Println("size:", size)

	for i := 0; i < 5; i++ {
		pushNode(queue, i)
	}
	traverse(queue)
	fmt.Println("size:", size)


}
