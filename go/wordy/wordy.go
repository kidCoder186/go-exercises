package wordy

import (
	"strconv"
	"strings"
)

type Stack struct {
	head int
	data []int
}

func (s *Stack) push(x int) {
	if len(s.data) <= s.head {
		s.data = append(s.data, x)
	} else {
		s.data[s.head] = x
	}
	s.head++
}

func (s *Stack) pop() int {
	if s.head > 0 {
		s.head--
		return s.data[s.head]
	}
	return 0
}

// Answer returns answer of the question `s`.
func Answer(s string) (int, bool) {
	words := strings.Fields(s)
	if strings.Compare(words[0], "What") != 0 ||
		strings.Compare(words[1], "is") != 0 {
		return 0, false
	}
	n := len(words)
	words[n-1] = words[n-1][:len(words[n-1])-1]
	words = append(words, "?")
	stack := Stack{}
	for i := 2; i < len(words); {
		num, err := strconv.Atoi(words[i])
		if err == nil {
			if stack.head > 0 {
				return 0, false
			}
			stack.push(num)
			i++
		} else {
			switch words[i] {
			case "plus", "minus":
				nextNum, err := strconv.Atoi(words[i+1])
				if err != nil || stack.head == 0 {
					return 0, false
				}
				if strings.Compare("minus", words[i]) == 0 {
					nextNum = -nextNum
				}
				stack.push(nextNum)
				i = i + 2
			case "multiplied", "divided":
				nextNum, err := strconv.Atoi(words[i+2])
				if err != nil || stack.head == 0 {
					return 0, false
				}
				preNum := stack.pop()
				if strings.Compare("multiplied", words[i]) == 0 {
					stack.push(nextNum * preNum)
				} else {
					stack.push(preNum / nextNum)
				}
				i = i + 3
			case "?":
				res := 0
				for stack.head > 0 {
					res += stack.pop()
				}
				return res, true
			default:
				return 0, false
			}
		}
	}
	return 0, false
}
