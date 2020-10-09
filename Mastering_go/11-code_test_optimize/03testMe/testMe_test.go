package main

import "testing"


func Test_fib2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib2(tt.args.n); got != tt.want {
				t.Errorf("fib2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_s1(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := s1(tt.args.s); got != tt.want {
				t.Errorf("s1() = %v, want %v", got, tt.want)
			}
		})
	}
}

