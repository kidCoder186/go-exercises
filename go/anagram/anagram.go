package anagram

import "strings"

// Detect returns all anagrams of the `subject` from given `candidates`.
func Detect(subject string, candidates []string) []string {
	res := []string{}
	counts := map[rune]int{}
	for _, v := range strings.ToLower(subject) {
		counts[v]++
	}
	for _, candidate := range candidates {
		lowerCandidate := strings.ToLower(candidate)
		if len(lowerCandidate) == len(subject) &&
			strings.Compare(lowerCandidate, strings.ToLower(subject)) != 0 {
			visit := map[rune]int{}
			for _, v := range lowerCandidate {
				visit[v]++
				if visit[v] > counts[v] {
					break
				}
			}
			ok := true
			for _, v := range subject {
				if visit[v] != counts[v] {
					ok = false
					break
				}
			}
			if ok {
				res = append(res, candidate)
			}
		}

	}
	return res
}
