package armstrong

import "math"

// IsNumber checks a number whether is a armstrong number or not
func IsNumber(n int) bool {
	var sum float64
	digits := []int{}
	temp := n
	for temp > 0 {
		r := temp % 10
		digits = append(digits, r)
		temp /= 10
	}
	length := len(digits)
	for i := 0; i < len(digits); i++ {
		sum = sum + math.Pow(float64(digits[i]), float64(length))
	}
	return int(sum) == n
}
