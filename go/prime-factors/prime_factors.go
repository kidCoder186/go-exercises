package prime

import "math"

// Factors returns all prime factors of a given natural number.
func Factors(n int64) []int64 {
	res := []int64{}
	limit := int64(math.Sqrt(float64(n)))
	for n%2 == 0 {
		res = append(res, 2)
		n = n / 2
	}
	for i := int64(3); i <= limit; i += 2 {
		if n%i == 0 {
			for n%i == 0 {
				res = append(res, i)
				n = n / i
			}
		}
	}
	if n > 1 {
		res = append(res, n)
	}
	return res
}
