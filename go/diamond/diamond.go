package diamond

import (
	"errors"
	"strings"
)

func spaces(x byte) string {
	res := ""
	for i := byte(0); i < x; i++ {
		res += " "
	}
	return res
}

// Gen creates a dimond shape with a letter input.
func Gen(c byte) (string, error) {
	if (c < 'A') || (c > 'Z') {
		return "", errors.New("input is out of range(A-Z)")
	}
	res := []string{}
	n := int(c-'A')*2 + 1
	mid := n / 2
	char := 'A'
	delta := 0
	dir := 1
	for i := 0; i < n; i++ {
		line := ""
		for j := 0; j < n; j++ {
			if j == mid-delta || j == mid+delta {
				line = line + string(char)
			} else {
				line += " "
			}
		}
		line += "\n"
		res = append(res, line)
		if i >= mid {
			dir = -1
		}
		delta += dir
		char += rune(dir)
	}
	return strings.Join(res, ""), nil
}
