package main

func fib1(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fib1(n-1) + fib1(n-2)
}

func fib2(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 2 // bug
	}
	return fib2(n-1) + fib2(n-2)
}

func s1(s string) int {
	if s == "" {
		return 0
	}
	n := 1 // bug
	for range s {
		n++
	}
	return n
}

func s2(s string) int  {
	return len(s)
}

func main() {

}

/*
go test -cover -v
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
go tool cover -html=coverage.out -o coverage.html

*/