package helpers_test

import (
	"testing"

	"github.com/i-spirin/geekbrains/lesson_04/insertion_sort/helpers"
)

func TestStrToSlice(t *testing.T) {
	numbers := "54321"
	numbersList := helpers.StrToSlice(numbers)
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
