package pangram

import (
	"strings"
)

// IsPangram checks a string whether is a pangram or not.
func IsPangram(s string) bool {
	s = strings.ToLower(s)
	visit := map[rune]bool{}
	count := 0
	for _, v := range s {
		if v >= 'a' && v <= 'z' {
			if !visit[v] {
				visit[v] = true
				count++
			}
		}
	}
	return count == 26
}
