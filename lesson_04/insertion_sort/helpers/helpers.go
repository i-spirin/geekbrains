package helpers

import (
	"log"
	"strconv"
)

// StrToSlice turns string with digits to slice of integers
func StrToSlice(numbers string) []int {
	var numbersList []int
	for _, el := range numbers {
		number, err := strconv.Atoi(string(el))
		if err != nil {
			log.Fatalf("%c is not a number", el)
		}
		numbersList = append(numbersList, number)
	}
	return numbersList
}
