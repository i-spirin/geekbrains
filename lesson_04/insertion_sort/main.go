package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func strToSlice(numbers string) []int {
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

func insertionSort(numbersList []int) []int {
	var numbersListCopy []int
	for _, el := range numbersList {
		numbersListCopy = append(numbersListCopy, el)
	}

	for i := 0; i < len(numbersList); i++ {
		j := i - 1
		key := numbersList[i]
		for j >= 0 && numbersList[j] > key {
			numbersList[j+1] = numbersList[j]
			j--
			numbersList[j+1] = key
		}
	}
	return numbersList
}

func bubbleSort(numbersList []int) []int {
	var numbersListCopy []int
	for _, el := range numbersList {
		numbersListCopy = append(numbersListCopy, el)
	}

	if len(numbersList) < 2 {
		return numbersListCopy
	}

	for i := 0; i < len(numbersListCopy)-1; i++ {
		for j := 0; j < len(numbersListCopy)-i-1; j++ {
			if numbersListCopy[j] > numbersListCopy[j+1] {
				numbersListCopy[j], numbersListCopy[j+1] = numbersListCopy[j+1], numbersListCopy[j]
			}
		}
	}
	return numbersListCopy
}

func main() {

	var numbers string
	fmt.Print("Insert numbers: ")
	fmt.Fscan(os.Stdin, &numbers)

	numbersList := strToSlice(numbers)
	log.Printf("Sorting %v\n", numbersList)

	sortedList := insertionSort(numbersList)
	// sortedList := bubbleSort(numbersList)
	log.Printf("Sorted %v\n", sortedList)

}
