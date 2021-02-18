package main

import "testing"

func TestStrToSlice(t *testing.T) {
	numbers := "54321"
	numbersList := strToSlice(numbers)
	forCheck := []int{5, 4, 3, 2, 1}
	if len(numbersList) != len(forCheck) {
		t.Errorf("strToSlice returned wrong length")
	}
	for i, el := range numbersList {
		if el != forCheck[i] {
			t.Errorf("numbersList is different. Want %v, got %v", forCheck, numbersList)
		}
	}
}

func TestInsertionSort(t *testing.T) {
	numbersList := []int{5, 4, 3, 2, 1}
	forCheck := []int{1, 2, 3, 4, 5}

	sortedNumbersList := insertionSort(numbersList)
	if len(sortedNumbersList) != len(forCheck) {
		t.Errorf("insertionSort returned wrong length slice")
	}
	for i, el := range sortedNumbersList {
		if el != forCheck[i] {
			t.Errorf("numbersList is different. Want %v, got %v", forCheck, sortedNumbersList)
		}
	}

}

func TestBubbleSort(t *testing.T) {
	numbersList := []int{5, 4, 3, 2, 1}
	forCheck := []int{1, 2, 3, 4, 5}

	sortedNumbersList := bubbleSort(numbersList)
	if len(sortedNumbersList) != len(forCheck) {
		t.Errorf("bubbleSort returned wrong length slice")
	}
	for i, el := range sortedNumbersList {
		if el != forCheck[i] {
			t.Errorf("numbersList is different. Want %v, got %v", forCheck, sortedNumbersList)
		}
	}

}

var GlobalF = []int{}

func BenchmarkInsertionSort(b *testing.B) {
	x := []int{}
	for i := 0; i < b.N; i++ {
		x = insertionSort([]int{5, 4, 3, 2, 1})
	}
	GlobalF = x
}

func BenchmarkBubbleSort(b *testing.B) {
	x := []int{}
	for i := 0; i < b.N; i++ {
		x = insertionSort([]int{5, 4, 3, 2, 1})
	}
	GlobalF = x
}
