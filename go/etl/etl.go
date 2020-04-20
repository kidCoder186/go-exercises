package etl

import "strings"

// Transform transforms the legacy data format to the new format.
func Transform(m map[int][]string) map[string]int {
	res := map[string]int{}
	for k, list := range m {
		for _, v := range list {
			res[strings.ToLower(v)] = k
		}
	}
	return res
}
