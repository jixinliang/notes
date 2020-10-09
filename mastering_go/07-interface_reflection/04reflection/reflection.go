package main

import (
	"fmt"
	"os"
	"reflect"
)

type a struct {
	X int
	Y float64
	Z string
}

type b struct {
	F int
	G int
	H string
	I float64
}

func main() {
	x := 100
	xElem := reflect.ValueOf(&x).Elem()
	xType := xElem.Type()
	fmt.Printf("The type of x is %v\n", xType)

	A := a{100, 100.1, "struct a"}
	B := b{1, 2, "struct b", -2.3}

	var rValEle reflect.Value

	args := os.Args
	if len(args) == 1 {
		rValEle = reflect.ValueOf(&A).Elem()
	} else {
		rValEle = reflect.ValueOf(&B).Elem()
	}

	iType := rValEle.Type()
	fmt.Printf("type of r %v\n", iType)

	fmt.Printf("The %d fields of %v are:\n", rValEle.NumField(), iType)

	for i := 0; i < rValEle.NumField(); i++ {
		fmt.Printf("name of type of field: %s\n", iType.Field(i).Name)
		fmt.Printf("type of field of element is: %v\n", rValEle.Field(i).Type())
		fmt.Printf("interface of field of element is: %v\n", rValEle.Field(i).Interface())
	}

}
