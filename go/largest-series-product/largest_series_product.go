package lsproduct

import (
	"errors"
	"unicode"
)

// LargestSeriesProduct calculates the largest product for a contiguous substring of digits of length n.
func LargestSeriesProduct(input string, span int) (int64, error) {
	if span > len(input) || span < 0 {
		return 0, errors.New("wrong span")
	}
	digits := []rune(input)
	var res int64
	for i := 0; i < len(digits)-span+1; i++ {
		var product int64 = 1
		for j := 0; j < span && i+j < len(digits); j++ {
			if !unicode.IsDigit(digits[i+j]) {
				return 0, errors.New("wrong digits")
			}
			product *= int64(digits[i+j] - '0')
		}
		if product > res {
			res = product
		}
	}
	return res, nil
}
