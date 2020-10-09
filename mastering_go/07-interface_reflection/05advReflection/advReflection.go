package main

import (
	"fmt"
	"os"
	"reflect"
)

type t1 int
type t2 int

type a struct {
	X    int
	Y    float64
	Text string
}

func (a1 a) compareStruct(a2 a) bool {
	r1 := reflect.ValueOf(&a1).Elem()
	r2 := reflect.ValueOf(&a2).Elem()

	for i := 0; i < r1.NumField(); i++ {
		if r1.Field(i).Interface() != r2.Field(i).Interface() {
			return false
		}
	}
	return true
}

func getMethods(i interface{}) {
	r := reflect.ValueOf(i)
	t := r.Type()

	fmt.Printf("Type is %v\n", t)

	for j := 0; j < r.NumMethod(); j++ {
		m := r.Method(j).Type()
		fmt.Println(t.Method(j).Name, "---> ", m)
	}
}

func main() {
	x1 := t1(100)
	x2 := t2(100)

	fmt.Println("type of x1 is",reflect.TypeOf(x1))
	fmt.Println("type of x2 is",reflect.TypeOf(x2))

	var s struct{}
	r := reflect.New(reflect.ValueOf(&s).Type().Elem())
	fmt.Println("type of r is",reflect.TypeOf(r))

	a1 := a{1,1.1,"a1"}
	if a1.compareStruct(a1) {
		fmt.Println("Equal!")
	}

	var f *os.File
	getMethods(f)
}
