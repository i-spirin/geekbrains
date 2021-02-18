package sorting_test

import (
	"testing"

	"github.com/i-spirin/geekbrains/lesson_04/insertion_sort/sorting"
)

func TestInsertionSort(t *testing.T) {
	numbersList := []int{5, 4, 3, 2, 1}
	forCheck := []int{1, 2, 3, 4, 5}

	sortedNumbersList := sorting.InsertionSort(numbersList)
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

	sortedNumbersList := sorting.BubbleSort(numbersList)
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
		x = sorting.InsertionSort([]int{5, 4, 3, 2, 1})
	}
	GlobalF = x
}

func BenchmarkBubbleSort(b *testing.B) {
	x := []int{}
	for i := 0; i < b.N; i++ {
		x = sorting.InsertionSort([]int{5, 4, 3, 2, 1})
	}
	GlobalF = x
}
