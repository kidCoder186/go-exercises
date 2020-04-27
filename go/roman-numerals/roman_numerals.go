package romannumerals

import "fmt"

var (
	arabics = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans  = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

// ToRomanNumeral converts arabic numbers to roman numerals
func ToRomanNumeral(arabic int) (roman string, err error) {
	if !(arabic > 0 && arabic <= 3000) {
		return roman, fmt.Errorf("number is out of range")
	}
	for i := 0; i < len(arabics); i++ {
		for arabic >= arabics[i] {
			arabic -= arabics[i]
			roman += romans[i]
		}
	}
	return roman, nil
}
