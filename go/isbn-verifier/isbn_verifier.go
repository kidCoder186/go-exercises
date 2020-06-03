package isbn

import (
	"strings"
)

func isValid(c byte, pos int) bool {
	return (c >= '0' && c <= '9') || (pos == 9 && c == 'X')
}

// IsValidISBN checks the validation of a isbn string
func IsValidISBN(isbn string) bool {
	isbn = strings.Join(strings.Split(isbn, "-"), "")
	if len(isbn) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < len(isbn); i++ {
		if isValid(isbn[i], i) {
			if isbn[i] == 'X' {
				sum += 10
			} else {
				sum += int(isbn[i]-48) * (10 - i)
			}

		} else {
			return false
		}
	}
	return sum%11 == 0
}
