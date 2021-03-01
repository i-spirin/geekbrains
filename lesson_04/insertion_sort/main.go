package main

import (
	"fmt"
	"log"
	"os"

	"github.com/i-spirin/geekbrains/lesson_04/insertion_sort/helpers"
	"github.com/i-spirin/geekbrains/lesson_04/insertion_sort/sorting"
)

func main() {

	var numbers string
	fmt.Print("Insert numbers: ")
	fmt.Fscan(os.Stdin, &numbers)

	numbersList := helpers.StrToSlice(numbers)
	log.Printf("Sorting %v\n", numbersList)

	sortedList := sorting.InsertionSort(numbersList)
	// sortedList := bubbleSort(numbersList)
	log.Printf("Sorted %v\n", sortedList)

}
