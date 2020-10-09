package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing/quick"
	"time"
)

type point3D struct {
	X, Y, Z int8
	S       float32
}

func main() {
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	typeOf := reflect.TypeOf(point3D{})

	value, _ := quick.Value(typeOf, ran)
	fmt.Println(value)
}
