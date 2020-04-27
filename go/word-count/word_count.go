package wordcount

import (
	"strings"
	"unicode"
)

// Frequency type
type Frequency map[string]int

//WordCount counts the occurrences of each word in the phrase.
func WordCount(phrase string) Frequency {
	freq := Frequency{}
	words := strings.FieldsFunc(phrase, func(r rune) bool {
		return (unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r)) && r != '\''
	})
	for _, v := range words {
		word := strings.TrimFunc(v, func(r rune) bool {
			return !(unicode.IsLetter(r) || unicode.IsDigit(r))
		})
		word = strings.ToLower(word)
		freq[word]++
	}
	return freq
}
