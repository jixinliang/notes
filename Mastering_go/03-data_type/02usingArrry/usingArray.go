package main

import "fmt"

func myArray() {
	myArr := [4]int{1, 2, 3, 4}
	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	threeD := [2][2][2]int{{{1, 2}, {3, 4}}, {{5, 6}, {7, 8}}}
	fmt.Println("len of myArr:", len(myArr))
	fmt.Println("the first element of twoD:", twoD[0][0])
	fmt.Println("the len of threeD：", len(threeD))

	for _, x := range threeD {
		for _, y := range x {
			for _, z := range y {
				fmt.Print(z, " ")
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func mySlice() {
	mySlice1 := make([]int, 20)
	for val := range mySlice1 {
		fmt.Print(val, " ")
	}
	fmt.Println()
	mySlice1 = nil
	fmt.Println("mySlice1:", mySlice1)
	fmt.Printf("type of mySlice1 %T\n", mySlice1)

	mySlice1 = append(mySlice1, 1, 2, 3, 4)
	fmt.Print(mySlice1, " ")
	fmt.Println()

	fmt.Println("the first element:", mySlice1[0])
	fmt.Println("the last element:", mySlice1[len(mySlice1)-1])
	fmt.Println("the second to third", mySlice1[1:3])

	mySlice2 := mySlice1[1:3]
	fmt.Println("mySlice2:", mySlice2)
	mySlice2[0]=-100
	mySlice2[1]=123
	fmt.Println("mySlice1:",mySlice1)
	fmt.Println("mySlice2:",mySlice2)

}

func main() {
	mySlice()
}
