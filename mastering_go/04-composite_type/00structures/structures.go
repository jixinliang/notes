package main

import "fmt"

type XYZ struct {
	x, y, z int
}

func xyzDemo() {
	s1 := XYZ{4, 5, 6}
	s2 := XYZ{z: 10, y: 12}
	fmt.Println(s1)
	fmt.Println(s2)

	var s3 XYZ
	fmt.Println(s3.x, s3.y)

	s4 := [4]XYZ{}
	s4[0] = s1
	s4[2] = s2
	fmt.Println(s4)
}

type info struct {
	Name    string
	SurName string
	Height  uint32
}

func newInfoPtr(n, s string, h uint32) *info {
	if h > 300 {
		h = 0
	}
	return &info{n, s, h}
}

func newInfo(n, s string, h uint32) info {
	if h > 300 {
		h = 0
	}
	return info{n, s, h}
}

func infoDemo() {
	s1 := newInfoPtr("Jack", "Ji", 182)
	s2 := newInfo("Gina", "Gi", 172)

	fmt.Println(s1.Name)
	fmt.Println(s2.Name)
	fmt.Println(s1)
	fmt.Println(s2)

	s3 := new(info)
	fmt.Println(s3)

	s4 := new([]info)
	fmt.Println(s4)
}

// emulate tuple
func retThree(x int) (int, int, int) {
	return 2 * x, x * x, -x
}

func main() {
	n1, n2, n3 := retThree(10)
	fmt.Println(n1, n2, n3)
}

/*
if you use key value style you do not need to define
an initial value for every field of the structure

*/
