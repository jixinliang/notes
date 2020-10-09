package main

import "fmt"

type info struct {
	ID      string
	Name    string
	SurName string
}

var data = make(map[string]info)

func cat(key string) *info {
	_, ok := data[key]
	if ok {
		val := data[key]
		return &val
	}
	return nil
}

func add(key string, val info) bool {
	if key == "" {
		return false
	}
	if cat(key) == nil {
		data[key] = val
		return true
	}
	return false
}

func del(key string) bool {
	if cat(key) != nil {
		delete(data, key)
		return true
	}
	return false
}

func edit(key string, val info) bool {
	data[key] = val
	return true
}

func get() {
	for k, v := range data {
		fmt.Printf("key: %s value: %v\n", k, v)
	}
}

func main() {
	p1 := info{"1", "jack", "ji"}
	add("1", p1)
	//get()
	fmt.Println(del("1"))
	get()
}
