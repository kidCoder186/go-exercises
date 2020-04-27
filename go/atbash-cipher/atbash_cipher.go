package atbash

import (
	"strings"
	"unicode"
)

// Atbash returns atbash cipher of a string.
func Atbash(s string) string {
	res := []rune{}
	word := []rune{}
	for _, val := range strings.ToLower(s) {
		if unicode.IsLetter(val) {
			word = append(word, 'z'-val+'a')
		} else if !unicode.IsSpace(val) && !unicode.IsPunct(val) {
			word = append(word, val)
		}
		if len(word) == 5 {
			res = append(res, word...)
			res = append(res, ' ')
			word = []rune{}
		}
	}
	res = append(res, word...)
	return strings.TrimSpace(string(res))
}
