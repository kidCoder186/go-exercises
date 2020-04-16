package hamming

import "errors"

// Distance measures hamming distance between two DNA strands.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("2 strands has diffrent length")
	}
	res := 0
	for i := range a {
		if a[i] != b[i] {
			res++
		}
	}
	return res, nil
}
