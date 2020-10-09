package func1

import "fmt"

func init() {
	fmt.Println("package func1 init()")
}

func FuncA()  {
	fmt.Println("From package func1/funcA()")
}