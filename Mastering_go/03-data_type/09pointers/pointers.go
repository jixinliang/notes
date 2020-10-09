package main

import "fmt"

func main() {
	v1 := -1
	v2 := 2
	pV1 := &v1
	pV2 := &v2
	fmt.Println("addr of pV1:", pV1)
	fmt.Println("addr of pV2:", pV2)
	fmt.Println("value of pV1:", *pV1)
	fmt.Println("value of pV2:", *pV2)

	*pV1 = 123
	*pV1--
	fmt.Println("v1:",v1)


}

/*
use * to get the value of a pointer, which is called deReferencing the pointer,
and & to get the memory address of a non-pointer variable.
*/
