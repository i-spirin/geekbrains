package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/i-spirin/geekbrains/lesson_03/prime_numbers/helpers"
)

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
		if helpers.IsPrime(i) {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
