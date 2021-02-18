package main

import "testing"

func TestFib(t *testing.T) {
	cases := []struct {
		argument int
		expected int
	}{
		{
			argument: 1,
			expected: 1,
		},
		{
			argument: 2,
			expected: 1,
		},
		{
			argument: 10,
			expected: 55,
		},
	}

	for _, el := range cases {
		result := fib(el.argument)
		if result != el.expected {
			t.Fatalf("TestFib failed %d given, %d got, %d expected", el.argument, result, el.expected)
		}
	}
}

var GlobalF int

func BenchmarkFib(t *testing.B) {
	x := 0
	for i := 0; i < t.N; i++ {
		x = fib(10)
	}
	GlobalF = x
}
