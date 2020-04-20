package summultiples

// SumMultiples returns sum of all multiples of a list number in divisors that less than limit.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	for i := 1; i < limit; i++ {
		for _, val := range divisors {
			if val != 0 && i%val == 0 {
				sum += i
				break
			}
		}
	}
	return sum
}
