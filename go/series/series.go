package series

// All returns all substring of s that has length = n.
func All(n int, s string) []string {
	first := 0
	res := []string{}
	for {
		if first+n > len(s) {
			break
		}
		res = append(res, s[first:first+n])
		first++
	}
	return res
}

// UnsafeFirst return first substring of s that has length = n.
func UnsafeFirst(n int, s string) string {
	return s[0:n]
}

// First return first substring of s that has length = n.
func First(n int, s string) (string, bool) {
	if n > len(s) {
		return "", false
	}
	return s[0:n], true
}
