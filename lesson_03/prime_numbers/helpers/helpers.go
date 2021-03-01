package helpers

import "math"

// IsPrime function to check number for primary
func IsPrime(n int) bool {
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
