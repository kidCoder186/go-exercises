package allergies

import (
	"strings"
)

var (
	numSubstance = 8
	Substances   = []string{
		"eggs",
		"peanuts",
		"shellfish",
		"strawberries",
		"tomatoes",
		"chocolate",
		"pollen",
		"cats",
	}
)

// Allergies returns all the substances that has the corresponding allergy score.
func Allergies(score uint) []string {
	res := []string{}
	for i := 0; i < numSubstance; i++ {
		if (score & (1 << i)) == (1 << i) {
			res = append(res, Substances[i])
		}
	}
	return res
}

// AllergicTo checks whether or not a given substance in the list of allergies
// that the corresponding allergy score.
func AllergicTo(score uint, substance string) bool {
	var k int
	for i := 0; i < numSubstance; i++ {
		if strings.Compare(substance, Substances[i]) == 0 {
			k = i
			break
		}
	}
	return score&(1<<k) == (1 << k)
}
