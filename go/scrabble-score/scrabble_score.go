package scrabble

import "strings"

func scrabbleScore(c rune) int {
	switch c {
	case 'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T':
		return 1
	case 'D', 'G':
		return 2
	case 'B', 'C', 'M', 'P':
		return 3
	case 'F', 'H', 'V', 'W', 'Y':
		return 4
	case 'K':
		return 5
	case 'J', 'X':
		return 8
	case 'Q', 'Z':
		return 10
	default:
		return 0
	}
}

// Score calculates Scrabble score for the string input.
func Score(s string) int {
	res := 0
	for _, v := range strings.ToUpper(s) {
		res += scrabbleScore(v)
	}
	return res
}
