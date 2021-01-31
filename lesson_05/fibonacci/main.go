package main

import (
	"fmt"
	"os"
)

func main() {
	var n int
	cache := make(map[int]int)

	fibCache := func(n int) int {
		result, isExists := cache[n]
		if isExists == true {
			return result
		}
		result = fib(n)
		cache[n] = result
		return result
	}

	for {
		fmt.Print("Insert number: ")
		fmt.Fscan(os.Stdin, &n)

		fmt.Printf("%d\n", fibCache(n))
	}

}

func fib(n int) int {
	if n <= 2 {
		return 1
	}

	return fib(n-1) + fib(n-2)
}
