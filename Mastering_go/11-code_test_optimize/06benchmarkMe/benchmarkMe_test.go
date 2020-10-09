package main

import "testing"

var result int

// not run automatically because they begin with a lowercase b instead of an uppercase B
func benchmarkfib1(b *testing.B, n int) {
	var res int
	for i := 0; i < b.N; i++ {
		res = fib1(n)
	}
	result = res
}

func benchmarkfib2(b *testing.B, n int) {
	var res int
	for i := 0; i < b.N; i++ {
		res = fib2(n)
	}
	result = res
}

func benchmarkfib3(b *testing.B, n int) {
	var res int
	for i := 0; i < b.N; i++ {
		res = fib3(n)
	}
	result = res
}

func Benchmark20fib1(b *testing.B) {
	benchmarkfib1(b, 20)
}

func Benchmark20fib2(b *testing.B) {
	benchmarkfib2(b, 20)
}

func Benchmark20fib3(b *testing.B) {
	benchmarkfib3(b, 20)
}

func Benchmark40fib1(b *testing.B) {
	benchmarkfib1(b, 40)
}

func Benchmark40fib2(b *testing.B) {
	benchmarkfib2(b, 40)
}

func Benchmark40fib3(b *testing.B) {
	benchmarkfib3(b, 40)
}