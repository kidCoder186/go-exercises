// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import "strings"

var (
	bobAnswers []string = []string{"Sure.",
		"Whoa, chill out!",
		"Calm down, I know what I'm doing!",
		"Fine. Be that way!",
		"Whatever.",
	}
)

func isQuestion(s string) bool {
	if strings.ContainsRune(s, '?') && strings.IndexRune(s, '?') == len(s)-1 {
		return true
	}
	return false
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func isYell(s string) bool {
	return strings.Compare(strings.ToUpper(s), s) == 0 &&
		strings.IndexFunc(s, isLetter) != -1

}

func isEmpty(s string) bool {
	return len(s) == 0
}

// Hey should have a comment documenting it.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	var res string = bobAnswers[4]
	if isYell(remark) {
		if isQuestion(remark) {
			res = bobAnswers[2]
		} else {
			res = bobAnswers[1]
		}
	} else if isQuestion(remark) {
		res = bobAnswers[0]
	} else if isEmpty(remark) {
		res = bobAnswers[3]
	}
	return res
}
