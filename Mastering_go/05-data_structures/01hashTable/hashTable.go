package main

import "fmt"

const SIZE = 15

type Node struct {
	Value int
	Next  *Node
}

type HashTable struct {
	Table map[int]*Node
	Size  int
}

func hashFunc(value, size int) int {
	return value % size
}

func insert(hashTable *HashTable, value int) int {
	//  does not check for duplicate values
	index := hashFunc(value, hashTable.Size)
	element := &Node{Value: value, Next: hashTable.Table[index]}
	hashTable.Table[index] = element
	return index
}

func traverse(hashTable *HashTable) {
	for k := range hashTable.Table {
		if hashTable.Table[k] != nil {
			t := hashTable.Table[k]
			for t != nil {
				fmt.Printf("%d -> ", t.Value)
				t = t.Next
			}
			fmt.Println()
		}
	}
}

func lookUp(hashTable *HashTable, value int) bool {
	index := hashFunc(value, hashTable.Size)
	if hashTable.Table[index] != nil {
		t := hashTable.Table[index]
		for t != nil {
			if t.Value == value {
				return true
			}
			t = t.Next
		}
	}
	return false
}

func main() {
	table := make(map[int]*Node, SIZE)
	hashTable := &HashTable{Table: table, Size: SIZE}
	fmt.Println("Number of spaces:", hashTable.Size)
	for i := 0; i < 120; i++ {
		insert(hashTable, i)
	}
	traverse(hashTable)
}
