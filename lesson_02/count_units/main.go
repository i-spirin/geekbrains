package main

import (
	"flag"
	"fmt"
	"os"
)

// Number is a struct for string of integers length of 3
type Number struct {
	hundreds rune
	tens     rune
	units    rune
}

// Print Number as count of hundreds, tens and units
func (n Number) Print() {
	fmt.Println(string(n.hundreds) + " hundreds")
	fmt.Println(string(n.tens) + " tens")
	fmt.Println(string(n.units) + " units")
}

// Parse string to Number structure
func (n *Number) Parse(input string) {
	runes := []rune(input)

	n.hundreds = runes[0]
	n.tens = runes[1]
	n.units = runes[2]

}

func main() {
	var userInput *string
	var n Number

	userInput = flag.String("number", "", "number to decompose")
	flag.Parse()

	if *userInput == "" {
		fmt.Print("Insert number: ")
		fmt.Fscan(os.Stdin, userInput)
	}

	n.Parse(*userInput)

	n.Print()

}
