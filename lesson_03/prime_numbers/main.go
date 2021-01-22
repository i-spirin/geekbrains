package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			// composite number
			return false
		}
	}
	return true
}

func main() {
	var n string
	fmt.Print("Insert number: ")
	_, err := fmt.Scanln(&n)
	if err != nil {
		log.Fatalf("Error reading value: %v", err)
	}

	newN, err := strconv.Atoi(n)
	if err != nil {
		log.Fatalf("Error parsing value: %v", err)
	}

	for i := 2; i <= newN; i++ {
		if isPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
