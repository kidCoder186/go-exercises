package strain

type Ints []int
type Lists [][]int
type Strings []string

// Keep returns a collection containing those elements where the function `f` is true
func (i Ints) Keep(f func(int) bool) Ints {
	var res Ints
	for _, val := range i {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

// Discard returns a collection containing those elements where the function `f` is false
func (i Ints) Discard(f func(int) bool) Ints {
	var res Ints
	for _, val := range i {
		if !f(val) {
			res = append(res, val)
		}
	}
	return res
}

// Keep returns a collection containing those elements where the function `f` is true
func (l Lists) Keep(f func([]int) bool) Lists {
	var res Lists
	for _, val := range l {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}

// Keep returns a collection containing those elements where the function `f` is true
func (s Strings) Keep(f func(string) bool) Strings {
	var res Strings
	for _, val := range s {
		if f(val) {
			res = append(res, val)
		}
	}
	return res
}
