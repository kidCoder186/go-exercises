package prime

import "math"

func isPrime(x int) bool {
	if x == 2 {
		return true
	}
	for i := 2; i <= int(math.Sqrt(float64(x))); i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// Nth returns n-th prime number.
func Nth(n int) (int, bool) {
	if n <= 0 {
		return 0, false
	}
	var i int
	for i = 2; n > 0; i++ {
		if isPrime(i) {
			n--
		}
	}
	return i - 1, true
}
