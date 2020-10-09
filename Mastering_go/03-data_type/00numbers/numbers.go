package main

import "fmt"

func main() {
	c1 := 12 + 1i
	c2 := complex(5, 7)
	fmt.Printf("type of c1: %T\n", c1)
	fmt.Printf("type of c2: %T\n", c2)

	c3 := complex64(c1 + c2)
	fmt.Printf("type of c3: %T\n", c3)
	fmt.Println("c3: ", c3)

	cZero := c3 - c3
	fmt.Println("cZero: ", cZero)

	fmt.Println("----------------------------------------")

	x := 12
	y := 5
	fmt.Printf("type of x: %T\n", x)
	fmt.Println(x)

	div := x / y
	fmt.Println("div: ", div)

	fmt.Println("----------------------------------------")

	var m float64 = 1.223
	fmt.Println("m: ", m)

	var m1 float64
	fmt.Println("m1: ", m1)

	n := 4 / 2.3
	fmt.Println("n: ", n)

	divFloat := float64(x) / float64(y)

	fmt.Printf("type of divFloat: %T\n", divFloat)
	fmt.Println(divFloat)

}

/*
Floating-point numbers
Go supports only two types of floating-point numbers: float32 and float64. The first
one provides about six decimal digits of precision, whereas the second one gives you 15
digits of precision.

Complex numbers
Similar to floating-point numbers, Go offers two complex number types named complex64
and complex128. The first one uses two float32: one for the real part and the other for
the imaginary part of the complex number, whereas complex128 uses two float64.

*/
