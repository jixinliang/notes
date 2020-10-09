package main

import (
	"fmt"
	"unsafe"
)

func unSafe() {
	var val int64 = 5
	p1 := &val
	p2 := (*int32)(unsafe.Pointer(p1))

	fmt.Println("*p1: ", *p1)
	fmt.Println("*p2: ", *p2)
	*p1 = 5434123412312431212
	fmt.Println("val: ", val)
	fmt.Println("*p2", *p2) //  a 32-bit pointer cannot store a 64-bit integer
	*p1 = 54341234
	fmt.Println("val: ", val)
	fmt.Println("*p2", *p2)
}

func moreUnSafe() {
	arr := [...]int{0, 1, -2, 3, 4}
	ptr := &arr[0]
	fmt.Print(*ptr, " ")
	memAddr := uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])
	for i := 0; i < len(arr)-1; i++ {
		ptr = (*int)(unsafe.Pointer(memAddr))
		fmt.Print(*ptr, " ")
		memAddr = uintptr(unsafe.Pointer(ptr)) + unsafe.Sizeof(arr[0])
	}
}

func main() {
	moreUnSafe()
}
