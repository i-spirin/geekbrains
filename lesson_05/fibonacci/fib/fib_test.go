package fib_test

import (
	"testing"

	"github.com/i-spirin/geekbrains/lesson_05/fibonacci/fib"
)

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
		result := fib.Fib(el.argument)
		if result != el.expected {
			t.Fatalf("TestFib failed %d given, %d got, %d expected", el.argument, result, el.expected)
		}
	}
}

var GlobalF int

func BenchmarkFib(t *testing.B) {
	x := 0
	for i := 0; i < t.N; i++ {
		x = fib.Fib(10)
	}
	GlobalF = x
}
