package ocr

import "strings"

var (
	Ocr2Dec = map[string]string{
		" _ | ||_|   ": "0",
		"     |  |   ": "1",
		" _  _||_    ": "2",
		" _  _| _|   ": "3",
		"   |_|  |   ": "4",
		" _ |_  _|   ": "5",
		" _ |_ |_|   ": "6",
		" _   |  |   ": "7",
		" _ |_||_|   ": "8",
		" _ |_| _|   ": "9",
	}
)

func recognizeDigit(s string) string {
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n'
	})
	digit := strings.Join(fields, "")
	res, ok := Ocr2Dec[digit]
	if ok {
		return res
	}
	return "?"
}

// Recognize converts all ocr numbers to decimal number.
func Recognize(s string) []string {
	res := []string{}
	lines := strings.FieldsFunc(s, func(r rune) bool {
		return r == '\n'
	})
	for i := 0; i < len(lines)-3; i = i + 4 {
		var line string
		for j := 0; j < len(lines[i])-2; j = j + 3 {
			digit := lines[i][j:j+3] + lines[i+1][j:j+3] +
				lines[i+2][j:j+3] + lines[i+3][j:j+3]
			line = line + recognizeDigit(digit)
		}
		res = append(res, line)
	}
	return res
}
