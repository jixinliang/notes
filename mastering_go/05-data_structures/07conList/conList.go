package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func getList(l *list.List) {
	// negative order
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	// positive order
	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	val := list.New()
	e1 := val.PushBack("one")
	e2 := val.PushBack("two")
	val.PushFront("three")
	val.InsertBefore("four", e1)
	val.InsertAfter("five", e2)

	val.Remove(e2)
	val.PushBackList(val)

	getList(val)
	val.Init() // Init initializes or clears list l.
	fmt.Println("len of list:", val.Len())

	fmt.Println("Pushing after init():", *val)

	for i := 1; i <= 20; i++ {
		val.PushFront(strconv.Itoa(i))
	}
	getList(val)
}
