package piglatin

import (
	"strings"
)

var (
	vowels = []string{"a", "u", "o", "e", "i", "xr", "yt"}
)

func word2Pig(s string) string {
	var res, vowel string
	res = s
	posVowel := -1
	for _, v := range vowels {
		if index := strings.Index(s, v); index != -1 {
			if index < posVowel || posVowel == -1 {
				posVowel = index
				vowel = v
			}
		}
	}
	posY := strings.Index(s, "y")
	posQU := strings.Index(s, "qu")
	if posY == 2 || posVowel == -1 {
		posVowel = posY
	}
	if posQU != -1 && strings.Compare(vowel, "u") == 0 {
		posVowel = posQU + 2
	}
	if posVowel != -1 {
		res = s[posVowel:]
		res = res + s[:posVowel] + "ay"
	}
	return res
}

// Sentence translates a sentence from English to Pig Latin language.
func Sentence(s string) string {
	res := []string{}
	words := strings.Fields(s)
	for _, word := range words {
		res = append(res, word2Pig(word))
	}
	return strings.Join(res, " ")
}
