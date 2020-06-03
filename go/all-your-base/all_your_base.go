package allyourbase

import (
	"fmt"
)

// ConvertToBase converts a number from `inBase` to `outBase`
func ConvertToBase(inBase int, digits []int, outBase int) ([]int, error) {
	if inBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}
	if len(digits) == 0 {
		digits = append(digits, 0)
	}

	var res []int
	var decNum int
	for _, digit := range digits {
		if digit < 0 || digit >= inBase {
			return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
		}
		decNum = decNum*inBase + digit
	}
	if decNum == 0 {
		res = append(res, 0)
	}
	for decNum > 0 {
		d := decNum % outBase
		res = append(res, 0)
		copy(res[1:], res[0:])
		res[0] = d
		decNum = decNum / outBase
	}
	return res, nil
}
