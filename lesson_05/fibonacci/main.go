package main

import (
	"fmt"
	"os"

	"github.com/i-spirin/geekbrains/lesson_05/fibonacci/fib"
)

func main() {
	var n int
	cache := make(map[int]int)

	fibCache := func(n int) int {
		result, isExists := cache[n]
		if isExists == true {
			return result
		}
		result = fib.Fib(n)
		cache[n] = result
		return result
	}

	for {
		fmt.Print("Insert number: ")
		fmt.Fscan(os.Stdin, &n)

		fmt.Printf("%d\n", fibCache(n))
	}

}
