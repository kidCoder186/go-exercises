package rotationalcipher

import (
	"unicode"
)

// RotationalCipher implements Ceasar cipher with `shiftKey` value
func RotationalCipher(plain string, shiftKey int) string {
	var res []rune
	for _, r := range plain {
		if unicode.IsLetter(r) {
			if unicode.IsLower(r) {
				res = append(res, 97+(r+rune(shiftKey)-97)%26)
			} else {
				res = append(res, 65+(r+rune(shiftKey)-65)%26)
			}
		} else {
			res = append(res, r)
		}
	}
	return string(res)
}
