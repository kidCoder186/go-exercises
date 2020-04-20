package brackets

var (
	pairsMap = map[rune]rune{
		']': '[',
		'}': '{',
		')': '(',
	}
)

// Stack structure
type Stack struct {
	top  int
	data []rune
}

func (s *Stack) isEmpty() bool {
	return s.top == 0
}

func (s *Stack) pop() rune {
	if s.top > 0 {
		s.top--
		return s.data[s.top]
	}
	return 0
}

func (s *Stack) push(x rune) {
	if s.top >= len(s.data) {
		s.data = append(s.data, x)
	} else {
		s.data[s.top] = x
	}
	s.top++
}

// Bracket checks a combination of brackets [], braces {}, parentheses () valid or not.
func Bracket(input string) bool {
	s := []rune(input)
	stack := Stack{0, []rune{}}
	for _, v := range s {
		switch v {
		case '{', '[', '(':
			stack.push(v)
		case '}', ']', ')':
			p := stack.pop()
			if p != pairsMap[v] {
				return false
			}
		}
	}
	return stack.isEmpty()
}
