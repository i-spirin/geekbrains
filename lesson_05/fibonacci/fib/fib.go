package fib

// Fib function to get Fibonacci number
func Fib(n int) int {
	if n <= 2 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)
}
