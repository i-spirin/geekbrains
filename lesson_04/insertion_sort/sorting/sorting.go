package sorting

// InsertionSort algorithm
func InsertionSort(numbersList []int) []int {
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

// BubbleSort algorithm
func BubbleSort(numbersList []int) []int {
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
