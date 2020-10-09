package main

import "fmt"

func mapDemo() {
	m1 := make(map[string]int)
	fmt.Println(m1)

	m2 := map[string]int{
		"k1": 1,
		"k2": 2,
	}
	fmt.Println(m2)
	fmt.Println("get k1:", m2["k1"])

	delete(m2, "k1")
	// multiple times does not make any difference and does not generate any warning messages
	delete(m2, "k1")
	fmt.Println("after delete k1:", m2)
	fmt.Println("get k1:", m2["k1"])

	_, ok := m2["ifExists"]
	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does not exists")
	}

	for k, v := range m2 {
		fmt.Println(k, ":", v)
	}
}

func failMap() {
	m := map[string]int{}
	m["k1"] = 1
	fmt.Println(m)
	m = nil
	fmt.Println(m)
	//m["k2"] = 2 // Assignment to entry may panic because of 'nil' map

}

func main() {
	failMap()
}
