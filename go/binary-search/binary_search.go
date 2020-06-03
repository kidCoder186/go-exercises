package binarysearch

// SearchInts finds index of `key` value in the slice.
func SearchInts(s []int, key int) int {
	if len(s) == 0 {
		return -1
	}
	l := 0
	r := len(s) - 1
	for r-l > 1 {
		mid := (l + r) / 2
		if key > s[mid] {
			l = mid
		} else if key < s[mid] {
			r = mid
		} else {
			return mid
		}
	}
	if s[l] == key {
		return l
	}
	if s[r] == key {
		return r
	}
	return -1
}
