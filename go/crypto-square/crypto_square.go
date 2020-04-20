package cryptosquare

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

// Rect structure
type Rect struct {
	nRows int
	nCols int
	data  [][]rune
	dataT [][]rune
}

func newRect(s []rune) *Rect {
	var r, c int
	n := len(s)
	r = int(math.Sqrt(float64(n)))
	if r*r == n {
		c = r
	} else if r*(r+1) >= n {
		c = r + 1
	} else {
		r, c = r+1, r+1
	}
	for i := n; i < r*c; i++ {
		s = append(s, ' ')
	}
	data := make([][]rune, r)
	for i := 0; i < r; i++ {
		data[i] = []rune{}
		data[i] = append(data[i], s[i*c:i*c+c]...)
	}

	dataT := make([][]rune, c)
	for j := 0; j < c; j++ {
		dataT[j] = make([]rune, r)
		for i := 0; i < r; i++ {
			dataT[j][i] = data[i][j]
		}
	}
	return &Rect{r, c, data, dataT}
}

func (r *Rect) String() string {
	res := []string{}
	for i := 0; i < r.nCols; i++ {
		res = append(res, string(r.dataT[i]))
		res = append(res, " ")
	}
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return strings.Join(res, "")
}

// Encode encodes the input string `s` to crypto square.
func Encode(s string) string {
	s = strings.ToLower(s)
	norStr := []rune{}
	for _, c := range s {
		if unicode.IsLetter(c) || unicode.IsDigit(c) {
			norStr = append(norStr, c)
		}
	}
	rect := newRect(norStr)
	return fmt.Sprint(rect)
}
