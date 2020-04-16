// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isSepChar(c rune) bool {
	switch c {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0, '-':
		return true
	}
	return false
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	var res []rune
	words := strings.FieldsFunc(s, isSepChar)
	for _, word := range words {
		if firstLetter := strings.IndexFunc(word, isLetter); firstLetter != -1 {
			res = append(res, rune(word[firstLetter]))
		}
	}
	return strings.ToUpper(string(res))
}
