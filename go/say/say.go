package say

import "strings"

var (
	num2Eng = map[int64]string{
		0: "zero", 1: "one", 2: "two", 3: "three", 4: "four",
		5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
		10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen",
		15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen",
		20: "twenty", 30: "thirty", 40: "forty", 50: "fifty",
		60: "sixty", 70: "seventy", 80: "eighty", 90: "ninety",
	}
	suffix = []string{"hundred", "thousand", "million", "billion", "trillion"}
)

// Spell spells out numbers from 0-999
func Spell(n int64) string {
	if n > 0 && n < 100 {
		if n < 20 {
			return num2Eng[n]
		}
		r := n % 10
		res := num2Eng[n-r]
		if r != 0 {
			res += "-"
			res += num2Eng[r]
		}
		return res
	} else if n >= 100 && n < 1000 {
		r := n % 100
		res := Spell(n/100) + " hundred"
		if r != 0 {
			res += " " + Spell(r)
		}
		return res
	}
	return ""
}

// Say spells out a number to English.
func Say(n int64) (string, bool) {
	if n < 0 || n >= 1e12 {
		return "", false
	}
	if n == 0 {
		return "zero", true
	}
	parts := []int64{}
	for n > 0 {
		parts = append(parts, n%1000)
		n /= 1000
	}
	res := ""
	for i := len(parts) - 1; i > 0; i-- {
		if parts[i] > 0 {
			res = res + Spell(parts[i]) + " " + suffix[i] + " "
		}
	}
	res = res + Spell(parts[0])
	return strings.TrimSpace(res), true
}
