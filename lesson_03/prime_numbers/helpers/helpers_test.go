package helpers_test

import (
	"testing"

	"github.com/i-spirin/geekbrains/lesson_03/prime_numbers/helpers"
)

func TestIsPrime(t *testing.T) {
	cases := []struct {
		argument int
		expected bool
	}{
		{
			argument: 1,
			expected: false,
		},
		{
			argument: 2,
			expected: true,
		},
		{
			argument: 3,
			expected: true,
		},
		{
			argument: 4,
			expected: false,
		},
		{
			argument: 90,
			expected: false,
		},
		{
			argument: 97,
			expected: true,
		},
	}

	var result bool
	for _, el := range cases {
		result = helpers.IsPrime(el.argument)
		if result != el.expected {
			t.Fatalf("TestIsPrime failed: %d given, %t got, %t expected", el.argument, result, el.expected)
		}
	}
}
