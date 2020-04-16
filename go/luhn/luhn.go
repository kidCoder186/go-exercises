package luhn

import (
	"strings"
	"unicode"
)

// Valid checks a string whether or not it is valid luhn formula.
func Valid(s string) bool {
	fields := strings.FieldsFunc(s, unicode.IsSpace)
	number := strings.Join(fields, "")
	sum, count := 0, 0
	if len(number) < 2 {
		return false
	}
	for i := len(number) - 1; i >= 0; i = i - 1 {
		if unicode.IsDigit(rune(number[i])) {
			alpha := (int(number[i]) - 48) * (count%2 + 1)
			if alpha > 9 {
				alpha -= 9
			}
			sum += alpha
		} else {
			return false
		}
		count++
	}
	if sum%10 != 0 {
		return false
	}
	return true
}
