package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram chekcks a string whether is an isogram or not.
func IsIsogram(s string) bool {
	visit := map[rune]int{}
	for _, v := range strings.ToUpper(s) {
		if unicode.IsLetter(v) {
			if visit[v] == 0 {
				visit[v]++
			} else {
				return false
			}
		}
	}
	return true
}
