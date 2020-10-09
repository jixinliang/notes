package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func traverse(tree *Tree) {
	if tree == nil {
		return
	}
	traverse(tree.Left)
	fmt.Print(tree.Value, " ")
	traverse(tree.Right)
}

func create(n int) *Tree {
	var tree *Tree
	rand.Seed(time.Now().Unix())

	// binary tree double n
	for i := 0; i < 2*n; i++ {
		tmp := rand.Intn(2 * n)
		tree = insert(tree, tmp)
	}
	return tree
}

func insert(tree *Tree, val int) *Tree {
	// root of the tree
	if tree == nil {
		return &Tree{nil, val, nil}
	}
	if val == tree.Value {
		return tree
	}
	if val < tree.Value {
		tree.Left = insert(tree.Left, val)
		return tree
	}
	tree.Right = insert(tree.Right, val)
	return tree
}

func main() {
	tree := create(10)
	fmt.Println("the value of the root of tree is:", tree.Value)
	traverse(tree)
	fmt.Println()

	tree = insert(tree, -10)
	tree = insert(tree, -2)
	traverse(tree)
	fmt.Println()

	fmt.Println("the value of the root of the tree is:", tree.Value)
}
